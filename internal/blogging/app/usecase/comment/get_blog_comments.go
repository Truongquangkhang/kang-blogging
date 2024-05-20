package comment

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
)

type GetBlogCommentsParams struct {
	Page     int32
	PageSize int32
	BlogID   string
}

type GetBlogCommentsResult struct {
	Comments   []comment.ResultGetBlogComments
	Pagination model.Pagination
}

type GetBlogCommentsHandler decorator.UsecaseHandler[GetBlogCommentsParams, GetBlogCommentsResult]

type getBlogCommentsHandler struct {
	commentRepo comment.Repository
}

func NewGetBlogCommentsHandler(
	commentRepo comment.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetBlogCommentsHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetBlogCommentsParams, GetBlogCommentsResult](
		getBlogCommentsHandler{
			commentRepo: commentRepo,
		},
		logger,
		metrics,
	)
}

func (g getBlogCommentsHandler) Handle(ctx context.Context, param GetBlogCommentsParams) (GetBlogCommentsResult, error) {
	err := param.Validate()
	if err != nil {
		return GetBlogCommentsResult{}, err
	}

	comments, total, err := g.commentRepo.GetBlogComments(ctx, comment.ParamGetBlogComments{
		Page:     param.Page,
		PageSize: param.PageSize,
		BlogID:   param.BlogID,
	})
	if err != nil {
		return GetBlogCommentsResult{}, err
	}

	return GetBlogCommentsResult{
		Comments: comments,
		Pagination: model.Pagination{
			Total:    total,
			Page:     param.Page,
			PageSize: param.PageSize,
		},
	}, err
}

func (p *GetBlogCommentsParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return nil
}
