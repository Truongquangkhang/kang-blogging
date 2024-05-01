package blog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type GetBlogDetailParams struct {
	BlogID string
}

type GetBlogDetailResult struct {
	Blog *model.Blog
}

type GetBlogDetailHandler decorator.UsecaseHandler[GetBlogDetailParams, GetBlogDetailResult]

type getBlogDetailHandler struct {
	blogRepo blog.Repository
}

func NewGetBlogDetailHandler(
	blogRepo blog.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetBlogDetailHandler {
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetBlogDetailParams, GetBlogDetailResult](
		getBlogDetailHandler{
			blogRepo: blogRepo,
		},
		logger,
		metrics,
	)
}

func (g getBlogDetailHandler) Handle(ctx context.Context, param GetBlogDetailParams) (GetBlogDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return GetBlogDetailResult{}, err
	}

	rs, err := g.blogRepo.GetBlogByID(ctx, param.BlogID)
	if err != nil {
		return GetBlogDetailResult{}, err
	}
	if rs == nil {
		return GetBlogDetailResult{}, errors.NewNotFoundError("blog not found")
	}
	return GetBlogDetailResult{
		Blog: rs,
	}, nil
}

func (p *GetBlogDetailParams) Validate() error {
	if p.BlogID == "" {
		return errors.NewBadRequestError("BlogID is required")
	}
	return nil
}
