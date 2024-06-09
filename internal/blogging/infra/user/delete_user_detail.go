package user

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
)

func (g GrpcService) DeleteUserDetail(
	ctx context.Context,
	request *blogging.DeleteUserDetailRequest,
) (*blogging.DeleteUserDetailResponse, error) {
	_, role, err := jwt.GetIDAndRoleFromRequest(ctx)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	if *role != constants.ADMIN_ROLE {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	params := user.DeleteUserDetailParams{
		UserID: request.UserId,
	}

	_, err = g.usecase.DeleteUserDetail.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}
	return &blogging.DeleteUserDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
