package user

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) GetUserDetail(
	ctx context.Context,
	request *blogging.GetUserDetailRequest,
) (*blogging.GetUserDetailResponse, error) {
	return &blogging.GetUserDetailResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
