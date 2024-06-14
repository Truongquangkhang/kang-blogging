package blog

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/blog"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetBlogs(
	ctx context.Context,
	request *blogging.GetBlogsRequest,
) (*blogging.GetBlogsResponse, error) {
	param := blog.GetBlogsParams{
		Page:         request.Page,
		PageSize:     request.PageSize,
		SearchBy:     utils.WrapperValueString(request.SearchBy),
		SearchName:   utils.WrapperValueString(request.SearchName),
		AuthorIds:    utils.WrapperValueString(request.AuthorIds),
		CategoryIds:  utils.WrapperValueString(request.CategoryIds),
		SortBy:       utils.WrapperValueString(request.SortBy),
		IsDeprecated: utils.WrapperValueBool(request.IsDeprecated),
		Published:    utils.WrapperValueBool(request.Published),
	}
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err == nil && auth != nil {
		param.CurrentUserID = utils.ToStringPointerValue(auth.UserID)
		param.GetBlogRelevant = utils.WrapperValueBool(request.GetRelevant)
	}
	rs, err := g.usecase.GetBlogs.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.GetBlogsResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetBlogsResponse_Data{
			Blogs:      common.MapToListBlogMetadataResponse(rs.Blogs),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}
