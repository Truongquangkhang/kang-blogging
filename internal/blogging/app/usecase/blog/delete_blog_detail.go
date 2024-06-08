package blog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type DeleteBlogDetailParams struct {
	DeletedById string
	BlogID      string
}

type DeleteBlogDetailResult struct {
	Blog *model.Blog
}

type DeleteBlogDetailHandler decorator.UsecaseHandler[DeleteBlogDetailParams, DeleteBlogDetailResult]

type deleteBlogDetailHandler struct {
	blogRepo blog.Repository
}

func NewDeleteBlogDetailHandler(
	blogRepo blog.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) DeleteBlogDetailHandler {
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[DeleteBlogDetailParams, DeleteBlogDetailResult](
		deleteBlogDetailHandler{
			blogRepo: blogRepo,
		},
		logger,
		metrics,
	)
}

func (g deleteBlogDetailHandler) Handle(ctx context.Context, param DeleteBlogDetailParams) (DeleteBlogDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return DeleteBlogDetailResult{}, err
	}

	rs, err := g.blogRepo.GetBlogByID(ctx, param.BlogID)
	if err != nil {
		return DeleteBlogDetailResult{}, err
	}
	if rs == nil {
		return DeleteBlogDetailResult{}, errors.NewNotFoundError("blog not found")
	}
	//if rs.AuthorID != param.DeletedById {
	//	return DeleteBlogDetailResult{}, errors.NewForbiddenDefaultError()
	//}
	err = g.blogRepo.DeprecatedBlog(ctx, rs.ID)
	if err != nil {
		return DeleteBlogDetailResult{}, err
	}

	return DeleteBlogDetailResult{
		Blog: rs,
	}, nil
}

func (p *DeleteBlogDetailParams) Validate() error {
	if p.BlogID == "" {
		return errors.NewBadRequestError("BlogID is required")
	}
	return nil
}
