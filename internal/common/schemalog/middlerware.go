package schemalog

import (
	"fmt"
	"kang-blogging/internal/common/server/httpheader"
	"kang-blogging/internal/common/server/httpresp"
	"kang-blogging/internal/common/utils"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchemaLogMiddleware struct {
}

func (t SchemaLogMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(ignoredURL, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		requestEvent := getRequestEventForMiddleware(r)
		logRequestEvent(requestEvent)

		rww := httpresp.NewResponseWriterWrapper(w)
		defer func() {
			responseEvent := getResponseEventForMiddleware(rww, r)
			logResponseEvent(responseEvent)
		}()

		next.ServeHTTP(rww, r)
	})
}

func getRequestEventForMiddleware(r *http.Request) RequestEvent {
	span := trace.SpanFromContext(r.Context())

	params := []string{}
	for key, value := range r.URL.Query() {
		paramAsStr := fmt.Sprintf("%v:[%v]", key, strings.Join(value, ","))
		params = append(params, paramAsStr)
	}

	headers := []string{}
	includedHeaders := []string{
		httpheader.CONTENT_TYPE,
		httpheader.USER_AGENT,
	}
	for _, includedHeader := range includedHeaders {
		headers = append(
			headers,
			fmt.Sprintf("%v:%v", includedHeader, r.Header.Get(includedHeader)),
		)
	}

	return RequestEvent{
		TraceId:   span.SpanContext().TraceID().String(),
		Url:       r.URL.Path,
		Method:    r.Method,
		Params:    params,
		Headers:   headers,
		UserId:    r.Header.Get(httpheader.X_USER_ID),
		Timestamp: timestamppb.New(utils.GetServerNow()),
	}
}

func getResponseEventForMiddleware(w httpresp.ResponseWriterWrapper, r *http.Request) ResponseEvent {
	span := trace.SpanFromContext(r.Context())

	code, message := w.GetBodyData()

	return ResponseEvent{
		TraceId:    span.SpanContext().TraceID().String(),
		Url:        r.URL.Path,
		StatusCode: int32(*w.GetStatusCode()),
		Code:       fmt.Sprint(code),
		Message:    message,
		Timestamp:  timestamppb.New(utils.GetServerNow()),
	}
}
