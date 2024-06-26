// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: blogging/violation.proto

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
	ViolationService_GetViolations_FullMethodName = "/blogging.ViolationService/GetViolations"
)

// ViolationServiceClient is the client API for ViolationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViolationServiceClient interface {
	GetViolations(ctx context.Context, in *GetViolationsRequest, opts ...grpc.CallOption) (*GetViolationsResponse, error)
}

type violationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViolationServiceClient(cc grpc.ClientConnInterface) ViolationServiceClient {
	return &violationServiceClient{cc}
}

func (c *violationServiceClient) GetViolations(ctx context.Context, in *GetViolationsRequest, opts ...grpc.CallOption) (*GetViolationsResponse, error) {
	out := new(GetViolationsResponse)
	err := c.cc.Invoke(ctx, ViolationService_GetViolations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViolationServiceServer is the server API for ViolationService service.
// All implementations should embed UnimplementedViolationServiceServer
// for forward compatibility
type ViolationServiceServer interface {
	GetViolations(context.Context, *GetViolationsRequest) (*GetViolationsResponse, error)
}

// UnimplementedViolationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedViolationServiceServer struct {
}

func (UnimplementedViolationServiceServer) GetViolations(context.Context, *GetViolationsRequest) (*GetViolationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViolations not implemented")
}

// UnsafeViolationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViolationServiceServer will
// result in compilation errors.
type UnsafeViolationServiceServer interface {
	mustEmbedUnimplementedViolationServiceServer()
}

func RegisterViolationServiceServer(s grpc.ServiceRegistrar, srv ViolationServiceServer) {
	s.RegisterService(&ViolationService_ServiceDesc, srv)
}

func _ViolationService_GetViolations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViolationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViolationServiceServer).GetViolations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViolationService_GetViolations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViolationServiceServer).GetViolations(ctx, req.(*GetViolationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ViolationService_ServiceDesc is the grpc.ServiceDesc for ViolationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViolationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blogging.ViolationService",
	HandlerType: (*ViolationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetViolations",
			Handler:    _ViolationService_GetViolations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogging/violation.proto",
}