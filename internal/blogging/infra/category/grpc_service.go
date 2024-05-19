package category

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.CategoryUsecase
}

func NewGrpcService(usecase app.CategoryUsecase) GrpcService {
	return GrpcService{usecase}
}
