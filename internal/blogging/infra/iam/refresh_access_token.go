package iam

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) RefreshAccessToken(
	ctx context.Context,
	request *blogging.RefreshAccessTokenRequest,
) (*blogging.RefreshAccessTokenResponse, error) {
	userId, role, err := infra.GetIDAndRoleFromJwtToken(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	params := iam.RefreshAccessTokenParams{
		UserID: *userId,
		Role:   *role,
	}
	rs, err := g.usecase.RefreshAccessToken.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.RefreshAccessTokenResponse{
		Code:    0,
		Message: "Succes",
		Data: &blogging.RefreshAccessTokenResponse_Data{
			AccessToken: rs.AccessToken,
		},
	}, nil
}
