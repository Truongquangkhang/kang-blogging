// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: blogging/iam.proto

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
	IAMService_Login_FullMethodName              = "/blogging.IAMService/Login"
	IAMService_Register_FullMethodName           = "/blogging.IAMService/Register"
	IAMService_CheckExistUsername_FullMethodName = "/blogging.IAMService/CheckExistUsername"
	IAMService_RefreshAccessToken_FullMethodName = "/blogging.IAMService/RefreshAccessToken"
)

// IAMServiceClient is the client API for IAMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IAMServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CheckExistUsername(ctx context.Context, in *CheckExistUsernameRequest, opts ...grpc.CallOption) (*CheckExistUsernameResponse, error)
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
}

type iAMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMServiceClient(cc grpc.ClientConnInterface) IAMServiceClient {
	return &iAMServiceClient{cc}
}

func (c *iAMServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, IAMService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, IAMService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CheckExistUsername(ctx context.Context, in *CheckExistUsernameRequest, opts ...grpc.CallOption) (*CheckExistUsernameResponse, error) {
	out := new(CheckExistUsernameResponse)
	err := c.cc.Invoke(ctx, IAMService_CheckExistUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, IAMService_RefreshAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMServiceServer is the server API for IAMService service.
// All implementations should embed UnimplementedIAMServiceServer
// for forward compatibility
type IAMServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CheckExistUsername(context.Context, *CheckExistUsernameRequest) (*CheckExistUsernameResponse, error)
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

// UnimplementedIAMServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIAMServiceServer struct {
}

func (UnimplementedIAMServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIAMServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIAMServiceServer) CheckExistUsername(context.Context, *CheckExistUsernameRequest) (*CheckExistUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistUsername not implemented")
}
func (UnimplementedIAMServiceServer) RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}

// UnsafeIAMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMServiceServer will
// result in compilation errors.
type UnsafeIAMServiceServer interface {
	mustEmbedUnimplementedIAMServiceServer()
}

func RegisterIAMServiceServer(s grpc.ServiceRegistrar, srv IAMServiceServer) {
	s.RegisterService(&IAMService_ServiceDesc, srv)
}

func _IAMService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CheckExistUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExistUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CheckExistUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CheckExistUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CheckExistUsername(ctx, req.(*CheckExistUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_RefreshAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IAMService_ServiceDesc is the grpc.ServiceDesc for IAMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blogging.IAMService",
	HandlerType: (*IAMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IAMService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _IAMService_Register_Handler,
		},
		{
			MethodName: "CheckExistUsername",
			Handler:    _IAMService_CheckExistUsername_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _IAMService_RefreshAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogging/iam.proto",
}
