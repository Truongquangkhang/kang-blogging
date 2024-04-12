package query

import (
	"context"
	"github.com/sirupsen/logrus"
	repository "kang-blogging/internal/blogging/domain"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/logs"
)

type AllApplicableRepositorys struct {
}

type AllApplicableRepositorysHandler decorator.QueryHandler[AllApplicableRepositorys, []repository.Repository]

type allApplicableRepositorysHandler struct {
	repositoryRepo repository.Repository
}

func NewAllApplicableRepositorysHandler(
	repositoryRepo repository.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) AllApplicableRepositorysHandler {
	if repositoryRepo == nil {
		panic("nil repositoryRepo")
	}

	return decorator.ApplyQueryDecorators[AllApplicableRepositorys, []repository.Repository](
		allApplicableRepositorysHandler{
			repositoryRepo: repositoryRepo,
		},
		logger,
		metricsClient,
	)
}

func (h allApplicableRepositorysHandler) Handle(
	ctx context.Context,
	query AllApplicableRepositorys,
) (r []repository.Repository, err error) {
	defer func() {
		logs.LogQueryExecution("AllApplicableRepositorysHandler", query, err)
	}()

	return nil, nil
}
