package comment

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type GetCommentsParams struct {
	Page       int32
	PageSize   int32
	SearchName *string
	SortBy     *string
	IsToxicity *bool
	UserIds    *string
}

type GetCommentsResult struct {
	Comments   []model.Comment
	Pagination model.Pagination
}

type GetCommentsHandler decorator.UsecaseHandler[GetCommentsParams, GetCommentsResult]

type getCommentsHandler struct {
	commentRepo comment.Repository
}

func NewGetCommentsHandler(
	commentRepo comment.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetCommentsHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetCommentsParams, GetCommentsResult](
		getCommentsHandler{
			commentRepo: commentRepo,
		},
		logger,
		metrics,
	)
}

func (g getCommentsHandler) Handle(ctx context.Context, param GetCommentsParams) (GetCommentsResult, error) {
	err := param.Validate()
	if err != nil {
		return GetCommentsResult{}, err
	}

	comments, count, err := g.commentRepo.GetCommentsByParams(
		ctx, comment.ParamsGetComments{
			SearchName: param.SearchName,
			SortBy:     param.SortBy,
			IsToxicity: param.IsToxicity,
			Page:       param.Page,
			PageSize:   param.PageSize,
			UserIds:    utils.SplitStringSeparateCommaToSlice(param.UserIds),
		},
	)
	if err != nil {
		return GetCommentsResult{}, err
	}

	return GetCommentsResult{
		Comments: comments,
		Pagination: model.Pagination{
			Page:     param.Page,
			PageSize: param.PageSize,
			Total:    count,
		},
	}, err
}

func (p *GetCommentsParams) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return nil
}
