package management

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/management"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) GetPolicies(
	ctx context.Context,
	request *blogging.GetPoliciesRequest,
) (*blogging.GetPoliciesResponse, error) {
	_, role, err := jwt.GetIDAndRoleFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if *role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	result, err := g.usecase.GetPolicies.Handle(ctx, management.GetPoliciesParams{})
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetPoliciesResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetPoliciesResponse_Data{
			Policies: common.MapToPoliciesResponse(result.Policies),
		},
	}, nil
}
