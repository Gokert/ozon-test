package posts

import (
	"context"
	"ozon-test/services/posts/delivery/graph/model"
)

//go:generate mockgen -source=posts_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IPostsRepository interface {
	GetPost(ctx context.Context, id uint64) (*model.Post, error)
	GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(ctx context.Context, comment *model.Comment) (bool, error)
	CheckPost(ctx context.Context, id uint64) (bool, error)
	CheckComment(ctx context.Context, id uint64) (bool, error)
	GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
}
