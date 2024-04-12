package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

func ApplyUsecaseDecorators[P any, R any](
	handler UsecaseHandler[P, R],
	logger *logrus.Entry,
	metricsClient MetricsClient,
) UsecaseHandler[P, R] {
	return usecaseLoggingDecorator[P, R]{
		base: usecaseMetricsDecorator[P, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type UsecaseHandler[P any, R any] interface {
	Handle(ctx context.Context, param P) (R, error)
}
