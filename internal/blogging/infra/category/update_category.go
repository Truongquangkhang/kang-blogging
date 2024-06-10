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

func (g GrpcService) UpdateCategory(
	ctx context.Context,
	request *blogging.UpdateCategoryRequest,
) (*blogging.UpdateCategoryResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}
	params := category.UpdateCategoryParams{
		CategoryID:  request.CategoryId,
		Name:        utils.WrapperValueString(request.Name),
		Description: utils.WrapperValueString(request.Description),
	}
	rs, err := g.usecase.UpdateCategory.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.UpdateCategoryResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UpdateCategoryResponse_Data{
			Category: common.MapToCategoryResponse(&rs.Category),
		},
	}, nil
}
