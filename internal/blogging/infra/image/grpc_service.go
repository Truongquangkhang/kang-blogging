package image

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.ImageUsecase
}

func NewGrpcService(usecase app.ImageUsecase) GrpcService {
	return GrpcService{usecase}
}
