package user

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.UserUsecase
}

func NewGrpcService(usecase app.UserUsecase) GrpcService {
	return GrpcService{usecase}
}
