package command

import (
	"context"
	"github.com/sirupsen/logrus"
	repository "kang-blogging/internal/blogging/domain"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/logs"
)

type DoSomething struct{}

type DoSomethingHandler decorator.CommandHandler[DoSomething]

type doSomethingHandler struct {
	repositoryRepo repository.Repository
}

func NewDoSomethingHandler(
	repositoryRepo repository.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) decorator.CommandHandler[DoSomething] {
	if repositoryRepo == nil {
		panic("nil repositoryRepo")
	}

	return decorator.ApplyCommandDecorators[DoSomething](
		doSomethingHandler{
			repositoryRepo: repositoryRepo,
		},
		logger,
		metricsClient,
	)
}

func (h doSomethingHandler) Handle(
	ctx context.Context,
	cmd DoSomething,
) (err error) {
	defer func() {
		logs.LogCommandExecution("DoSomethingHandler", cmd, err)
	}()

	// Do something

	return nil
}
