package user

import (
	"context"
	"kang-blogging/internal/blogging/app/usecase/user"
	"kang-blogging/internal/blogging/infra"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) GetUserDetail(
	ctx context.Context,
	request *blogging.GetUserDetailRequest,
) (*blogging.GetUserDetailResponse, error) {
	params := user.GetUserDetailParams{
		ID: request.UserId,
	}

	rs, err := g.usecase.GetUserDetail.Handle(ctx, params)
	if err != nil {
		return nil, infra.ParseGrpcError(err)
	}

	return &blogging.GetUserDetailResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.GetUserDetailResponse_Data{
			Users: mapUserToUserInfoMetadataResponse(rs.User),
		},
	}, nil
}
