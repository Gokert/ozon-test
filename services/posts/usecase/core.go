package usecase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"ozon-test/configs"
	auth "ozon-test/services/authorization/delivery/proto"
	"ozon-test/services/posts/delivery/graph/model"
	posts_repo "ozon-test/services/posts/repository/posts"
	"strconv"
)

//go:generate mockgen -source=core.go -destination=../mocks/core_mock.go -package=mocks
type ICore interface {
	GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error)
	GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (bool, error)
	CreateComment(ctx context.Context, comment *model.Comment) (bool, error)
	GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)
	GetCommentsByCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error)

	GetUserId(ctx context.Context, sid string) (uint64, error)
	GetRole(ctx context.Context, userId uint64) (string, error)
}

type Core struct {
	log    *logrus.Logger
	posts  posts_repo.IPostsRepository
	client auth.AuthorizationClient
}

func GetClient(address string) (auth.AuthorizationClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc connect err: %w", err)
	}
	client := auth.NewAuthorizationClient(conn)

	return client, nil
}

func GetCore(grpcCfg *configs.GrpcConfig, psxCfg *configs.DbPsxConfig, log *logrus.Logger) (*Core, error) {
	repo, err := posts_repo.GetPsxRepo(psxCfg, log)
	if err != nil {
		return nil, fmt.Errorf("get psx error error: %s", err.Error())
	}
	log.Info("Psx created successful")

	authRepo, err := GetClient(grpcCfg.Addr + ":" + grpcCfg.Port)
	if err != nil {
		return nil, fmt.Errorf("get auth repo error: %s", err.Error())
	}

	core := &Core{
		log:    log,
		posts:  repo,
		client: authRepo,
	}

	return core, nil
}

func (c *Core) GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error) {
	postItem, err := c.posts.GetPost(id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get poster repo: %s", err.Error())
	}

	if postItem.ID == "" {
		return nil, nil
	}

	userId, err := strconv.ParseUint(postItem.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse user id err: %s", err.Error())
	}

	res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: userId})
	if err != nil {
		return nil, fmt.Errorf("client request error: %s", err.Error())
	}

	postItem.Author.Login = res.Name

	return postItem, nil
}

func (c *Core) GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	posts, err := c.posts.GetPosts(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get posts repo error: %s", err.Error())
	}

	return posts, nil
}

func (c *Core) GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	comments, err := c.posts.GetCommentsByPostId(id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get comments repo error: %s", err.Error())
	}

	//if len(comments) == 0 {
	//	return nil, nil
	//}

	//res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: userId})
	//if err != nil {
	//	return nil, fmt.Errorf("client request error: %s", err.Error())
	//}

	return comments, nil
}

func (c *Core) GetCommentsByCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	comments, err := c.posts.GetCommentsCommentID(id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get comments repo error: %s", err.Error())
	}

	return comments, nil
}

func (c *Core) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
	result, err := c.posts.CreatePost(post)
	if err != nil {
		return false, fmt.Errorf("create poster repo error: %s", err.Error())
	}

	return result, nil
}

func (c *Core) CreateComment(ctx context.Context, comment *model.Comment) (bool, error) {
	result, err := c.posts.CreateComment(comment)
	if err != nil {
		return false, fmt.Errorf("create comment repo error: %s", err.Error())
	}

	return result, nil
}

func (c *Core) GetUserId(ctx context.Context, sid string) (uint64, error) {
	response, err := c.client.GetId(ctx, &auth.FindIdRequest{Sid: sid})
	if err != nil {
		return 0, fmt.Errorf("get user id err: %w", err)
	}
	return response.Value, nil
}

func (c *Core) GetRole(ctx context.Context, userId uint64) (string, error) {
	role, err := c.client.GetRole(ctx, &auth.RoleRequest{Id: userId})
	if err != nil {
		return "", fmt.Errorf("get role error: %s", err.Error())
	}

	return role.Role, nil
}
