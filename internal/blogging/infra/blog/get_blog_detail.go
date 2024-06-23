package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetBlogDetail(
	ctx context.Context,
	request *blogging.GetBlogDetailRequest,
) (*blogging.GetBlogDetailResponse, error) {
	param := blog.GetBlogDetailParams{
		BlogID: request.BlogId,
	}
	rs, err := g.usecase.GetBlogDetail.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	canEdit := false
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err == nil && auth != nil {
		if auth.Role == constants.ADMIN_ROLE {
			canEdit = true
		} else {
			canEdit = auth.UserID == rs.Blog.User.ID
		}
	}
	return &blogging.GetBlogDetailResponse{
		Code:    0,
		Message: "Success",
		Data:    buildGetBlogDetailResponseData(rs, canEdit),
	}, nil
}

func buildGetBlogDetailResponseData(rs blog.GetBlogDetailResult, canEdit bool) *blogging.GetBlogDetailResponse_Data {
	return &blogging.GetBlogDetailResponse_Data{
		Blog: &blogging.BlogInfo{
			BlogInfo: common.MapBlogModelToBlogMetadataResponse(rs.Blog),
			Content:  utils.WrapperStringFromString(rs.Blog.Content),
			CanEdit:  canEdit,
		},
	}
}
