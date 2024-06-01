package category

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/category"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) CreateCategory(
	ctx context.Context,
	request *blogging.CreateCategoryRequest,
) (*blogging.CreateCategoryResponse, error) {
	_, role, err := jwt.GetIDAndRoleFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if *role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}
	params := category.CreateCategoryParams{
		Name:        request.Name,
		Description: utils.WrapperValueString(request.Description),
	}
	rs, err := g.usecase.CreateCategory.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.CreateCategoryResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.CreateCategoryResponse_Data{
			Category: common.MapToCategoryResponse(&rs.Category),
		},
	}, nil
}
