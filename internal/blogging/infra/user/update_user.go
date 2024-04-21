package user

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) UpdateUserDetail(
	ctx context.Context,
	request *blogging.UpdateUserDetailRequest,
) (*blogging.UpdateUserDetailResponse, error) {
	return &blogging.UpdateUserDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
