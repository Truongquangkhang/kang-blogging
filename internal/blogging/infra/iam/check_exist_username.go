package iam

import (
	"context"
	iamUsecase "kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) CheckExistUsername(
	ctx context.Context,
	request *blogging.CheckExistUsernameRequest,
) (*blogging.CheckExistUsernameResponse, error) {
	param := iamUsecase.CheckExistUsernameParam{
		Username: request.Username,
	}
	result, err := g.usecase.CheckExistUsername.Handle(ctx, param)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.CheckExistUsernameResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.CheckExistUsernameResponse_Data{
			AlreadyExist: result.AlreadyExists,
		},
	}, nil
}
