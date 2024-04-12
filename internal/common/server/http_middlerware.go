package server

import (
	"kang-blogging/internal/common/schemalog"
	"kang-blogging/internal/common/tracing"
	"net/http"
)

func AddSchemaLogMiddleware(next http.Handler) http.Handler {
	return schemalog.SchemaLogMiddleware{}.Middleware(next)
}

func AddTracingMiddleware(next http.Handler) http.Handler {
	return tracing.TracingHttpMiddleware{}.Middleware(next)
}
