// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: blogging/blogging_management.proto

package blogging

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BloggingManagementService_GetDashboard_FullMethodName = "/blogging.BloggingManagementService/GetDashboard"
)

// BloggingManagementServiceClient is the client API for BloggingManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BloggingManagementServiceClient interface {
	GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error)
}

type bloggingManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBloggingManagementServiceClient(cc grpc.ClientConnInterface) BloggingManagementServiceClient {
	return &bloggingManagementServiceClient{cc}
}

func (c *bloggingManagementServiceClient) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error) {
	out := new(GetDashboardResponse)
	err := c.cc.Invoke(ctx, BloggingManagementService_GetDashboard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BloggingManagementServiceServer is the server API for BloggingManagementService service.
// All implementations should embed UnimplementedBloggingManagementServiceServer
// for forward compatibility
type BloggingManagementServiceServer interface {
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error)
}

// UnimplementedBloggingManagementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBloggingManagementServiceServer struct {
}

func (UnimplementedBloggingManagementServiceServer) GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}

// UnsafeBloggingManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BloggingManagementServiceServer will
// result in compilation errors.
type UnsafeBloggingManagementServiceServer interface {
	mustEmbedUnimplementedBloggingManagementServiceServer()
}

func RegisterBloggingManagementServiceServer(s grpc.ServiceRegistrar, srv BloggingManagementServiceServer) {
	s.RegisterService(&BloggingManagementService_ServiceDesc, srv)
}

func _BloggingManagementService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloggingManagementServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BloggingManagementService_GetDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloggingManagementServiceServer).GetDashboard(ctx, req.(*GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BloggingManagementService_ServiceDesc is the grpc.ServiceDesc for BloggingManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BloggingManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blogging.BloggingManagementService",
	HandlerType: (*BloggingManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDashboard",
			Handler:    _BloggingManagementService_GetDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogging/blogging_management.proto",
}
