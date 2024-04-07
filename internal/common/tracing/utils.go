package tracing

import (
	"context"
	"kang-edu/common/errors"
	"kang-edu/common/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Receives a span and error instance. If error is not nil,
// update span to have status error and record the error.
// Otherwise do nothing.
func UpdateSpanError(span trace.Span, err error) {
	if err == nil {
		return
	}
	if baseErr, ok := err.(errors.BaseError); ok {
		span.SetStatus(codes.Error, baseErr.ErrorMessage())
		span.RecordError(baseErr)
		return
	}
	span.SetStatus(codes.Error, err.Error())
	span.RecordError(err)
}

const handlerDepth = 3

// A shortcut function to create a span from `context` and `spanName`
// and start it immediately.
func StartSpan(ctx context.Context) (
	context.Context,
	trace.Span,
) {
	funcName := utils.GetFuncNameFromCallStack(handlerDepth)
	return StartSpanWithName(ctx, funcName)
}

func StartSpanWithName(ctx context.Context, spanName string) (
	context.Context,
	trace.Span,
) {
	return otel.Tracer(serviceName).Start(ctx, spanName)
}
