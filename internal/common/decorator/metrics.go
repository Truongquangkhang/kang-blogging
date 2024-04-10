package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type commandMetricsDecorator[C any] struct {
	base   CommandHandler[C]
	client MetricsClient
}

func (d commandMetricsDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(cmd))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("commands.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("commands.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("commands.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryMetricsDecorator[Q any, R any] struct {
	base   QueryHandler[Q, R]
	client MetricsClient
}

func (d queryMetricsDecorator[Q, R]) Handle(ctx context.Context, query Q) (result R, err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(query))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("queries.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("queries.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("queries.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, query)
}

type usecaseMetricsDecorator[P any, R any] struct {
	base   QueryHandler[P, R]
	client MetricsClient
}

func (d usecaseMetricsDecorator[P, R]) Handle(ctx context.Context, param P) (result R, err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(param))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("usecases.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("usecases.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("usecases.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, param)
}
