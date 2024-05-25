package posts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"ozon-test/configs"
	utils "ozon-test/pkg"
	"ozon-test/services/posts/delivery/graph/model"
	"time"
)

//go:generate mockgen -source=posts_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IPostsRepository interface {
	GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error)
	GetPosts(limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(comment *model.Comment) (bool, error)
	CheckPost(id uint64) (bool, error)
	CheckCommentByPost(postId uint64, parentId uint64) (bool, error)
	GetCommentsByPostId(id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsCommentID(id uint64, limit *int, offset *int) ([]*model.Comment, error)
}

type Repository struct {
	db *sql.DB
}

func GetPsxRepo(config *configs.DbPsxConfig, logger *logrus.Logger) (*Repository, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open error: %s", err.Error())
	}

	repo := &Repository{db: db}

	errs := make(chan error)
	go func() {
		errs <- repo.pingDb(3, logger)
	}()

	if err := <-errs; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	logger.Info("Successfully connected to database")

	return repo, nil
}

func (r *Repository) pingDb(timer uint32, log *logrus.Logger) error {
	var err error
	var retries int

	for retries < utils.MaxRetries {
		err = r.db.Ping()
		if err == nil {
			return nil
		}

		retries++
		log.Errorf("sql ping error: %s", err.Error())
		time.Sleep(time.Duration(timer) * time.Second)
	}

	return fmt.Errorf("sql max pinging error: %s", err.Error())
}

func (r *Repository) GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error) {
	post := &model.Post{}
	author := &model.User{}

	rows, err := r.db.Query("SELECT id, user_id, content, created_at, comments_allowed  FROM posts WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("sql get posts error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		post.Author = author
		err = rows.Scan(&post.ID, &author.ID, &post.Content, &post.CreatedAt, &post.AllowComments)
		if err != nil {
			return nil, fmt.Errorf("sql scan error: %s", err.Error())
		}

		comments, err := r.GetCommentsByPostId(id, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("get comments error: %s", err.Error())
		}

		post.Comments = comments
	}

	return post, nil
}

func (r *Repository) GetPosts(limit *int, offset *int) ([]*model.Post, error) {
	var results []*model.Post

	rows, err := r.db.Query("SELECT id, user_id, content, created_at, comments_allowed  FROM posts OFFSET $1 LIMIT $2", *offset, *limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := &model.Post{Author: &model.User{}}

		err = rows.Scan(&result.ID, &result.Author.ID, &result.Content, &result.CreatedAt, &result.AllowComments)
		if err != nil {
			return nil, fmt.Errorf("scan error: %s", err.Error())
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *Repository) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
	err := r.db.QueryRow("INSERT INTO posts(content, comments_allowed) VALUES ($1, $2) RETURNING id", post.Content, post.AllowComments).Scan(&post.ID)
	if err != nil {
		return false, fmt.Errorf("insert post error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) CheckPost(id uint64) (bool, error) {
	var postId uint64

	err := r.db.QueryRow("SELECT id FROM posts where id = $1", id).Scan(&postId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("check comment error: %s", err.Error())
	}

	if postId == 0 {
		return false, nil
	}

	return true, nil
}

func (r *Repository) CheckCommentByPost(postId uint64, parentId uint64) (bool, error) {
	var id uint64

	err := r.db.QueryRow("SELECT comments.id FROM comments where comments.id = $1 and comments.post_id = $2", parentId, postId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("check comment error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) CreateComment(comment *model.Comment) (bool, error) {
	err := r.db.QueryRow("INSERT INTO comments(user_id, post_id, content, parent_id) VALUES ($1, $2, $3, $4) RETURNING id",
		comment.Author.ID, comment.Post.ID, comment.Content, comment.ParentID).Scan(&comment.ID)
	if err != nil {
		return false, fmt.Errorf("insert comment error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) GetCommentsByPostId(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	var results []*model.Comment

	rows, err := r.db.Query("SELECT id, user_id, parent_id, content, created_at  FROM comments WHERE post_id = $1 OFFSET $2 LIMIT $3", id, *offset, *limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := &model.Comment{Author: &model.User{}}

		err = rows.Scan(&result.ID, &result.Author.ID, &result.ParentID, &result.Content, &result.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan comments error: %s", err.Error())
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *Repository) GetCommentsCommentID(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	var results []*model.Comment

	rows, err := r.db.Query("SELECT id, user_id, parent_id, content, created_at  FROM comments WHERE parent_id = $1 OFFSET $2 LIMIT $3", id, *offset, *limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := &model.Comment{Author: &model.User{}}

		err = rows.Scan(&result.ID, &result.Author.ID, &result.ParentID, &result.Content, &result.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan comments error: %s", err.Error())
		}

		results = append(results, result)
	}

	return results, nil
}
