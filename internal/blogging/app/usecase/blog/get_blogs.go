package blog

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
	"strings"
)

type GetBlogsParams struct {
	Page        int32
	PageSize    int32
	SearchName  *string
	SearchBy    *string
	CategoryIds *string
	AuthorIds   *string
}

type GetBlogsResult struct {
	Blogs      []model.Blog
	Pagination model.Pagination
}

type GetBlogsHandler decorator.UsecaseHandler[GetBlogsParams, GetBlogsResult]

type getBogsHandler struct {
	userRepo user.Repository
	blogRepo blog.Repository
}

func NewGetBlogsHandler(
	userRepo user.Repository,
	blogRepo blog.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetBlogsHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetBlogsParams, GetBlogsResult](
		getBogsHandler{
			userRepo: userRepo,
			blogRepo: blogRepo,
		},
		logger,
		metrics,
	)
}

func (g getBogsHandler) Handle(ctx context.Context, param GetBlogsParams) (GetBlogsResult, error) {
	err := param.Validate()
	if err != nil {
		return GetBlogsResult{}, err
	}

	var authorIds []string
	if param.AuthorIds != nil {
		authorIds = strings.Split(*param.AuthorIds, ",")
	}
	var categoryIds []string
	if param.CategoryIds != nil {
		categoryIds = strings.Split(*param.CategoryIds, ",")
	}

	blogs, total, err := g.blogRepo.GetBlogsByParam(ctx, blog.BlogsParams{
		Page:        param.Page,
		PageSize:    param.PageSize,
		SearchName:  param.SearchName,
		SearchBy:    param.SearchBy,
		AuthorIds:   authorIds,
		CategoryIds: categoryIds,
	})
	if err != nil {
		return GetBlogsResult{}, err
	}

	return GetBlogsResult{
		Blogs: blogs,
		Pagination: model.Pagination{
			Total:    total,
			Page:     param.Page,
			PageSize: param.PageSize,
		},
	}, err
}

func (p *GetBlogsParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return nil
}
