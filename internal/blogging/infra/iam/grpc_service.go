package iam

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.IAMUsecases
}

func NewGrpcService(usecase app.IAMUsecases) GrpcService {
	return GrpcService{usecase}
}
