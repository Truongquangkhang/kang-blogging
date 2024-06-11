package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/common"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/constants"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/jwt"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) UpdateUserDetail(
	ctx context.Context,
	request *blogging.UpdateUserDetailRequest,
) (*blogging.UpdateUserDetailResponse, error) {
	auth, err := jwt.GetAuthenticationFromRequest(ctx)
	if err != nil || auth == nil {
		return nil, infra.ParseGrpcError(err)
	}
	if auth.Role != constants.ADMIN_ROLE && auth.UserID != request.UserId {
		return nil, infra.ParseGrpcError(errors.NewForbiddenDefaultError())
	}

	params := user.UpdateUserParams{
		ID:          request.UserId,
		Email:       utils.WrapperValueString(request.Email),
		Avatar:      utils.WrapperValueString(request.Avatar),
		DisplayName: utils.WrapperValueString(request.DisplayName),
		Name:        utils.WrapperValueString(request.Name),
		PhoneNumber: utils.WrapperValueString(request.PhoneNumber),
		Gender:      utils.WrapperValueBool(request.Gender),
		Description: utils.WrapperValueString(request.Description),
	}
	rs, err := g.usecase.UpdateUser.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdateUserDetailResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UpdateUserDetailResponse_Data{
			User: common.MapUserToUserInfoMetadataResponse(rs.User),
		},
	}, nil
}
