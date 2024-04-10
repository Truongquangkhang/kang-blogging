package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

func ApplyQueryDecorators[Q any, R any](
	handler QueryHandler[Q, R],
	logger *logrus.Entry,
	metricsClient MetricsClient,
) QueryHandler[Q, R] {
	return queryLoggingDecorator[Q, R]{
		base: queryMetricsDecorator[Q, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}
