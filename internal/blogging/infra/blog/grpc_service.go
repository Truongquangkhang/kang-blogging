package blog

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.BlogUsecase
}

func NewGrpcService(usecase app.BlogUsecase) GrpcService {
	return GrpcService{usecase}
}
