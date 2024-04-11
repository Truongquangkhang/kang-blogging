package command

import (
	"context"

	"github.com/sirupsen/logrus"
)

type DoSomething struct{}

type DoSomethingHandler decorator.CommandHandler[DoSomething]

type doSomethingHandler struct {
	voucherRepo voucher.Repository
}

func NewDoSomethingHandler(
	voucherRepo voucher.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) decorator.CommandHandler[DoSomething] {
	if voucherRepo == nil {
		panic("nil voucherRepo")
	}

	return decorator.ApplyCommandDecorators[DoSomething](
		doSomethingHandler{
			voucherRepo: voucherRepo,
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
