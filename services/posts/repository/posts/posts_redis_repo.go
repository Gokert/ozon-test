package posts

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/stdlib"
	"math/rand"
	"ozon-test/configs"
	"ozon-test/services/posts/delivery/graph/model"
	"strconv"
)

type IPostRepo interface {
	GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error)
	GetPosts(limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(comment *model.Comment) (bool, error)
	CheckPost(id uint64) (bool, error)
	CheckCommentByPost(postId uint64, parentId uint64) (bool, error)
	GetCommentsByPostId(id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsCommentID(id uint64, limit *int, offset *int) ([]*model.Comment, error)
}

type RedisRepo struct {
	db *redis.Client
}

func GetRedisRepo(cfg *configs.DbRedisCfg) (IPostRepo, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DbNumber,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("ping redis error: %s", err.Error())
	}

	return &RedisRepo{db: redisClient}, nil
}

//func (I RedisRepo) GetPost(id uint64, limit *int, offset *int) (*model.Post, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (I RedisRepo) GetPosts(limit *int, offset *int) ([]*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

//func (I RedisRepo) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (I RedisRepo) CreateComment(comment *model.Comment) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (I RedisRepo) CheckPost(id uint64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (I RedisRepo) CheckCommentByPost(postId uint64, parentId uint64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (I RedisRepo) GetCommentsByPostId(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (I RedisRepo) GetCommentsCommentID(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisRepo) GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error) {
	post := &model.Post{}

	strId := strconv.FormatUint(id, 10)

	// Получение всех полей и значений хеша из Redis
	result, err := r.db.HGetAll(ctx, fmt.Sprintf("post:%s", strId)).Result()
	if err != nil {
		return nil, fmt.Errorf("get post error: %s", err.Error())
	}

	fmt.Println(result, id)

	// Проверка, что хеш найден
	if len(result) == 0 {
		return nil, nil
	}

	// Преобразование значений из хеша в свойства сообщения
	boolen, _ := strconv.ParseBool(result["comments_allowed"])

	post.ID = strconv.FormatUint(id, 10)
	post.Content = result["content"]
	post.AllowComments = &boolen
	return post, nil
}

func (r *RedisRepo) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
	// Генерация ID сообщения. В реальном приложении вам потребуется более надежный способ генерации ID.
	//max1 := uint64(math.MaxUint64)
	randomNumber := strconv.FormatUint(rand.Uint64(), 10)

	// Создание хеша в Redis для хранения сообщения
	err := r.db.HSet(ctx, fmt.Sprintf("post:%s", randomNumber), "content", post.Content, "comments_allowed", *post.AllowComments).Err()
	if err != nil {
		return false, fmt.Errorf("insert post error: %s", err.Error())
	}

	return true, nil
}

//func (r *RedisRepo) GetPost(id uint64, limit *int, offset *int) (*model.Post, error) {
//	post := &model.Post{}
//	author := &model.User{}
//
//	rows, err := r.db.Query("SELECT id, user_id, content, created_at, comments_allowed  FROM posts WHERE id = $1", id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, nil
//		}
//		return nil, fmt.Errorf("sql get posts error: %s", err.Error())
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		post.Author = author
//		err = rows.Scan(&post.ID, &author.ID, &post.Content, &post.CreatedAt, &post.AllowComments)
//		if err != nil {
//			return nil, fmt.Errorf("sql scan error: %s", err.Error())
//		}
//
//		comments, err := r.GetCommentsByPostId(id, limit, offset)
//		if err != nil {
//			return nil, fmt.Errorf("get comments error: %s", err.Error())
//		}
//
//		post.Comments = comments
//	}
//
//	return post, nil
//}

//func (r *RedisRepo) GetPosts(limit *int, offset *int) ([]*model.Post, error) {
//	var results []*model.Post
//
//	rows, err := r.db.Query("SELECT id, user_id, content, created_at, comments_allowed  FROM posts OFFSET $1 LIMIT $2", *offset, *limit)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		result := &model.Post{Author: &model.User{}}
//
//		err = rows.Scan(&result.ID, &result.Author.ID, &result.Content, &result.CreatedAt, &result.AllowComments)
//		if err != nil {
//			return nil, fmt.Errorf("scan error: %s", err.Error())
//		}
//
//		results = append(results, result)
//	}
//
//	return results, nil
//}

//func (r *RedisRepo) CreatePost(post *model.Post) (bool, error) {
//	err := r.db.QueryRow("INSERT INTO posts(content, comments_allowed) VALUES ($1, $2) RETURNING id", post.Content, post.AllowComments).Scan(&post.ID)
//	if err != nil {
//		return false, fmt.Errorf("insert post error: %s", err.Error())
//	}
//
//	return true, nil
//}

//func (r *RedisRepo) CheckPost(id uint64) (bool, error) {
//	var postId uint64
//
//	err := r.db.QueryRow("SELECT id FROM posts where id = $1", id).Scan(&postId)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return false, nil
//		}
//		return false, fmt.Errorf("check comment error: %s", err.Error())
//	}
//
//	if postId == 0 {
//		return false, nil
//	}
//
//	return true, nil
//}
//
//func (r *RedisRepo) CheckCommentByPost(postId uint64, parentId uint64) (bool, error) {
//	var id uint64
//
//	err := r.db.QueryRow("SELECT comments.id FROM comments where comments.id = $1 and comments.post_id = $2", parentId, postId).Scan(&id)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return false, nil
//		}
//		return false, fmt.Errorf("check comment error: %s", err.Error())
//	}
//
//	return true, nil
//}
//
//func (r *RedisRepo) CreateComment(comment *model.Comment) (bool, error) {
//	err := r.db.QueryRow("INSERT INTO comments(user_id, post_id, content, parent_id) VALUES ($1, $2, $3, $4) RETURNING id",
//		comment.Author.ID, comment.Post.ID, comment.Content, comment.ParentID).Scan(&comment.ID)
//	if err != nil {
//		return false, fmt.Errorf("insert comment error: %s", err.Error())
//	}
//
//	return true, nil
//}
//
//func (r *RedisRepo) GetCommentsByPostId(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
//	var results []*model.Comment
//
//	rows, err := r.db.Query("SELECT id, user_id, parent_id, content, created_at  FROM comments WHERE post_id = $1 OFFSET $2 LIMIT $3", id, *offset, *limit)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		result := &model.Comment{Author: &model.User{}}
//
//		err = rows.Scan(&result.ID, &result.Author.ID, &result.ParentID, &result.Content, &result.CreatedAt)
//		if err != nil {
//			return nil, fmt.Errorf("scan comments error: %s", err.Error())
//		}
//
//		results = append(results, result)
//	}
//
//	return results, nil
//}
//
//func (r *RedisRepo) GetCommentsCommentID(id uint64, limit *int, offset *int) ([]*model.Comment, error) {
//	var results []*model.Comment
//
//	rows, err := r.db.Query("SELECT id, user_id, parent_id, content, created_at  FROM comments WHERE parent_id = $1 OFFSET $2 LIMIT $3", id, *offset, *limit)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		result := &model.Comment{Author: &model.User{}}
//
//		err = rows.Scan(&result.ID, &result.Author.ID, &result.ParentID, &result.Content, &result.CreatedAt)
//		if err != nil {
//			return nil, fmt.Errorf("scan comments error: %s", err.Error())
//		}
//
//		results = append(results, result)
//	}
//
//	return results, nil
//}
