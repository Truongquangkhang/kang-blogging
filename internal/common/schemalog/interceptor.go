package schemalog

import (
	"context"
	"fmt"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/server/grpcheader"
	"kang-blogging/internal/common/utils"
	"reflect"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UnaryServerInterceptor(
	ctx context.Context,
	request interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	grpcMethod := info.FullMethod
	if slices.Contains(ignoredURL, grpcMethod) {
		response, err := handler(ctx, request)
		return response, err
	}
	requestEvent := getRequestEventForUnaryServerInterceptor(ctx, grpcMethod)
	logRequestEvent(requestEvent)

	response, err := handler(ctx, request)

	var responseEvent ResponseEvent
	if err == nil {
		responseEvent = getResponseEventForUnaryServerInterceptor(ctx, grpcMethod, response)
	} else {
		responseEvent = getErrorResponseEventForUnaryServerInterceptor(ctx, grpcMethod, err)
	}
	logResponseEvent(responseEvent)

	return response, err
}

func getRequestEventForUnaryServerInterceptor(
	ctx context.Context,
	grpcMethod string,
	// Uncomment when use func getParams
	// request interface{},
) RequestEvent {
	span := trace.SpanFromContext(ctx)

	// params := getParams(request)
	headers := getHeaders(ctx)

	return RequestEvent{
		TraceId: span.SpanContext().SpanID().String(),
		Url:     grpcMethod,
		// Params:    params,
		Headers:   headers,
		Timestamp: timestamppb.New(utils.GetServerNow()),
	}
}

func getHeaders(ctx context.Context) []string {
	mtd, err := metadata.FromIncomingContext(ctx)
	if !err {
		return []string{}
	}

	headers := []string{}
	includedHeaders := []string{
		grpcheader.CONTENT_TYPE,
		grpcheader.USER_AGENT,
		grpcheader.GRPC_CONTENT_TYPE,
		grpcheader.GRPC_USER_AGENT,
	}
	for _, includedHeader := range includedHeaders {
		header := mtd.Get(includedHeader)
		if len(header) > 0 {
			headers = append(
				headers,
				fmt.Sprintf("%v:%v", includedHeader, header),
			)
		}
	}
	return headers
}

// This func use to get params from request
// but it may get data from request body
// So it's commented out
// func getParams(request interface{}) []string {
// 	value := reflect.ValueOf(request)
// 	vtype := value.Kind()
// 	if vtype != reflect.Ptr {
// 		return []string{}
// 	}

// 	derefValue := value.Elem()
// 	if derefValue.Kind() != reflect.Struct {
// 		return []string{}
// 	}

// 	var params []string
// 	derefType := derefValue.Type()
// 	for index := 0; index < derefType.NumField(); index++ {
// 		fieldType := derefType.Field(index)
// 		fieldValue := derefValue.Field(index)
// 		if fieldType.IsExported() {
// 			params = append(params,
// 				fmt.Sprintf("%v:%v",
// 					fieldType.Name,
// 					fieldValue,
// 				))
// 		}
// 	}
// 	return params
// }

func getResponseEventForUnaryServerInterceptor(
	ctx context.Context,
	grpcMethod string,
	response interface{},
) ResponseEvent {
	span := trace.SpanFromContext(ctx)

	code, message := getResponseCodeAndMessage(response)

	return ResponseEvent{
		TraceId:    span.SpanContext().SpanID().String(),
		Url:        grpcMethod,
		StatusCode: 200,
		Code:       fmt.Sprint(code),
		Message:    message,
		Timestamp:  timestamppb.New(utils.GetServerNow()),
	}
}

func getResponseCodeAndMessage(response interface{}) (int32, string) {
	value := reflect.ValueOf(response)
	vtype := value.Kind()
	if vtype != reflect.Ptr {
		return -1, fmt.Sprintf("Type %v is not supported response", vtype)
	}

	derefValue := value.Elem()
	if derefValue.Kind() != reflect.Struct {
		return -1, fmt.Sprintf("Type %v is not supported response", derefValue.Kind())
	}

	code := int32(derefValue.FieldByName("Code").Int())
	message := derefValue.FieldByName("Message").String()
	return code, message
}

func getErrorResponseEventForUnaryServerInterceptor(
	ctx context.Context,
	grpcMethod string,
	err error,
) ResponseEvent {
	span := trace.SpanFromContext(ctx)
	baseErr := errors.ParseGrpcError(err)
	code, message := baseErr.ErrorCode(), baseErr.ErrorMessage()

	return ResponseEvent{
		TraceId:    span.SpanContext().SpanID().String(),
		Url:        grpcMethod,
		StatusCode: code / 1000,
		Code:       fmt.Sprint(code),
		Message:    message,
		Timestamp:  timestamppb.New(utils.GetServerNow()),
	}
}

func StreamServerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	wrapped := grpc_middleware.WrapServerStream(ss)
	err := handler(srv, wrapped)
	return err
}
