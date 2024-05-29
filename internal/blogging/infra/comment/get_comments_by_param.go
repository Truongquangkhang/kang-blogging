package comment

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/comment"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/model"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetCommentsByParam(
	ctx context.Context,
	request *blogging.GetCommentsByParamRequest,
) (*blogging.GetCommentsByParamResponse, error) {
	params := comment.GetCommentsParams{
		Page:       request.Page,
		PageSize:   request.PageSize,
		SearchName: utils.WrapperValueString(request.SearchName),
		SortBy:     utils.WrapperValueString(request.SortBy),
		IsToxicity: utils.WrapperValueBool(request.IsToxicity),
		UserIds:    utils.WrapperValueString(request.UserIds),
	}

	result, err := g.usecase.GetComments.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetCommentsByParamResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetCommentsByParamResponse_Data{
			Comments:   buildCommentItemResponse(result.Comments),
			Pagination: common.MapToPaginationResponse(result.Pagination),
		},
	}, nil
}

func buildCommentItemResponse(comments []model.Comment) []*blogging.GetCommentsByParamResponse_CommentItem {
	var rs []*blogging.GetCommentsByParamResponse_CommentItem
	for _, cm := range comments {
		rs = append(rs, &blogging.GetCommentsByParamResponse_CommentItem{
			CommentInfo:    common.MapModelCommentToResponse(cm),
			ReplyCommentId: utils.WrapperStringFromString(cm.ReplyCommentID),
			BlogId:         cm.BlogID,
		})
	}
	return rs
}
