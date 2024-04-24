package blog

import (
	"context"
	"github.com/sirupsen/logrus"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
)

type GetBlogsParams struct {
}

type GetBlogsResult struct {
}

type GetBlogsHandler decorator.UsecaseHandler[GetBlogsParams, GetBlogsResult]

type getBogsHandler struct {
	userRepo user.Repository
	//blogRepo blog.go.Repository
}

func NewGetBlogsHandler(
	userRepo user.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) GetBlogsHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[GetBlogsParams, GetBlogsResult](
		getBogsHandler{
			userRepo: userRepo,
		},
		logger,
		metrics,
	)
}

func (g getBogsHandler) Handle(ctx context.Context, param GetBlogsParams) (GetBlogsResult, error) {
	//TODO implement me
	panic("implement me")
}
