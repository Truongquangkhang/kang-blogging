package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
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

	return &blogging.GetBlogDetailResponse{
		Code:    0,
		Message: "Success",
		Data:    buildGetBlogDetailResponseData(rs),
	}, nil
}

func buildGetBlogDetailResponseData(rs blog.GetBlogDetailResult) *blogging.GetBlogDetailResponse_Data {
	return &blogging.GetBlogDetailResponse_Data{
		Blog: &blogging.BlogInfo{
			BlogInfo: common.MapBlogModelToBlogMetadataResponse(rs.Blog),
			Content:  utils.WrapperStringFromString(rs.Blog.Content),
		},
	}
}
