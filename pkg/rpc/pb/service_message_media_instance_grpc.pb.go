// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_message_media_instance.proto

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
	MessageMediaInstanceService_CreateMessageMediaInstance_FullMethodName           = "/pb.MessageMediaInstanceService/createMessageMediaInstance"
	MessageMediaInstanceService_UpdateMessageMediaInstance_FullMethodName           = "/pb.MessageMediaInstanceService/updateMessageMediaInstance"
	MessageMediaInstanceService_DeleteMessageMediaInstance_FullMethodName           = "/pb.MessageMediaInstanceService/deleteMessageMediaInstance"
	MessageMediaInstanceService_CountAllEnabledMessageMediaInstances_FullMethodName = "/pb.MessageMediaInstanceService/countAllEnabledMessageMediaInstances"
	MessageMediaInstanceService_ListEnabledMessageMediaInstances_FullMethodName     = "/pb.MessageMediaInstanceService/listEnabledMessageMediaInstances"
	MessageMediaInstanceService_FindEnabledMessageMediaInstance_FullMethodName      = "/pb.MessageMediaInstanceService/findEnabledMessageMediaInstance"
)

// MessageMediaInstanceServiceClient is the client API for MessageMediaInstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageMediaInstanceServiceClient interface {
	// 创建消息媒介实例
	CreateMessageMediaInstance(ctx context.Context, in *CreateMessageMediaInstanceRequest, opts ...grpc.CallOption) (*CreateMessageMediaInstanceResponse, error)
	// 修改消息媒介实例
	UpdateMessageMediaInstance(ctx context.Context, in *UpdateMessageMediaInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除消息媒介实例
	DeleteMessageMediaInstance(ctx context.Context, in *DeleteMessageMediaInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算消息媒介实例数量
	CountAllEnabledMessageMediaInstances(ctx context.Context, in *CountAllEnabledMessageMediaInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页消息媒介实例
	ListEnabledMessageMediaInstances(ctx context.Context, in *ListEnabledMessageMediaInstancesRequest, opts ...grpc.CallOption) (*ListEnabledMessageMediaInstancesResponse, error)
	// 查找单个消息媒介实例信息
	FindEnabledMessageMediaInstance(ctx context.Context, in *FindEnabledMessageMediaInstanceRequest, opts ...grpc.CallOption) (*FindEnabledMessageMediaInstanceResponse, error)
}

type messageMediaInstanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageMediaInstanceServiceClient(cc grpc.ClientConnInterface) MessageMediaInstanceServiceClient {
	return &messageMediaInstanceServiceClient{cc}
}

func (c *messageMediaInstanceServiceClient) CreateMessageMediaInstance(ctx context.Context, in *CreateMessageMediaInstanceRequest, opts ...grpc.CallOption) (*CreateMessageMediaInstanceResponse, error) {
	out := new(CreateMessageMediaInstanceResponse)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_CreateMessageMediaInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaInstanceServiceClient) UpdateMessageMediaInstance(ctx context.Context, in *UpdateMessageMediaInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_UpdateMessageMediaInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaInstanceServiceClient) DeleteMessageMediaInstance(ctx context.Context, in *DeleteMessageMediaInstanceRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_DeleteMessageMediaInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaInstanceServiceClient) CountAllEnabledMessageMediaInstances(ctx context.Context, in *CountAllEnabledMessageMediaInstancesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_CountAllEnabledMessageMediaInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaInstanceServiceClient) ListEnabledMessageMediaInstances(ctx context.Context, in *ListEnabledMessageMediaInstancesRequest, opts ...grpc.CallOption) (*ListEnabledMessageMediaInstancesResponse, error) {
	out := new(ListEnabledMessageMediaInstancesResponse)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_ListEnabledMessageMediaInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaInstanceServiceClient) FindEnabledMessageMediaInstance(ctx context.Context, in *FindEnabledMessageMediaInstanceRequest, opts ...grpc.CallOption) (*FindEnabledMessageMediaInstanceResponse, error) {
	out := new(FindEnabledMessageMediaInstanceResponse)
	err := c.cc.Invoke(ctx, MessageMediaInstanceService_FindEnabledMessageMediaInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageMediaInstanceServiceServer is the server API for MessageMediaInstanceService service.
// All implementations should embed UnimplementedMessageMediaInstanceServiceServer
// for forward compatibility
type MessageMediaInstanceServiceServer interface {
	// 创建消息媒介实例
	CreateMessageMediaInstance(context.Context, *CreateMessageMediaInstanceRequest) (*CreateMessageMediaInstanceResponse, error)
	// 修改消息媒介实例
	UpdateMessageMediaInstance(context.Context, *UpdateMessageMediaInstanceRequest) (*RPCSuccess, error)
	// 删除消息媒介实例
	DeleteMessageMediaInstance(context.Context, *DeleteMessageMediaInstanceRequest) (*RPCSuccess, error)
	// 计算消息媒介实例数量
	CountAllEnabledMessageMediaInstances(context.Context, *CountAllEnabledMessageMediaInstancesRequest) (*RPCCountResponse, error)
	// 列出单页消息媒介实例
	ListEnabledMessageMediaInstances(context.Context, *ListEnabledMessageMediaInstancesRequest) (*ListEnabledMessageMediaInstancesResponse, error)
	// 查找单个消息媒介实例信息
	FindEnabledMessageMediaInstance(context.Context, *FindEnabledMessageMediaInstanceRequest) (*FindEnabledMessageMediaInstanceResponse, error)
}

// UnimplementedMessageMediaInstanceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMessageMediaInstanceServiceServer struct {
}

func (UnimplementedMessageMediaInstanceServiceServer) CreateMessageMediaInstance(context.Context, *CreateMessageMediaInstanceRequest) (*CreateMessageMediaInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageMediaInstance not implemented")
}
func (UnimplementedMessageMediaInstanceServiceServer) UpdateMessageMediaInstance(context.Context, *UpdateMessageMediaInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageMediaInstance not implemented")
}
func (UnimplementedMessageMediaInstanceServiceServer) DeleteMessageMediaInstance(context.Context, *DeleteMessageMediaInstanceRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageMediaInstance not implemented")
}
func (UnimplementedMessageMediaInstanceServiceServer) CountAllEnabledMessageMediaInstances(context.Context, *CountAllEnabledMessageMediaInstancesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledMessageMediaInstances not implemented")
}
func (UnimplementedMessageMediaInstanceServiceServer) ListEnabledMessageMediaInstances(context.Context, *ListEnabledMessageMediaInstancesRequest) (*ListEnabledMessageMediaInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledMessageMediaInstances not implemented")
}
func (UnimplementedMessageMediaInstanceServiceServer) FindEnabledMessageMediaInstance(context.Context, *FindEnabledMessageMediaInstanceRequest) (*FindEnabledMessageMediaInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledMessageMediaInstance not implemented")
}

// UnsafeMessageMediaInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageMediaInstanceServiceServer will
// result in compilation errors.
type UnsafeMessageMediaInstanceServiceServer interface {
	mustEmbedUnimplementedMessageMediaInstanceServiceServer()
}

func RegisterMessageMediaInstanceServiceServer(s grpc.ServiceRegistrar, srv MessageMediaInstanceServiceServer) {
	s.RegisterService(&MessageMediaInstanceService_ServiceDesc, srv)
}

func _MessageMediaInstanceService_CreateMessageMediaInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageMediaInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).CreateMessageMediaInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_CreateMessageMediaInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).CreateMessageMediaInstance(ctx, req.(*CreateMessageMediaInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaInstanceService_UpdateMessageMediaInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageMediaInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).UpdateMessageMediaInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_UpdateMessageMediaInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).UpdateMessageMediaInstance(ctx, req.(*UpdateMessageMediaInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaInstanceService_DeleteMessageMediaInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageMediaInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).DeleteMessageMediaInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_DeleteMessageMediaInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).DeleteMessageMediaInstance(ctx, req.(*DeleteMessageMediaInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaInstanceService_CountAllEnabledMessageMediaInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledMessageMediaInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).CountAllEnabledMessageMediaInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_CountAllEnabledMessageMediaInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).CountAllEnabledMessageMediaInstances(ctx, req.(*CountAllEnabledMessageMediaInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaInstanceService_ListEnabledMessageMediaInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledMessageMediaInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).ListEnabledMessageMediaInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_ListEnabledMessageMediaInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).ListEnabledMessageMediaInstances(ctx, req.(*ListEnabledMessageMediaInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaInstanceService_FindEnabledMessageMediaInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledMessageMediaInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaInstanceServiceServer).FindEnabledMessageMediaInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageMediaInstanceService_FindEnabledMessageMediaInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaInstanceServiceServer).FindEnabledMessageMediaInstance(ctx, req.(*FindEnabledMessageMediaInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageMediaInstanceService_ServiceDesc is the grpc.ServiceDesc for MessageMediaInstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageMediaInstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageMediaInstanceService",
	HandlerType: (*MessageMediaInstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMessageMediaInstance",
			Handler:    _MessageMediaInstanceService_CreateMessageMediaInstance_Handler,
		},
		{
			MethodName: "updateMessageMediaInstance",
			Handler:    _MessageMediaInstanceService_UpdateMessageMediaInstance_Handler,
		},
		{
			MethodName: "deleteMessageMediaInstance",
			Handler:    _MessageMediaInstanceService_DeleteMessageMediaInstance_Handler,
		},
		{
			MethodName: "countAllEnabledMessageMediaInstances",
			Handler:    _MessageMediaInstanceService_CountAllEnabledMessageMediaInstances_Handler,
		},
		{
			MethodName: "listEnabledMessageMediaInstances",
			Handler:    _MessageMediaInstanceService_ListEnabledMessageMediaInstances_Handler,
		},
		{
			MethodName: "findEnabledMessageMediaInstance",
			Handler:    _MessageMediaInstanceService_FindEnabledMessageMediaInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_message_media_instance.proto",
}
