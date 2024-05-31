package blog

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) UpdateBlogDetail(
	ctx context.Context,
	request *blogging.UpdateBlogDetailRequest,
) (*blogging.UpdateBlogDetailResponse, error) {
	authorId, _, err := jwt.GetIDAndRoleFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	param := blog.UpdateBlogDetailParams{
		AuthorID:    *authorId,
		BlogID:      request.BlogId,
		Name:        utils.WrapperValueString(request.Name),
		Description: utils.WrapperValueString(request.Description),
		Thumbnail:   utils.WrapperValueString(request.Thumbnail),
		Content:     utils.WrapperValueString(request.Content),
		CategoryIDs: request.CategoryIds,
		Published:   utils.WrapperValueBool(request.Published),
	}

	rs, err := g.usecase.UpdateBlogDetail.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdateBlogDetailResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UpdateBlogDetailResponse_Data{
			Blog: common.MapToBlogInfoResponse(rs.Blog),
		},
	}, nil
}
