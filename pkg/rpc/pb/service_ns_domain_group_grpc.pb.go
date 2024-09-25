// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ns_domain_group.proto

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
	NSDomainGroupService_CreateNSDomainGroup_FullMethodName             = "/pb.NSDomainGroupService/createNSDomainGroup"
	NSDomainGroupService_UpdateNSDomainGroup_FullMethodName             = "/pb.NSDomainGroupService/updateNSDomainGroup"
	NSDomainGroupService_DeleteNSDomainGroup_FullMethodName             = "/pb.NSDomainGroupService/deleteNSDomainGroup"
	NSDomainGroupService_FindAllNSDomainGroups_FullMethodName           = "/pb.NSDomainGroupService/findAllNSDomainGroups"
	NSDomainGroupService_CountAllAvailableNSDomainGroups_FullMethodName = "/pb.NSDomainGroupService/countAllAvailableNSDomainGroups"
	NSDomainGroupService_FindAllAvailableNSDomainGroups_FullMethodName  = "/pb.NSDomainGroupService/findAllAvailableNSDomainGroups"
	NSDomainGroupService_FindNSDomainGroup_FullMethodName               = "/pb.NSDomainGroupService/findNSDomainGroup"
)

// NSDomainGroupServiceClient is the client API for NSDomainGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSDomainGroupServiceClient interface {
	// 创建分组
	CreateNSDomainGroup(ctx context.Context, in *CreateNSDomainGroupRequest, opts ...grpc.CallOption) (*CreateNSDomainGroupResponse, error)
	// 修改分组
	UpdateNSDomainGroup(ctx context.Context, in *UpdateNSDomainGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除分组
	DeleteNSDomainGroup(ctx context.Context, in *DeleteNSDomainGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询所有分组
	FindAllNSDomainGroups(ctx context.Context, in *FindAllNSDomainGroupsRequest, opts ...grpc.CallOption) (*FindAllNSDomainGroupsResponse, error)
	// 查询可用分组数量
	CountAllAvailableNSDomainGroups(ctx context.Context, in *CountAllAvailableNSDomainGroupsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 查询所有启用分组
	FindAllAvailableNSDomainGroups(ctx context.Context, in *FindAllAvailableNSDomainGroupsRequest, opts ...grpc.CallOption) (*FindAllAvailableNSDomainGroupsResponse, error)
	// 查找单个分组
	FindNSDomainGroup(ctx context.Context, in *FindNSDomainGroupRequest, opts ...grpc.CallOption) (*FindNSDomainGroupResponse, error)
}

type nSDomainGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSDomainGroupServiceClient(cc grpc.ClientConnInterface) NSDomainGroupServiceClient {
	return &nSDomainGroupServiceClient{cc}
}

func (c *nSDomainGroupServiceClient) CreateNSDomainGroup(ctx context.Context, in *CreateNSDomainGroupRequest, opts ...grpc.CallOption) (*CreateNSDomainGroupResponse, error) {
	out := new(CreateNSDomainGroupResponse)
	err := c.cc.Invoke(ctx, NSDomainGroupService_CreateNSDomainGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) UpdateNSDomainGroup(ctx context.Context, in *UpdateNSDomainGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainGroupService_UpdateNSDomainGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) DeleteNSDomainGroup(ctx context.Context, in *DeleteNSDomainGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NSDomainGroupService_DeleteNSDomainGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) FindAllNSDomainGroups(ctx context.Context, in *FindAllNSDomainGroupsRequest, opts ...grpc.CallOption) (*FindAllNSDomainGroupsResponse, error) {
	out := new(FindAllNSDomainGroupsResponse)
	err := c.cc.Invoke(ctx, NSDomainGroupService_FindAllNSDomainGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) CountAllAvailableNSDomainGroups(ctx context.Context, in *CountAllAvailableNSDomainGroupsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, NSDomainGroupService_CountAllAvailableNSDomainGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) FindAllAvailableNSDomainGroups(ctx context.Context, in *FindAllAvailableNSDomainGroupsRequest, opts ...grpc.CallOption) (*FindAllAvailableNSDomainGroupsResponse, error) {
	out := new(FindAllAvailableNSDomainGroupsResponse)
	err := c.cc.Invoke(ctx, NSDomainGroupService_FindAllAvailableNSDomainGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSDomainGroupServiceClient) FindNSDomainGroup(ctx context.Context, in *FindNSDomainGroupRequest, opts ...grpc.CallOption) (*FindNSDomainGroupResponse, error) {
	out := new(FindNSDomainGroupResponse)
	err := c.cc.Invoke(ctx, NSDomainGroupService_FindNSDomainGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSDomainGroupServiceServer is the server API for NSDomainGroupService service.
// All implementations should embed UnimplementedNSDomainGroupServiceServer
// for forward compatibility
type NSDomainGroupServiceServer interface {
	// 创建分组
	CreateNSDomainGroup(context.Context, *CreateNSDomainGroupRequest) (*CreateNSDomainGroupResponse, error)
	// 修改分组
	UpdateNSDomainGroup(context.Context, *UpdateNSDomainGroupRequest) (*RPCSuccess, error)
	// 删除分组
	DeleteNSDomainGroup(context.Context, *DeleteNSDomainGroupRequest) (*RPCSuccess, error)
	// 查询所有分组
	FindAllNSDomainGroups(context.Context, *FindAllNSDomainGroupsRequest) (*FindAllNSDomainGroupsResponse, error)
	// 查询可用分组数量
	CountAllAvailableNSDomainGroups(context.Context, *CountAllAvailableNSDomainGroupsRequest) (*RPCCountResponse, error)
	// 查询所有启用分组
	FindAllAvailableNSDomainGroups(context.Context, *FindAllAvailableNSDomainGroupsRequest) (*FindAllAvailableNSDomainGroupsResponse, error)
	// 查找单个分组
	FindNSDomainGroup(context.Context, *FindNSDomainGroupRequest) (*FindNSDomainGroupResponse, error)
}

// UnimplementedNSDomainGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNSDomainGroupServiceServer struct {
}

func (UnimplementedNSDomainGroupServiceServer) CreateNSDomainGroup(context.Context, *CreateNSDomainGroupRequest) (*CreateNSDomainGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNSDomainGroup not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) UpdateNSDomainGroup(context.Context, *UpdateNSDomainGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNSDomainGroup not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) DeleteNSDomainGroup(context.Context, *DeleteNSDomainGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNSDomainGroup not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) FindAllNSDomainGroups(context.Context, *FindAllNSDomainGroupsRequest) (*FindAllNSDomainGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllNSDomainGroups not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) CountAllAvailableNSDomainGroups(context.Context, *CountAllAvailableNSDomainGroupsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllAvailableNSDomainGroups not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) FindAllAvailableNSDomainGroups(context.Context, *FindAllAvailableNSDomainGroupsRequest) (*FindAllAvailableNSDomainGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllAvailableNSDomainGroups not implemented")
}
func (UnimplementedNSDomainGroupServiceServer) FindNSDomainGroup(context.Context, *FindNSDomainGroupRequest) (*FindNSDomainGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNSDomainGroup not implemented")
}

// UnsafeNSDomainGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSDomainGroupServiceServer will
// result in compilation errors.
type UnsafeNSDomainGroupServiceServer interface {
	mustEmbedUnimplementedNSDomainGroupServiceServer()
}

func RegisterNSDomainGroupServiceServer(s grpc.ServiceRegistrar, srv NSDomainGroupServiceServer) {
	s.RegisterService(&NSDomainGroupService_ServiceDesc, srv)
}

func _NSDomainGroupService_CreateNSDomainGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNSDomainGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).CreateNSDomainGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_CreateNSDomainGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).CreateNSDomainGroup(ctx, req.(*CreateNSDomainGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_UpdateNSDomainGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNSDomainGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).UpdateNSDomainGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_UpdateNSDomainGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).UpdateNSDomainGroup(ctx, req.(*UpdateNSDomainGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_DeleteNSDomainGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNSDomainGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).DeleteNSDomainGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_DeleteNSDomainGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).DeleteNSDomainGroup(ctx, req.(*DeleteNSDomainGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_FindAllNSDomainGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllNSDomainGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).FindAllNSDomainGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_FindAllNSDomainGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).FindAllNSDomainGroups(ctx, req.(*FindAllNSDomainGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_CountAllAvailableNSDomainGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllAvailableNSDomainGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).CountAllAvailableNSDomainGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_CountAllAvailableNSDomainGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).CountAllAvailableNSDomainGroups(ctx, req.(*CountAllAvailableNSDomainGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_FindAllAvailableNSDomainGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllAvailableNSDomainGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).FindAllAvailableNSDomainGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_FindAllAvailableNSDomainGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).FindAllAvailableNSDomainGroups(ctx, req.(*FindAllAvailableNSDomainGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSDomainGroupService_FindNSDomainGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNSDomainGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSDomainGroupServiceServer).FindNSDomainGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NSDomainGroupService_FindNSDomainGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSDomainGroupServiceServer).FindNSDomainGroup(ctx, req.(*FindNSDomainGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSDomainGroupService_ServiceDesc is the grpc.ServiceDesc for NSDomainGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSDomainGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSDomainGroupService",
	HandlerType: (*NSDomainGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNSDomainGroup",
			Handler:    _NSDomainGroupService_CreateNSDomainGroup_Handler,
		},
		{
			MethodName: "updateNSDomainGroup",
			Handler:    _NSDomainGroupService_UpdateNSDomainGroup_Handler,
		},
		{
			MethodName: "deleteNSDomainGroup",
			Handler:    _NSDomainGroupService_DeleteNSDomainGroup_Handler,
		},
		{
			MethodName: "findAllNSDomainGroups",
			Handler:    _NSDomainGroupService_FindAllNSDomainGroups_Handler,
		},
		{
			MethodName: "countAllAvailableNSDomainGroups",
			Handler:    _NSDomainGroupService_CountAllAvailableNSDomainGroups_Handler,
		},
		{
			MethodName: "findAllAvailableNSDomainGroups",
			Handler:    _NSDomainGroupService_FindAllAvailableNSDomainGroups_Handler,
		},
		{
			MethodName: "findNSDomainGroup",
			Handler:    _NSDomainGroupService_FindNSDomainGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns_domain_group.proto",
}