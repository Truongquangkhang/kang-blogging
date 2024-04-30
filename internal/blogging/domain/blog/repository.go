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
}
