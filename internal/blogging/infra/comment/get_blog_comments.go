package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	comment2 "kang-blogging/internal/blogging/domain/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) GetBlogComments(
	ctx context.Context,
	request *blogging.GetBlogCommentsRequest,
) (*blogging.GetBlogCommentsResponse, error) {
	params := comment.GetBlogCommentsParams{
		Page:     request.Page,
		PageSize: request.PageSize,
		BlogID:   request.BlogId,
	}
	rs, err := g.usecase.GetBlogComments.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetBlogCommentsResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetBlogCommentsResponse_Data{
			Comments:   mapToGetBlogCommentsResponseData(rs.Comments),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}

func mapToGetBlogCommentsResponseData(comments []comment2.ResultGetBlogComments) []*blogging.CommentWithReplies {
	var responseData []*blogging.CommentWithReplies
	for _, c := range comments {
		var replies []*blogging.Comment
		for _, reply := range c.Replies {
			replies = append(replies, common.MapModelCommentToResponse(reply))
		}
		responseData = append(responseData, &blogging.CommentWithReplies{
			Comment: common.MapModelCommentToResponse(c.Comment),
			Replies: replies,
		})
	}
	return responseData
}
