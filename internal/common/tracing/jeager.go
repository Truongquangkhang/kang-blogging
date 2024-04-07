package tracing

import (
	"context"
	"fmt"
	"kang-edu/common/config"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var serviceName = "unknown-service"

// Setting up Jaeger for `otel` package.
// It first assign arguments to global vars for reuse purpose.
// Then it creates a SpanExporter pointing at the endpoint of external Jaeger collector and resource.
// Then it wraps everything into a Trace Provider, tell `otel` to use it, and
// return a clean-up function.
func Init() func() {
	serviceName = fmt.Sprintf("%s-%s", config.GetNamespace(), config.GetServiceName())

	traceProvider, traceProviderCleanUp := NewTracerProvider()
	otel.SetTracerProvider(traceProvider)
	return traceProviderCleanUp
}

func NewTracerProvider() (
	*trace.TracerProvider,
	func(),
) {
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(newExporter()),
		trace.WithResource(newResource()),
	)
	cleanUp := func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			panic(fmt.Sprintf("Could not gracefully shutdown Jaeger tracing. Err=%v", err))
		}
	}
	return traceProvider, cleanUp
}

func newExporter() trace.SpanExporter {
	collectorEndpoint := os.Getenv("JAEGER_COLLECTOR_ENDPOINT")
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(collectorEndpoint)))
	if err != nil {
		panic(fmt.Sprintf("Setting up Jaeger tracing failed. Could not create new exporter. Err=%v", err))
	}
	return exporter
}

func newResource() *resource.Resource {
	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("Setting up Jaeger tracing failed. Could not create new resource. Err=%v", err))
	}
	return resource
}
