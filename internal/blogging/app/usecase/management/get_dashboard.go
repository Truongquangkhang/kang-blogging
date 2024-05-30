package management

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/category"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type GetDashboardParams struct {
}

type GetDashboardResult struct {
	SystemInfo     user.SystemInfo
	LatestBlogs    []model.Blog
	LatestComments []model.Comment
}

type GetDashboardHandler decorator.UsecaseHandler[GetDashboardParams, GetDashboardResult]

type getDashboardHandler struct {
	commentRepo  comment.Repository
	blogRepo     blog.Repository
	categoryRepo category.Repository
	userRepo     user.Repository
}

func NewGetDashboardHandler(
	commentRepo comment.Repository,
	blogRepo blog.Repository,
	categoryRepo category.Repository,
	userRepo user.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetDashboardHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	if categoryRepo == nil {
		panic("categoryRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetDashboardParams, GetDashboardResult](
		getDashboardHandler{
			commentRepo:  commentRepo,
			userRepo:     userRepo,
			blogRepo:     blogRepo,
			categoryRepo: categoryRepo,
		},
		logger,
		metrics,
	)
}

func (g getDashboardHandler) Handle(ctx context.Context, param GetDashboardParams) (GetDashboardResult, error) {
	err := param.Validate()
	if err != nil {
		return GetDashboardResult{}, err
	}

	tt, err := g.userRepo.GetInfoFromMultiTable(ctx)
	if err != nil {
		return GetDashboardResult{}, err
	}
	latestBlog, _, err := g.blogRepo.GetBlogsByParam(ctx, blog.BlogsParams{
		SortBy:   utils.ToStringPointerValue("updated_at"),
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return GetDashboardResult{}, err
	}
	latestComments, _, err := g.commentRepo.GetCommentsByParams(ctx, comment.ParamsGetComments{
		Page:     1,
		PageSize: 10,
		SortBy:   utils.ToStringPointerValue("created_at"),
	})
	if err != nil {
		return GetDashboardResult{}, err
	}

	return GetDashboardResult{
		SystemInfo:     *tt,
		LatestBlogs:    latestBlog,
		LatestComments: latestComments,
	}, err
}

func (p *GetDashboardParams) Validate() error {
	return nil
}
