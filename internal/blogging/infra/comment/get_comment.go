package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetComment(
	ctx context.Context,
	request *blogging.GetCommentRequest,
) (*blogging.GetCommentResponse, error) {
	params := comment.GetCommentParams{
		CommentID: request.CommentId,
	}
	rs, err := g.usecase.GetComment.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	var commentProcessed *string
	var prediction []int32
	if rs.PredictionComment != nil {
		commentProcessed = &rs.PredictionComment.Comment
		prediction = rs.PredictionComment.ToxicPredictionComment
	}

	return &blogging.GetCommentResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetCommentResponse_Data{
			Comment:          common.MapModelCommentToResponse(rs.Comment),
			ContentProcessed: utils.WrapperStringFromString(commentProcessed),
			Predictions:      prediction,
		},
	}, nil
}
