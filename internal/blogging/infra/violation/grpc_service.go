package violation

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.ViolationUsecase
}

func NewGrpcService(usecase app.ViolationUsecase) GrpcService {
	return GrpcService{usecase}
}
