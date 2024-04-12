package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func ApplyCommandDecorators[C any](
	handler CommandHandler[C],
	logger *logrus.Entry,
	metricsClient MetricsClient,
) CommandHandler[C] {
	return commandLoggingDecorator[C]{
		base: commandMetricsDecorator[C]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
