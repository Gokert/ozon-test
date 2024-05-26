package graph

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	utils "ozon-test/pkg"
	"ozon-test/pkg/middleware"
	httpResponse "ozon-test/pkg/response"
	"ozon-test/services/posts/delivery/graph/model"
	"ozon-test/services/posts/usecase"
	"strconv"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Core usecase.ICore
	Log  *logrus.Logger
}

func (r *Resolver) GetPost(ctx context.Context, id string, limit *int, offset *int) (*model.Post, error) {
	timeStart := time.Now()

	idConverted, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.Log.Errorf("failed to convert id to uint64: %s", err)
		httpResponse.SendLog(http.StatusBadRequest, utils.GetPost, timeStart, r.Log)
		return nil, fmt.Errorf(utils.ConvertedIdError)
	}

	if limit == nil || offset == nil {
		httpResponse.SendLog(http.StatusBadRequest, utils.GetPost, timeStart, r.Log)
		return nil, fmt.Errorf(utils.PaginatorError)
	}

	poster, err := r.Core.GetPost(ctx, idConverted, limit, offset)
	if err != nil {
		r.Log.Errorf("get post: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.GetPost, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	httpResponse.SendLog(http.StatusOK, utils.GetPost, timeStart, r.Log)
	return poster, nil
}

func (r *Resolver) GetPosts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	timeStart := time.Now()

	if limit == nil || offset == nil {
		httpResponse.SendLog(http.StatusBadRequest, utils.GetPosts, timeStart, r.Log)
		return nil, fmt.Errorf(utils.PaginatorError)
	}

	posts, err := r.Core.GetPosts(ctx, limit, offset)
	if err != nil {
		r.Log.Errorf("get posts: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.GetPosts, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	return posts, nil
}

func (r *Resolver) GetCommentsByPostID(ctx context.Context, id string, limit *int, offset *int) ([]*model.Comment, error) {
	timeStart := time.Now()

	idConverted, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.Log.Errorf("failed to convert id to uint64: %s", err)
		httpResponse.SendLog(http.StatusBadRequest, utils.CommentsByPostId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.ConvertedIdError)
	}

	if limit == nil || offset == nil {
		httpResponse.SendLog(http.StatusBadRequest, utils.CommentsByPostId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.PaginatorError)
	}

	comments, err := r.Core.GetCommentsByPostId(ctx, idConverted, limit, offset)
	if err != nil {
		r.Log.Errorf("get comments by post ID: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.CommentsByPostId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	httpResponse.SendLog(http.StatusOK, utils.CommentsByPostId, timeStart, r.Log)
	return comments, nil
}

func (r *Resolver) GetCommentsByCommentID(ctx context.Context, id string, limit *int, offset *int) ([]*model.Comment, error) {
	timeStart := time.Now()

	idConverted, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.Log.Errorf("failed to convert id to uint64: %s", err)
		httpResponse.SendLog(http.StatusBadRequest, utils.CommentsByCommentId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.ConvertedIdError)
	}

	if limit == nil || offset == nil {
		httpResponse.SendLog(http.StatusBadRequest, utils.CommentsByCommentId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.PaginatorError)
	}

	comments, err := r.Core.GetCommentsByCommentID(ctx, idConverted, limit, offset)
	if err != nil {
		r.Log.Errorf("get comments by comment ID: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.CommentsByCommentId, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	return comments, nil
}

func (r *Resolver) CreatePost(ctx context.Context, content string, allowComments bool) (*model.Post, error) {
	timeStart := time.Now()

	session := ctx.Value(middleware.UserIDKey)
	if session == nil {
		return nil, fmt.Errorf("session is nil")
	}

	id := strconv.FormatUint(session.(uint64), 10)

	post := &model.Post{
		Content:       content,
		AllowComments: &allowComments,
		Author: &model.User{
			ID:    id,
			Login: "",
		},
	}

	_, err := r.Core.CreatePost(ctx, post)
	if err != nil {
		r.Log.Errorf("create post error: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.CreatePost, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	httpResponse.SendLog(http.StatusOK, utils.CreatePost, timeStart, r.Log)
	return post, nil
}

func (r *Resolver) CreateComment(ctx context.Context, postID string, content string, parentID *string) (*model.Comment, error) {
	timeStart := time.Now()

	if parentID == nil {
		return nil, fmt.Errorf("parentID is required")
	}

	session := ctx.Value(middleware.UserIDKey)
	if session == nil {
		return nil, fmt.Errorf("session is nil")
	}

	comment := &model.Comment{
		Content: content,
		Post: &model.Post{
			ID: postID,
		},
		Author: &model.User{
			ID: "1",
		},
		ParentID: *parentID,
	}

	result, err := r.Core.CreateComment(ctx, comment)
	if err != nil {
		r.Log.Errorf("create comment error: %s", err)
		httpResponse.SendLog(http.StatusInternalServerError, utils.CreateComment, timeStart, r.Log)
		return nil, fmt.Errorf(utils.InternalError)
	}

	if result == false {
		httpResponse.SendLog(http.StatusNotFound, utils.CreateComment, timeStart, r.Log)
		return nil, fmt.Errorf(utils.PostOrCommentNotFound)
	}

	httpResponse.SendLog(http.StatusOK, utils.CreatePost, timeStart, r.Log)
	return comment, nil
}
