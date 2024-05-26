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
	"time"
)

type IPostRepo interface {
	GetPost(ctx context.Context, id uint64) (*model.Post, error)
	GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(ctx context.Context, comment *model.Comment) (bool, error)
	CheckPost(ctx context.Context, id uint64) (bool, error)
	CheckComment(ctx context.Context, id uint64) (bool, error)
	//CheckCommentByPost(ctx context.Context, postId uint64, parentId uint64) (bool, error)
	GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
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

func (r RedisRepo) GetPosts(ctx context.Context, limit, offset *int) ([]*model.Post, error) {
	var posts []*model.Post

	var cursor uint64
	var keys []string
	var err error
	var count int

	// Вычисляем начальный и конечный индексы для пагинации
	start := *offset
	end := *offset + *limit - 1

	for {
		keys, cursor, err = r.db.Scan(ctx, cursor, "post:*", 10).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve keys: %s", err.Error())
		}

		for _, key := range keys {
			if count < start {
				count++
				continue
			}

			if count > end {
				break
			}

			res, err := r.db.HGetAll(ctx, key).Result()
			if err != nil {
				return nil, fmt.Errorf("get post error: %s", err.Error())
			}

			boolen, err := strconv.ParseBool(res["comments_allowed"])
			if err != nil {
				return nil, fmt.Errorf("parse comments_allowed error: %s", err.Error())
			}

			numStr := key[len("post:"):]
			//num, err := strconv.ParseInt(numStr, 10, 64)
			//if err != nil {
			//	return nil, fmt.Errorf("parse post error: %s", err.Error())
			//}

			post := &model.Post{
				ID:            numStr,
				Content:       res["content"],
				AllowComments: &boolen,
				CreatedAt:     res["created_at"],
				Author: &model.User{
					ID:    res["user_id"],
					Login: "",
				},
				Comments: make([]*model.Comment, 0),
			}

			posts = append(posts, post)
			count++
		}

		if cursor == 0 || count > end {
			break
		}
	}

	return posts, nil
}

func (r RedisRepo) CreateComment(ctx context.Context, comment *model.Comment) (bool, error) {
	randomNumber := strconv.Itoa(rand.Intn(1 << 31))
	timeNow := time.Now()

	err := r.db.HSet(ctx, fmt.Sprintf("comment:%s", randomNumber), "content", comment.Content, "post_id", comment.Post.ID, "created_at", timeNow, "user_id", comment.Author.ID, "parent_id", comment.ParentID).Err()
	if err != nil {
		return false, fmt.Errorf("insert post error: %s", err.Error())
	}

	comment.ID = randomNumber
	comment.CreatedAt = timeNow.String()

	return true, nil
}

func (r RedisRepo) CheckPost(ctx context.Context, id uint64) (bool, error) {
	strId := strconv.FormatUint(id, 10)

	result, err := r.db.HGetAll(ctx, fmt.Sprintf("post:%s", strId)).Result()
	if err != nil {
		return false, fmt.Errorf("get post error: %s", err.Error())
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (r RedisRepo) CheckComment(ctx context.Context, id uint64) (bool, error) {
	strId := strconv.FormatUint(id, 10)

	result, err := r.db.HGetAll(ctx, fmt.Sprintf("comment:%s", strId)).Result()
	if err != nil {
		return false, fmt.Errorf("get post error: %s", err.Error())
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (r RedisRepo) CheckCommentByPost(ctx context.Context, postId uint64, parentId uint64) (bool, error) {
	strId := strconv.FormatUint(postId, 10)

	result, err := r.db.HGetAll(ctx, fmt.Sprintf("post:%s", strId)).Result()
	if err != nil {
		return false, fmt.Errorf("get post error: %s", err.Error())
	}

	if len(result) == 0 {
		return false, nil
	}

	return true, nil
}

func (r RedisRepo) GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	var comments []*model.Comment
	var cursor uint64
	var keys []string
	var err error
	var count int

	start := *offset
	end := *offset + *limit - 1

	for {
		keys, cursor, err = r.db.Scan(ctx, cursor, "comment:*", 10).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve keys: %s", err.Error())
		}

		for _, key := range keys {
			if count < start {
				count++
				continue
			}

			if count > end {
				break
			}

			res, err := r.db.HGetAll(ctx, key).Result()
			if err != nil {
				return nil, fmt.Errorf("get comment error: %s", err.Error())
			}

			numStr := key[len("comment:"):]
			postId := strconv.FormatUint(id, 10)

			if res["post_id"] != postId {
				continue
			}

			comment := &model.Comment{
				ID:        numStr,
				Content:   res["content"],
				CreatedAt: res["created_at"],
				ParentID:  res["parent_id"],
				Post: &model.Post{
					ID: res["post_id"],
				},
				Author: &model.User{
					ID:    res["user_id"],
					Login: "",
				},
			}

			comments = append(comments, comment)
			count++
		}

		if cursor == 0 || count > end {
			break
		}
	}

	return comments, nil
}

func (r RedisRepo) GetCommentsCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	var comments []*model.Comment
	var cursor uint64
	var keys []string
	var err error
	var count int

	start := *offset
	end := *offset + *limit - 1

	for {
		keys, cursor, err = r.db.Scan(ctx, cursor, "comment:*", 10).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve keys: %s", err.Error())
		}

		for _, key := range keys {
			if count < start {
				count++
				continue
			}

			if count > end {
				break
			}

			res, err := r.db.HGetAll(ctx, key).Result()
			if err != nil {
				return nil, fmt.Errorf("get comment error: %s", err.Error())
			}

			numStr := key[len("comment:"):]

			if res["parent_id"] != strconv.FormatUint(id, 10) {
				continue
			}

			comment := &model.Comment{
				ID:        numStr,
				Content:   res["content"],
				CreatedAt: res["created_at"],
				ParentID:  res["parent_id"],
				Post: &model.Post{
					ID: res["post_id"],
				},
				Author: &model.User{
					ID:    res["user_id"],
					Login: "",
				},
			}

			comments = append(comments, comment)
			count++
		}

		if cursor == 0 || count > end {
			break
		}
	}

	return comments, nil

}

func (r *RedisRepo) GetPost(ctx context.Context, id uint64) (*model.Post, error) {
	strId := strconv.FormatUint(id, 10)

	result, err := r.db.HGetAll(ctx, fmt.Sprintf("post:%s", strId)).Result()
	if err != nil {
		return nil, fmt.Errorf("get post error: %s", err.Error())
	}

	if len(result) == 0 {
		return nil, nil
	}

	boolen, err := strconv.ParseBool(result["comments_allowed"])
	if err != nil {
		return nil, fmt.Errorf("parse comments_allowed error: %s", err.Error())
	}

	post := &model.Post{
		ID:            strconv.FormatUint(id, 10),
		Content:       result["content"],
		AllowComments: &boolen,
		CreatedAt:     result["created_at"],
		Author: &model.User{
			ID:    result["user_id"],
			Login: "",
		},
		Comments: nil,
	}

	return post, nil
}

func (r *RedisRepo) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
	randomNumber := strconv.Itoa(rand.Intn(1 << 31))
	timeNow := time.Now()

	err := r.db.HSet(ctx, fmt.Sprintf("post:%s", randomNumber), "content", post.Content, "comments_allowed", *post.AllowComments, "created_at", timeNow, "user_id", post.Author.ID).Err()
	if err != nil {
		return false, fmt.Errorf("insert post error: %s", err.Error())
	}

	post.ID = randomNumber
	post.CreatedAt = timeNow.String()

	return true, nil
}
