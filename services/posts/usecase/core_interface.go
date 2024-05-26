package usecase

import (
	"context"
	"ozon-test/services/posts/delivery/graph/model"
)

//go:generate mockgen -source=core.go -destination=../mocks/core_mock.go -package=mocks
type ICore interface {
	GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error)
	GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(ctx context.Context, comment *model.Comment) (bool, error)
	GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsByCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
}
