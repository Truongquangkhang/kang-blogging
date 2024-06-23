package iam

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/iam"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) ChangePassword(
	ctx context.Context,
	request *blogging.ChangePasswordRequest,
) (*blogging.ChangePasswordResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil || auth == nil {
		return nil, infra.ParseGrpcError(err)
	}

	if request.UserId != auth.UserID && auth.Role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	params := iam.ChangePasswordParams{
		UserID:             request.UserId,
		OldPassword:        utils.WrapperValueString(request.OldPassword),
		NewPassword:        request.NewPassword,
		WithoutOldPassword: auth.Role == constants.ADMIN_ROLE,
	}
	_, err = g.usecase.ChangePassword.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.ChangePasswordResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
