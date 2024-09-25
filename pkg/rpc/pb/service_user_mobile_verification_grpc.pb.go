// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_mobile_verification.proto

package pb

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
	UserMobileVerificationService_VerifyUserMobile_FullMethodName                 = "/pb.UserMobileVerificationService/verifyUserMobile"
	UserMobileVerificationService_SendUserMobileVerification_FullMethodName       = "/pb.UserMobileVerificationService/sendUserMobileVerification"
	UserMobileVerificationService_FindLatestUserMobileVerification_FullMethodName = "/pb.UserMobileVerificationService/findLatestUserMobileVerification"
)

// UserMobileVerificationServiceClient is the client API for UserMobileVerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMobileVerificationServiceClient interface {
	// 认证手机号码
	VerifyUserMobile(ctx context.Context, in *VerifyUserMobileRequest, opts ...grpc.CallOption) (*VerifyUserMobileResponse, error)
	// 发送手机号码认证
	SendUserMobileVerification(ctx context.Context, in *SendUserMobileVerificationRequest, opts ...grpc.CallOption) (*SendUserMobileVerificationResponse, error)
	// 查找用户正在等待激活的认证
	FindLatestUserMobileVerification(ctx context.Context, in *FindLatestUserMobileVerificationRequest, opts ...grpc.CallOption) (*FindLatestUserMobileVerificationResponse, error)
}

type userMobileVerificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMobileVerificationServiceClient(cc grpc.ClientConnInterface) UserMobileVerificationServiceClient {
	return &userMobileVerificationServiceClient{cc}
}

func (c *userMobileVerificationServiceClient) VerifyUserMobile(ctx context.Context, in *VerifyUserMobileRequest, opts ...grpc.CallOption) (*VerifyUserMobileResponse, error) {
	out := new(VerifyUserMobileResponse)
	err := c.cc.Invoke(ctx, UserMobileVerificationService_VerifyUserMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMobileVerificationServiceClient) SendUserMobileVerification(ctx context.Context, in *SendUserMobileVerificationRequest, opts ...grpc.CallOption) (*SendUserMobileVerificationResponse, error) {
	out := new(SendUserMobileVerificationResponse)
	err := c.cc.Invoke(ctx, UserMobileVerificationService_SendUserMobileVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMobileVerificationServiceClient) FindLatestUserMobileVerification(ctx context.Context, in *FindLatestUserMobileVerificationRequest, opts ...grpc.CallOption) (*FindLatestUserMobileVerificationResponse, error) {
	out := new(FindLatestUserMobileVerificationResponse)
	err := c.cc.Invoke(ctx, UserMobileVerificationService_FindLatestUserMobileVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMobileVerificationServiceServer is the server API for UserMobileVerificationService service.
// All implementations should embed UnimplementedUserMobileVerificationServiceServer
// for forward compatibility
type UserMobileVerificationServiceServer interface {
	// 认证手机号码
	VerifyUserMobile(context.Context, *VerifyUserMobileRequest) (*VerifyUserMobileResponse, error)
	// 发送手机号码认证
	SendUserMobileVerification(context.Context, *SendUserMobileVerificationRequest) (*SendUserMobileVerificationResponse, error)
	// 查找用户正在等待激活的认证
	FindLatestUserMobileVerification(context.Context, *FindLatestUserMobileVerificationRequest) (*FindLatestUserMobileVerificationResponse, error)
}

// UnimplementedUserMobileVerificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserMobileVerificationServiceServer struct {
}

func (UnimplementedUserMobileVerificationServiceServer) VerifyUserMobile(context.Context, *VerifyUserMobileRequest) (*VerifyUserMobileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserMobile not implemented")
}
func (UnimplementedUserMobileVerificationServiceServer) SendUserMobileVerification(context.Context, *SendUserMobileVerificationRequest) (*SendUserMobileVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserMobileVerification not implemented")
}
func (UnimplementedUserMobileVerificationServiceServer) FindLatestUserMobileVerification(context.Context, *FindLatestUserMobileVerificationRequest) (*FindLatestUserMobileVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestUserMobileVerification not implemented")
}

// UnsafeUserMobileVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMobileVerificationServiceServer will
// result in compilation errors.
type UnsafeUserMobileVerificationServiceServer interface {
	mustEmbedUnimplementedUserMobileVerificationServiceServer()
}

func RegisterUserMobileVerificationServiceServer(s grpc.ServiceRegistrar, srv UserMobileVerificationServiceServer) {
	s.RegisterService(&UserMobileVerificationService_ServiceDesc, srv)
}

func _UserMobileVerificationService_VerifyUserMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMobileVerificationServiceServer).VerifyUserMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMobileVerificationService_VerifyUserMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMobileVerificationServiceServer).VerifyUserMobile(ctx, req.(*VerifyUserMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMobileVerificationService_SendUserMobileVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserMobileVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMobileVerificationServiceServer).SendUserMobileVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMobileVerificationService_SendUserMobileVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMobileVerificationServiceServer).SendUserMobileVerification(ctx, req.(*SendUserMobileVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMobileVerificationService_FindLatestUserMobileVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestUserMobileVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMobileVerificationServiceServer).FindLatestUserMobileVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserMobileVerificationService_FindLatestUserMobileVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMobileVerificationServiceServer).FindLatestUserMobileVerification(ctx, req.(*FindLatestUserMobileVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMobileVerificationService_ServiceDesc is the grpc.ServiceDesc for UserMobileVerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMobileVerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserMobileVerificationService",
	HandlerType: (*UserMobileVerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "verifyUserMobile",
			Handler:    _UserMobileVerificationService_VerifyUserMobile_Handler,
		},
		{
			MethodName: "sendUserMobileVerification",
			Handler:    _UserMobileVerificationService_SendUserMobileVerification_Handler,
		},
		{
			MethodName: "findLatestUserMobileVerification",
			Handler:    _UserMobileVerificationService_FindLatestUserMobileVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_mobile_verification.proto",
}