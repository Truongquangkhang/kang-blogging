package blog

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/blog"
	"kang-blogging/internal/blogging/domain/role"
	"kang-blogging/internal/common/constants"
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
	roleRepo role.Repository
}

func NewDeleteBlogDetailHandler(
	blogRepo blog.Repository,
	roleRepo role.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) DeleteBlogDetailHandler {
	if blogRepo == nil {
		panic("blogRepo is nil")
	}
	if roleRepo == nil {
		panic("roleRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[DeleteBlogDetailParams, DeleteBlogDetailResult](
		deleteBlogDetailHandler{
			blogRepo: blogRepo,
			roleRepo: roleRepo,
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
	if rs.AuthorID != param.DeletedById {
		r, err := g.roleRepo.GetRoleByUserId(ctx, param.DeletedById)
		if err != nil {
			return DeleteBlogDetailResult{}, err
		}
		if r == nil || r.Name != constants.ADMIN_ROLE {
			return DeleteBlogDetailResult{}, errors.NewForbiddenDefaultError()
		}
		return DeleteBlogDetailResult{}, errors.NewForbiddenDefaultError()
	}
	err = g.blogRepo.ChangeDeprecatedBlog(ctx, rs.ID, rs.IsDeprecated)
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
