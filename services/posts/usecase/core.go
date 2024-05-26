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

type Core struct {
	log    *logrus.Logger
	posts  posts_repo.IPostsRepository
	client auth.AuthorizationClient
}

func GetClient(address string) (auth.AuthorizationClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc connect err: %w", err)
	}
	client := auth.NewAuthorizationClient(conn)

	return client, nil
}

func GetPostgresCore(grpcCfg *configs.GrpcConfig, psxCfg *configs.DbPsxConfig, log *logrus.Logger) (*Core, error) {
	repo, err := posts_repo.GetPsxRepo(psxCfg, log)
	if err != nil {
		return nil, fmt.Errorf("get postgres repo error: %s", err.Error())
	}
	log.Info("Postgresql created successful")

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

func GetRedisCore(grpcCfg *configs.GrpcConfig, redisCfg *configs.DbRedisCfg, log *logrus.Logger) (*Core, error) {
	redisRepo, err := posts_repo.GetRedisRepo(redisCfg)
	if err != nil {
		return nil, fmt.Errorf("get auth repo error: %s", err.Error())
	}
	log.Info("Redis created successful")

	authRepo, err := GetClient(grpcCfg.Addr + ":" + grpcCfg.Port)
	if err != nil {
		return nil, fmt.Errorf("get auth repo error: %s", err.Error())
	}

	core := &Core{
		log:    log,
		posts:  redisRepo,
		client: authRepo,
	}

	return core, nil
}

func (c *Core) GetPost(ctx context.Context, id uint64, limit *int, offset *int) (*model.Post, error) {
	postItem, err := c.posts.GetPost(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get poster repo: %s", err.Error())
	}

	if postItem.ID == "" {
		return nil, nil
	}

	userId, err := strconv.ParseUint(postItem.Author.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse user id err: %s", err.Error())
	}

	res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: userId})
	if err != nil {
		return nil, fmt.Errorf("client request error: %s", err.Error())
	}

	if res == nil {
		return nil, nil
	}

	postItem.Author.Login = res.Name

	comments, err := c.posts.GetCommentsByPostId(ctx, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get post comments error: %s", err.Error())
	}

	postItem.Comments = comments

	for _, comment := range comments {
		formatUint, _ := strconv.ParseUint(comment.Author.ID, 10, 64)

		res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: formatUint})
		if err != nil {
			return nil, fmt.Errorf("client request error: %s", err.Error())
		}

		comment.Author.Login = res.Name
	}

	return postItem, nil
}

func (c *Core) GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	posts, err := c.posts.GetPosts(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get posts repo error: %s", err.Error())
	}

	for _, post := range posts {
		formatUint, _ := strconv.ParseUint(post.Author.ID, 10, 64)

		res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: formatUint})
		if err != nil {
			return nil, fmt.Errorf("client request error: %s", err.Error())
		}

		post.Author.Login = res.Name
	}

	return posts, nil
}

func (c *Core) GetCommentsByPostId(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	have, err := c.posts.CheckPost(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("check post error: %s", err.Error())
	}

	if !have {
		return nil, nil
	}

	comments, err := c.posts.GetCommentsByPostId(ctx, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get comments repo error: %s", err.Error())
	}

	for _, comment := range comments {
		formatUint, _ := strconv.ParseUint(comment.Author.ID, 10, 64)

		res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: formatUint})
		if err != nil {
			return nil, fmt.Errorf("client request error: %s", err.Error())
		}

		comment.Author.Login = res.Name
	}

	return comments, nil
}

func (c *Core) GetCommentsByCommentID(ctx context.Context, id uint64, limit *int, offset *int) ([]*model.Comment, error) {
	comments, err := c.posts.GetCommentsCommentID(ctx, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get comments repo error: %s", err.Error())
	}

	for _, comment := range comments {
		formatUint, _ := strconv.ParseUint(comment.Author.ID, 10, 64)

		res, err := c.client.GetUserName(ctx, &auth.UserItemRequest{Id: formatUint})
		if err != nil {
			return nil, fmt.Errorf("client request error: %s", err.Error())
		}

		comment.Author.Login = res.Name
	}

	return comments, nil
}

func (c *Core) CreatePost(ctx context.Context, post *model.Post) (bool, error) {
	result, err := c.posts.CreatePost(ctx, post)
	if err != nil {
		return false, fmt.Errorf("create poster repo error: %s", err.Error())
	}

	return result, nil
}

func (c *Core) CreateComment(ctx context.Context, comment *model.Comment) (bool, error) {
	var checked bool

	postId, err := strconv.ParseUint(comment.Post.ID, 10, 64)
	if err != nil {
		return false, fmt.Errorf("parse user id err: %s", err.Error())
	}

	parentID, err := strconv.ParseUint(comment.ParentID, 10, 64)
	if err != nil {
		return false, fmt.Errorf("parse comment id err: %s", err.Error())
	}

	checked, err = c.posts.CheckPost(ctx, postId)
	if err != nil {
		return false, fmt.Errorf("check post error: %s", err.Error())
	}

	if parentID != 0 {
		checked, err = c.posts.CheckComment(ctx, parentID)
		if err != nil {
			return false, fmt.Errorf("check comment error: %s", err.Error())
		}
	}

	if !checked {
		return false, nil
	}

	result, err := c.posts.CreateComment(ctx, comment)
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
