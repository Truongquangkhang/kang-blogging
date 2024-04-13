package infra

import (
	"context"
	"kang-blogging/internal/blogging/app"
	"kang-blogging/internal/blogging/infra/genproto/blogging"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app: app}
}

func (g GrpcServer) Login(
	ctx context.Context,
	request *blogging.LoginRequest,
) (*blogging.LoginResponse, error) {
	return &blogging.LoginResponse{
		Code:    0,
		Message: "Success",
		Data: &blogging.LoginResponse_Data{
			AccessToken:  "abc",
			RefreshToken: "def",
			UserInfo: &blogging.UserInfo{
				Id:        "id",
				Name:      "Quang Khang",
				Avatar:    "Image",
				TotalBlog: 2,
			},
		},
	}, nil
}

func (g GrpcServer) Register(
	ctx context.Context,
	request *blogging.RegisterRequest,
) (*blogging.RegisterResponse, error) {
	return &blogging.RegisterResponse{}, nil
}
