package user

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type GetUserDetailParams struct {
	CurrentUserID *string
	ID            string
}

type GetUserDetailResult struct {
	User model.User
}

type GetUserDetailHandler decorator.UsecaseHandler[GetUserDetailParams, GetUserDetailResult]

type getUserDetailHandler struct {
	userRepo user.Repository
	blogRepo blog.Repository
}

func NewGetUserDetailHandler(
	userRepo user.Repository,
	blogRepo blog.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUserDetailHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetUserDetailParams, GetUserDetailResult](
		&getUserDetailHandler{
			userRepo: userRepo,
			blogRepo: blogRepo,
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

	// Get info of a user
	u, err := g.userRepo.GetUserByID(ctx, param.ID)
	if err != nil || u == nil {
		return GetUserDetailResult{}, errors.NewNotFoundError("user not found")
	}

	paramGetBlog := blog.BlogsParams{
		Page:         1,
		PageSize:     200,
		AuthorIds:    []string{u.ID},
		Published:    utils.ToPointer(true),
		IsDeprecated: utils.ToPointer(false),
	}

	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err == nil && auth != nil {
		if auth.Role == constants.ADMIN_ROLE || auth.UserID == u.ID {
			paramGetBlog.Published = nil
		}
	}

	// Get blogs of the user
	blogs, totalBlog, err := g.blogRepo.GetBlogsByParam(ctx, paramGetBlog)
	if err != nil {
		return GetUserDetailResult{}, err
	}
	u.Blogs = blogs
	u.TotalBlogs = totalBlog
	return GetUserDetailResult{
		User: *u,
	}, nil
}

func (p GetUserDetailParams) Validate() error {
	if p.ID == "" {
		return errors.NewBadRequestError("user ID is required")
	}
	return nil
}
