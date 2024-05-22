package blog_comments

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	InsertBlogComment(
		ctx context.Context,
		blogComment *model.BlogComment,
	) (*model.BlogComment, error)
}
