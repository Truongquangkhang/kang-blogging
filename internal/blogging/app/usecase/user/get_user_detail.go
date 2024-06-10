package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type GetUserDetailParams struct {
	ID          string
	HasFullData bool
}

type GetUserDetailResult struct {
	User model.User
}

type GetUserDetailHandler decorator.UsecaseHandler[GetUserDetailParams, GetUserDetailResult]

type getUserDetailHandler struct {
	userRepo    user.Repository
	blogRepo    blog.Repository
	commentRepo comment.Repository
}

func NewGetUserDetailHandler(
	userRepo user.Repository,
	blogRepo blog.Repository,
	commentRepo comment.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUserDetailHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetUserDetailParams, GetUserDetailResult](
		&getUserDetailHandler{
			userRepo:    userRepo,
			blogRepo:    blogRepo,
			commentRepo: commentRepo,
		},
		logger,
		metricsClient,
	)
}

func (g getUserDetailHandler) Handle(
	ctx context.Context,
	param GetUserDetailParams,
) (GetUserDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return GetUserDetailResult{}, err
	}

	relateUserInfo, err := g.userRepo.GetRelateInfoOfUser(ctx, param.ID, param.HasFullData)
	if err != nil || relateUserInfo == nil {
		return GetUserDetailResult{}, err
	}
	userResponse := relateUserInfo.User
	blogsReponse := relateUserInfo.Blogs
	for index := range blogsReponse {
		blogsReponse[index].User = &userResponse
	}
	userResponse.Blogs = blogsReponse

	commentsResponse := relateUserInfo.Comments
	for index := range commentsResponse {
		commentsResponse[index].User = userResponse
	}
	userResponse.Comments = commentsResponse
	userResponse.TotalComments = &relateUserInfo.TotalComments
	userResponse.TotalBlogs = relateUserInfo.TotalBlogs

	return GetUserDetailResult{
		User: userResponse,
	}, nil
}

func (p GetUserDetailParams) Validate() error {
	if p.ID == "" {
		return errors.NewBadRequestError("user ID is required")
	}
	return nil
}
