package blog

import (
	"context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	InsertBlog(
		ctx context.Context,
		blog *model.Blog,
	) (*model.Blog, error)

	GetBlogsByParam(
		ctx context.Context,
		param BlogsParams,
	) ([]model.Blog, int32, error)

	GetBlogByID(
		ctx context.Context,
		blogId string,
	) (*model.Blog, error)

	UpdateBlog(
		ctx context.Context,
		blog *model.Blog,
		categoryIds []string,
	) (*model.Blog, error)

	ChangeDeprecatedBlog(
		ctx context.Context,
		blogId string,
		currentStatus bool,
	) error
}
