package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	GetBlogComments(
		ctx context.Context, param ParamGetBlogComments,
	) ([]ResultGetBlogComments, int32, error)
	InsertComment(
		ctx context.Context,
		comment *model.Comment,
	) (*model.Comment, error)
	GetCommentById(
		ctx context.Context, commentId string,
	) (*model.Comment, error)
	GetCommentsByParams(
		ctx context.Context, params ParamsGetComments,
	) ([]model.Comment, int32, error)
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

type ParamsGetComments struct {
	Page       int32
	PageSize   int32
	SearchName *string
	SortBy     *string
	IsToxicity *bool
	UserIds    []string
}
