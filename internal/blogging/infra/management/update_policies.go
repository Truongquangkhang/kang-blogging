package management

import (
	"fmt"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/management"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/model"
	"strconv"
)

func (g GrpcService) UpdatePolicies(
	ctx context.Context,
	request *blogging.UpdatePoliciesRequest,
) (*blogging.UpdatePoliciesResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil || auth == nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	policies := []model.Policy{}
	for _, policy := range request.Policies {
		if policy.Name == "" {
			return nil, infra.ParseGrpcError(
				errors.NewBadRequestError(fmt.Sprintf("invalid policy name with name")))
		}
		value, err := strconv.ParseInt(policy.Value, 10, 64)
		if err != nil {
			return nil, infra.ParseGrpcError(
				errors.NewBadRequestError(fmt.Sprintf("invalid policy value with name %s", policy.Name)))
		}
		policies = append(policies, model.Policy{
			Type:  &policy.Name,
			Value: &value,
		})
	}
	params := management.UpdatePoliciesParams{
		Policies: policies,
	}
	_, err = g.usecase.UpdatePolicies.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdatePoliciesResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
