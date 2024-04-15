package iam

import (
	"context"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

func (g GrpcService) Login(
	ctx context.Context,
	request *blogging.LoginRequest,
) (*blogging.LoginResponse, error) {
	return &blogging.LoginResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
