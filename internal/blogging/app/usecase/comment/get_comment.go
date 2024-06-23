package comment

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/domain/detection_client"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

type GetCommentParams struct {
	CommentID string
}

type GetCommentResult struct {
	Comment           model.Comment
	PredictionComment *detection_client.PredictionComment
}

type GetCommentHandler decorator.UsecaseHandler[GetCommentParams, GetCommentResult]

type getCommentHandler struct {
	commentRepo comment.Repository
	userRepo    user.Repository
}

func NewGetCommentHandler(
	commentRepo comment.Repository,
	userRepo user.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetCommentHandler {
	if commentRepo == nil {
		panic("commentRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetCommentParams, GetCommentResult](
		getCommentHandler{
			commentRepo: commentRepo,
			userRepo:    userRepo,
		},
		logger,
		metrics,
	)
}

func (g getCommentHandler) Handle(ctx context.Context, param GetCommentParams) (GetCommentResult, error) {
	err := param.Validate()
	if err != nil {
		return GetCommentResult{}, err
	}

	cmt, err := g.commentRepo.GetCommentById(ctx, param.CommentID)
	if err != nil {
		return GetCommentResult{}, err
	}
	if cmt == nil {
		return GetCommentResult{}, errors.NewNotFoundError("comment not found")
	}
	var prediction *detection_client.PredictionComment
	if cmt.Prediction != nil {
		err = json.Unmarshal([]byte(utils.ToStringValue(cmt.Prediction)), &prediction)
		if err != nil {
			return GetCommentResult{}, err
		}
	}

	u, err := g.userRepo.GetUserByID(ctx, cmt.UserID)
	if err != nil {
		return GetCommentResult{}, err
	}
	if u == nil {
		return GetCommentResult{}, errors.NewNotFoundError("user not found")
	}
	cmt.User = *u

	return GetCommentResult{
		Comment:           *cmt,
		PredictionComment: prediction,
	}, err
}

func (p *GetCommentParams) Validate() error {
	if p.CommentID == "" {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}
