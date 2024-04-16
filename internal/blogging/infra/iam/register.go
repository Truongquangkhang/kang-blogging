package iam

import (
	"context"
	iamUsecase "kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) Register(
	ctx context.Context,
	request *blogging.RegisterRequest,
) (*blogging.RegisterResponse, error) {
	params := iamUsecase.RegisterParams{
		Username:     request.Username,
		Password:     request.Password,
		Name:         request.Username,
		DisplayName:  request.DisplayName,
		Email:        request.Email,
		Gender:       utils.WrapperValueBool(request.Gender),
		Avatar:       utils.WrapperValueString(request.Avatar),
		BirthOfDay:   utils.WrapperValueInt64(request.BirthOfDay),
		PhoneNumbers: utils.WrapperValueString(request.PhoneNumbers),
	}
	_, err := g.usecase.Register.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.RegisterResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
