package tracing

import (
	"fmt"
	"net/http"
)

type TracingHttpMiddleware struct {
}

// Tracing Middleware will starts a span as soon as request hits the server.
// Name of the span is the request signature (Method + URL Path)
func (t TracingHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		newCtx, span := StartSpanWithName(ctx, requestSignatureFromHttpRequest(r))
		defer span.End()

		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

func requestSignatureFromHttpRequest(r *http.Request) string {
	return fmt.Sprintf("%s %s",
		r.Method,
		r.URL.Path,
	)
}
