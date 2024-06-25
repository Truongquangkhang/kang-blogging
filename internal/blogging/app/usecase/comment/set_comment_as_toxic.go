package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/detection_client"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/utils"
)

type SetCommentAsToxicParams struct {
	CommentID    string
	Content      string
	ToxicIndexes []int32
}

type SetCommentAsToxicResult struct {
}

type SetCommentAsToxicHandler decorator.UsecaseHandler[SetCommentAsToxicParams, SetCommentAsToxicResult]

type setCommentAsToxicHandler struct {
	commentRepo comment.Repository
}

func NewSetCommentAsToxicHandler(
	commentRepo comment.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) SetCommentAsToxicHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[SetCommentAsToxicParams, SetCommentAsToxicResult](
		setCommentAsToxicHandler{
			commentRepo: commentRepo,
		},
		logger,
		metrics,
	)
}

func (g setCommentAsToxicHandler) Handle(ctx context.Context, param SetCommentAsToxicParams) (SetCommentAsToxicResult, error) {
	err := param.Validate()
	if err != nil {
		return SetCommentAsToxicResult{}, err
	}
	cmt, err := g.commentRepo.GetCommentById(ctx, param.CommentID)
	if err != nil {
		return SetCommentAsToxicResult{}, err
	}
	if cmt == nil {
		return SetCommentAsToxicResult{}, errors.NewNotFoundError("comment not found")
	}
	isToxic := false
	for _, toxicIndex := range param.ToxicIndexes {
		if toxicIndex == 1 {
			isToxic = true
			break
		}
	}
	predictionComment := detection_client.PredictionComment{
		Comment:                param.Content,
		IsToxicComment:         isToxic,
		ToxicPredictionComment: param.ToxicIndexes,
	}
	marshall, err := json.Marshal(predictionComment)
	if err != nil {
		return SetCommentAsToxicResult{}, err
	}
	cmt.Prediction = utils.ToStringPointerValue(fmt.Sprintf("%s", string(marshall)))
	cmt.IsToxicity = isToxic

	cmt, err = g.commentRepo.UpdateComment(ctx, cmt)
	if err != nil {
		return SetCommentAsToxicResult{}, err
	}

	return SetCommentAsToxicResult{}, err
}

func (p *SetCommentAsToxicParams) Validate() error {
	if p.CommentID == "" || p.Content == "" || len(p.ToxicIndexes) == 0 {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}
