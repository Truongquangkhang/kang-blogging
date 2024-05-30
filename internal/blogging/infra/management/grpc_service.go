package management

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.ManagementUsecase
}

func NewGrpcService(usecase app.ManagementUsecase) GrpcService {
	return GrpcService{usecase}
}
