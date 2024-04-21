package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
	"kang-blogging/internal/common/utils"
)

func (g GrpcService) UpdateUserDetail(
	ctx context.Context,
	request *blogging.UpdateUserDetailRequest,
) (*blogging.UpdateUserDetailResponse, error) {
	params := user.UpdateUserParams{
		ID:          request.UserId,
		Email:       utils.WrapperValueString(request.Email),
		Avatar:      utils.WrapperValueString(request.Avatar),
		DisplayName: utils.WrapperValueString(request.DisplayName),
		Name:        utils.WrapperValueString(request.Name),
		PhoneNumber: utils.WrapperValueString(request.PhoneNumber),
		Gender:      utils.WrapperValueBool(request.Gender),
	}
	rs, err := g.usecase.UpdateUser.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.UpdateUserDetailResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.UpdateUserDetailResponse_Data{
			Users: mapUserToUserInfoMetadataResponse(rs.User),
		},
	}, nil
}
