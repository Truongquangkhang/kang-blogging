package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) UpdateComment(
	ctx context.Context,
	request *blogging.UpdateCommentRequest,
) (*blogging.UpdateCommentResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	params := comment.UpdateCommentParams{
		CommentID: request.CommentId,
		UserID:    auth.UserID,
		Content:   request.Content,
	}

	rs, err := g.usecase.UpdateComment.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdateCommentResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UpdateCommentResponse_Data{
			Comment: common.MapModelCommentToResponse(rs.Comment),
		},
	}, nil
}
