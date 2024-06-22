package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) DeleteComment(
	ctx context.Context,
	request *blogging.DeleteCommentRequest,
) (*blogging.DeleteCommentResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	params := comment.DeleteCommentParams{
		CommentID: request.CommentId,
		UserID:    auth.UserID,
	}
	_, err = g.usecase.DeleteComment.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.DeleteCommentResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
