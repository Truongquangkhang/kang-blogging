package blog

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/follow"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
	"strings"
)

type GetBlogsParams struct {
	Page            int32
	PageSize        int32
	SearchName      *string
	SearchBy        *string
	CategoryIds     *string
	AuthorIds       *string
	SortBy          *string
	Published       *bool
	IsDeprecated    *bool
	GetBlogRelevant *bool
	CurrentUserID   *string
}

type GetBlogsResult struct {
	Blogs      []model.Blog
	Pagination model.Pagination
}

type GetBlogsHandler decorator.UsecaseHandler[GetBlogsParams, GetBlogsResult]

type getBogsHandler struct {
	userRepo   user.Repository
	blogRepo   blog.Repository
	followRepo follow.Repository
}

func NewGetBlogsHandler(
	userRepo user.Repository,
	blogRepo blog.Repository,
	followRepo follow.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetBlogsHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	if followRepo == nil {
		panic("followRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetBlogsParams, GetBlogsResult](
		getBogsHandler{
			userRepo:   userRepo,
			blogRepo:   blogRepo,
			followRepo: followRepo,
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

	paramsGetBlogs := &blog.BlogsParams{
		Page:         param.Page,
		PageSize:     param.PageSize,
		SearchName:   param.SearchName,
		SearchBy:     param.SearchBy,
		SortBy:       param.SortBy,
		Published:    param.Published,
		IsDeprecated: param.IsDeprecated,
	}
	if param.AuthorIds != nil {
		paramsGetBlogs.AuthorIds = strings.Split(*param.AuthorIds, ",")
	}
	if param.CategoryIds != nil {
		paramsGetBlogs.CategoryIds = strings.Split(*param.CategoryIds, ",")
	}
	if param.GetBlogRelevant != nil && param.CurrentUserID != nil {
		followedIds, err := g.followRepo.GetFollowedIdsByFollowerId(ctx, *param.CurrentUserID)
		if err != nil {
			return GetBlogsResult{}, err
		}
		if len(followedIds) > 0 {
			followedIds = append(followedIds, utils.ToStringValue(param.CurrentUserID))
			paramsGetBlogs.AuthorIds = followedIds
		}
	}

	blogs, total, err := g.blogRepo.GetBlogsByParam(ctx, *paramsGetBlogs)
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
	if p.SortBy == nil {
		p.SortBy = utils.ToStringPointerValue("updated_at")
	} else {
		if *p.SortBy != "updated_at" && *p.SortBy != "total_comment" {
			return errors.NewBadRequestError("invalid params")
		}
	}
	return nil
}
