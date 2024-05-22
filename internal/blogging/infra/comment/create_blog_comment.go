package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) CreateBlogComment(
	ctx context.Context,
	request *blogging.CreateBlogCommentsRequest,
) (*blogging.CreateBlogCommentsResponse, error) {
	userId, _, err := infra.GetIDAndRoleFromJwtToken(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	param := comment.CreateBlogCommentParams{
		UserID:         *userId,
		BlogID:         request.BlogId,
		Content:        request.Content,
		ReplyCommentID: utils.WrapperValueString(request.ReplyCommentId),
	}
	rs, err := g.usecase.CreateBlogComment.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.CreateBlogCommentsResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.CreateBlogCommentsResponse_Data{
			Comment: common.MapModelCommentToResponse(rs.Comment),
		},
	}, nil
}
