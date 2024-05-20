package comment

import "kang-blogging/internal/blogging/app"

type GrpcService struct {
	usecase app.CommentUsecase
}

func NewGrpcService(usecase app.CommentUsecase) GrpcService {
	return GrpcService{usecase}
}
