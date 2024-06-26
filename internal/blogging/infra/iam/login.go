package iam

import (
	"context"
	iamUsecase "kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) Login(
	ctx context.Context,
	request *blogging.LoginRequest,
) (*blogging.LoginResponse, error) {
	param := iamUsecase.LoginParams{
		Username: request.Username,
		Password: request.Password,
	}
	rs, err := g.usecase.Login.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.LoginResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.LoginResponse_Data{
			AccessToken:  rs.AccessToken,
			RefreshToken: rs.RefreshToken,
			UserInfo:     common.MapUserToUserInfoMetadataResponse(rs.UserInfo),
		},
	}, nil
}
