package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) SetCommentAsToxic(
	ctx context.Context,
	request *blogging.SetCommentAsToxicRequest,
) (*blogging.SetCommentAsToxicResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}
	params := comment.SetCommentAsToxicParams{
		CommentID:    request.CommentId,
		Content:      request.Content,
		ToxicIndexes: request.ToxicIndexes,
	}
	_, err = g.usecase.SetCommentAsToxic.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.SetCommentAsToxicResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
