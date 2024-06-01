package blog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type UpdateBlogDetailParams struct {
	AuthorID    string
	BlogID      string
	Name        *string
	Description *string
	CategoryIDs []string
	Thumbnail   *string
	Content     *string
	Published   *bool
}

type UpdateBlogDetailResult struct {
	Blog *model.Blog
}

type UpdateBlogDetailHandler decorator.UsecaseHandler[UpdateBlogDetailParams, UpdateBlogDetailResult]

type updateBlogDetailHandler struct {
	blogRepo blog.Repository
}

func NewUpdateBlogDetailHandler(
	blogRepo blog.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) UpdateBlogDetailHandler {
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[UpdateBlogDetailParams, UpdateBlogDetailResult](
		updateBlogDetailHandler{
			blogRepo: blogRepo,
		},
		logger,
		metrics,
	)
}

func (g updateBlogDetailHandler) Handle(ctx context.Context, param UpdateBlogDetailParams) (UpdateBlogDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return UpdateBlogDetailResult{}, err
	}

	rs, err := g.blogRepo.GetBlogByID(ctx, param.BlogID)
	if err != nil {
		return UpdateBlogDetailResult{}, err
	}
	if rs == nil {
		return UpdateBlogDetailResult{}, errors.NewNotFoundError("blog not found")
	}
	if rs.AuthorID != param.AuthorID {
		return UpdateBlogDetailResult{}, errors.NewForbiddenDefaultError()
	}

	rs, err = g.updateBlog(ctx, param, rs)
	if err != nil {
		return UpdateBlogDetailResult{}, err
	}

	return UpdateBlogDetailResult{
		Blog: rs,
	}, nil
}

func (g updateBlogDetailHandler) updateBlog(
	ctx context.Context,
	param UpdateBlogDetailParams,
	blog *model.Blog,
) (*model.Blog, error) {
	if param.Name != nil {
		blog.Title = *param.Name
	}
	if param.Description != nil {
		blog.Summary = param.Description
	}
	if param.Thumbnail != nil {
		blog.Thumbnail = param.Thumbnail
	}
	if param.Content != nil {
		blog.Content = param.Content
	}
	if param.Published != nil {
		blog.Published = *param.Published
	}
	rs, err := g.blogRepo.UpdateBlog(ctx, blog, param.CategoryIDs)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (p *UpdateBlogDetailParams) Validate() error {
	if p.BlogID == "" {
		return errors.NewBadRequestError("BlogID is required")
	}
	return nil
}
