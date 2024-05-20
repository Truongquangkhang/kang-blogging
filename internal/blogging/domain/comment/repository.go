package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetBlogComments(
		ctx context.Context, param ParamGetBlogComments,
	) ([]ResultGetBlogComments, int32, error)
}

type ParamGetBlogComments struct {
	Page     int32
	PageSize int32
	BlogID   string
}

type ResultGetBlogComments struct {
	Comment model.Comment
	Replies []model.Comment
}
