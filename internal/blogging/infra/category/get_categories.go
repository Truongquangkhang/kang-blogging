package category

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/category"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) GetCategories(
	ctx context.Context,
	request *blogging.GetCategoriesRequest,
) (*blogging.GetCategoriesResponse, error) {
	param := category.GetCategoriesParams{
		Page:       request.Page,
		PageSize:   request.PageSize,
		SortBy:     utils.WrapperValueString(request.SortBy),
		SearchName: utils.WrapperValueString(request.SearchName),
	}

	rs, err := g.usecase.GetCategories.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetCategoriesResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetCategoriesResponse_Data{
			Categories: common.MapToCategoriesMetadata(rs.Categories),
			Pagination: common.MapToPaginationResponse(rs.Pagination),
		},
	}, nil
}
