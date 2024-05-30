// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: blogging/blogging_management.proto

/*
Package blogging is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package blogging

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_BloggingManagementService_GetDashboard_0(ctx context.Context, marshaler runtime.Marshaler, client BloggingManagementServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetDashboardRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetDashboard(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BloggingManagementService_GetDashboard_0(ctx context.Context, marshaler runtime.Marshaler, server BloggingManagementServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetDashboardRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetDashboard(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBloggingManagementServiceHandlerServer registers the http handlers for service BloggingManagementService to "mux".
// UnaryRPC     :call BloggingManagementServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBloggingManagementServiceHandlerFromEndpoint instead.
func RegisterBloggingManagementServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BloggingManagementServiceServer) error {

	mux.Handle("GET", pattern_BloggingManagementService_GetDashboard_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/blogging.BloggingManagementService/GetDashboard", runtime.WithHTTPPathPattern("/api/v1/management/dashboard"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BloggingManagementService_GetDashboard_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BloggingManagementService_GetDashboard_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBloggingManagementServiceHandlerFromEndpoint is same as RegisterBloggingManagementServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBloggingManagementServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBloggingManagementServiceHandler(ctx, mux, conn)
}

// RegisterBloggingManagementServiceHandler registers the http handlers for service BloggingManagementService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBloggingManagementServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBloggingManagementServiceHandlerClient(ctx, mux, NewBloggingManagementServiceClient(conn))
}

// RegisterBloggingManagementServiceHandlerClient registers the http handlers for service BloggingManagementService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BloggingManagementServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BloggingManagementServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BloggingManagementServiceClient" to call the correct interceptors.
func RegisterBloggingManagementServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BloggingManagementServiceClient) error {

	mux.Handle("GET", pattern_BloggingManagementService_GetDashboard_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/blogging.BloggingManagementService/GetDashboard", runtime.WithHTTPPathPattern("/api/v1/management/dashboard"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BloggingManagementService_GetDashboard_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BloggingManagementService_GetDashboard_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BloggingManagementService_GetDashboard_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "management", "dashboard"}, ""))
)

var (
	forward_BloggingManagementService_GetDashboard_0 = runtime.ForwardResponseMessage
)
