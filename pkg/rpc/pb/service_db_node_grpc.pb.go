// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_db_node.proto

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
	DBNodeService_CreateDBNode_FullMethodName           = "/pb.DBNodeService/createDBNode"
	DBNodeService_UpdateDBNode_FullMethodName           = "/pb.DBNodeService/updateDBNode"
	DBNodeService_DeleteDBNode_FullMethodName           = "/pb.DBNodeService/deleteDBNode"
	DBNodeService_CountAllEnabledDBNodes_FullMethodName = "/pb.DBNodeService/countAllEnabledDBNodes"
	DBNodeService_ListEnabledDBNodes_FullMethodName     = "/pb.DBNodeService/listEnabledDBNodes"
	DBNodeService_FindEnabledDBNode_FullMethodName      = "/pb.DBNodeService/findEnabledDBNode"
	DBNodeService_FindAllDBNodeTables_FullMethodName    = "/pb.DBNodeService/findAllDBNodeTables"
	DBNodeService_DeleteDBNodeTable_FullMethodName      = "/pb.DBNodeService/deleteDBNodeTable"
	DBNodeService_TruncateDBNodeTable_FullMethodName    = "/pb.DBNodeService/truncateDBNodeTable"
	DBNodeService_CheckDBNodeStatus_FullMethodName      = "/pb.DBNodeService/checkDBNodeStatus"
)

// DBNodeServiceClient is the client API for DBNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBNodeServiceClient interface {
	// 创建数据库节点
	CreateDBNode(ctx context.Context, in *CreateDBNodeRequest, opts ...grpc.CallOption) (*CreateDBNodeResponse, error)
	// 修改数据库节点
	UpdateDBNode(ctx context.Context, in *UpdateDBNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除节点
	DeleteDBNode(ctx context.Context, in *DeleteDBNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算可用的数据库节点数量
	CountAllEnabledDBNodes(ctx context.Context, in *CountAllEnabledDBNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的数据库节点
	ListEnabledDBNodes(ctx context.Context, in *ListEnabledDBNodesRequest, opts ...grpc.CallOption) (*ListEnabledDBNodesResponse, error)
	// 根据ID查找可用的数据库节点
	FindEnabledDBNode(ctx context.Context, in *FindEnabledDBNodeRequest, opts ...grpc.CallOption) (*FindEnabledDBNodeResponse, error)
	// 获取所有表信息
	FindAllDBNodeTables(ctx context.Context, in *FindAllDBNodeTablesRequest, opts ...grpc.CallOption) (*FindAllDBNodeTablesResponse, error)
	// 删除表
	DeleteDBNodeTable(ctx context.Context, in *DeleteDBNodeTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 清空表
	TruncateDBNodeTable(ctx context.Context, in *TruncateDBNodeTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 检查数据库节点状态
	CheckDBNodeStatus(ctx context.Context, in *CheckDBNodeStatusRequest, opts ...grpc.CallOption) (*CheckDBNodeStatusResponse, error)
}

type dBNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBNodeServiceClient(cc grpc.ClientConnInterface) DBNodeServiceClient {
	return &dBNodeServiceClient{cc}
}

func (c *dBNodeServiceClient) CreateDBNode(ctx context.Context, in *CreateDBNodeRequest, opts ...grpc.CallOption) (*CreateDBNodeResponse, error) {
	out := new(CreateDBNodeResponse)
	err := c.cc.Invoke(ctx, DBNodeService_CreateDBNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) UpdateDBNode(ctx context.Context, in *UpdateDBNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DBNodeService_UpdateDBNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) DeleteDBNode(ctx context.Context, in *DeleteDBNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DBNodeService_DeleteDBNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) CountAllEnabledDBNodes(ctx context.Context, in *CountAllEnabledDBNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, DBNodeService_CountAllEnabledDBNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) ListEnabledDBNodes(ctx context.Context, in *ListEnabledDBNodesRequest, opts ...grpc.CallOption) (*ListEnabledDBNodesResponse, error) {
	out := new(ListEnabledDBNodesResponse)
	err := c.cc.Invoke(ctx, DBNodeService_ListEnabledDBNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) FindEnabledDBNode(ctx context.Context, in *FindEnabledDBNodeRequest, opts ...grpc.CallOption) (*FindEnabledDBNodeResponse, error) {
	out := new(FindEnabledDBNodeResponse)
	err := c.cc.Invoke(ctx, DBNodeService_FindEnabledDBNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) FindAllDBNodeTables(ctx context.Context, in *FindAllDBNodeTablesRequest, opts ...grpc.CallOption) (*FindAllDBNodeTablesResponse, error) {
	out := new(FindAllDBNodeTablesResponse)
	err := c.cc.Invoke(ctx, DBNodeService_FindAllDBNodeTables_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) DeleteDBNodeTable(ctx context.Context, in *DeleteDBNodeTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DBNodeService_DeleteDBNodeTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) TruncateDBNodeTable(ctx context.Context, in *TruncateDBNodeTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, DBNodeService_TruncateDBNodeTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBNodeServiceClient) CheckDBNodeStatus(ctx context.Context, in *CheckDBNodeStatusRequest, opts ...grpc.CallOption) (*CheckDBNodeStatusResponse, error) {
	out := new(CheckDBNodeStatusResponse)
	err := c.cc.Invoke(ctx, DBNodeService_CheckDBNodeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBNodeServiceServer is the server API for DBNodeService service.
// All implementations should embed UnimplementedDBNodeServiceServer
// for forward compatibility
type DBNodeServiceServer interface {
	// 创建数据库节点
	CreateDBNode(context.Context, *CreateDBNodeRequest) (*CreateDBNodeResponse, error)
	// 修改数据库节点
	UpdateDBNode(context.Context, *UpdateDBNodeRequest) (*RPCSuccess, error)
	// 删除节点
	DeleteDBNode(context.Context, *DeleteDBNodeRequest) (*RPCSuccess, error)
	// 计算可用的数据库节点数量
	CountAllEnabledDBNodes(context.Context, *CountAllEnabledDBNodesRequest) (*RPCCountResponse, error)
	// 列出单页的数据库节点
	ListEnabledDBNodes(context.Context, *ListEnabledDBNodesRequest) (*ListEnabledDBNodesResponse, error)
	// 根据ID查找可用的数据库节点
	FindEnabledDBNode(context.Context, *FindEnabledDBNodeRequest) (*FindEnabledDBNodeResponse, error)
	// 获取所有表信息
	FindAllDBNodeTables(context.Context, *FindAllDBNodeTablesRequest) (*FindAllDBNodeTablesResponse, error)
	// 删除表
	DeleteDBNodeTable(context.Context, *DeleteDBNodeTableRequest) (*RPCSuccess, error)
	// 清空表
	TruncateDBNodeTable(context.Context, *TruncateDBNodeTableRequest) (*RPCSuccess, error)
	// 检查数据库节点状态
	CheckDBNodeStatus(context.Context, *CheckDBNodeStatusRequest) (*CheckDBNodeStatusResponse, error)
}

// UnimplementedDBNodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDBNodeServiceServer struct {
}

func (UnimplementedDBNodeServiceServer) CreateDBNode(context.Context, *CreateDBNodeRequest) (*CreateDBNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDBNode not implemented")
}
func (UnimplementedDBNodeServiceServer) UpdateDBNode(context.Context, *UpdateDBNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDBNode not implemented")
}
func (UnimplementedDBNodeServiceServer) DeleteDBNode(context.Context, *DeleteDBNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDBNode not implemented")
}
func (UnimplementedDBNodeServiceServer) CountAllEnabledDBNodes(context.Context, *CountAllEnabledDBNodesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledDBNodes not implemented")
}
func (UnimplementedDBNodeServiceServer) ListEnabledDBNodes(context.Context, *ListEnabledDBNodesRequest) (*ListEnabledDBNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledDBNodes not implemented")
}
func (UnimplementedDBNodeServiceServer) FindEnabledDBNode(context.Context, *FindEnabledDBNodeRequest) (*FindEnabledDBNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledDBNode not implemented")
}
func (UnimplementedDBNodeServiceServer) FindAllDBNodeTables(context.Context, *FindAllDBNodeTablesRequest) (*FindAllDBNodeTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDBNodeTables not implemented")
}
func (UnimplementedDBNodeServiceServer) DeleteDBNodeTable(context.Context, *DeleteDBNodeTableRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDBNodeTable not implemented")
}
func (UnimplementedDBNodeServiceServer) TruncateDBNodeTable(context.Context, *TruncateDBNodeTableRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TruncateDBNodeTable not implemented")
}
func (UnimplementedDBNodeServiceServer) CheckDBNodeStatus(context.Context, *CheckDBNodeStatusRequest) (*CheckDBNodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDBNodeStatus not implemented")
}

// UnsafeDBNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBNodeServiceServer will
// result in compilation errors.
type UnsafeDBNodeServiceServer interface {
	mustEmbedUnimplementedDBNodeServiceServer()
}

func RegisterDBNodeServiceServer(s grpc.ServiceRegistrar, srv DBNodeServiceServer) {
	s.RegisterService(&DBNodeService_ServiceDesc, srv)
}

func _DBNodeService_CreateDBNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDBNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).CreateDBNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_CreateDBNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).CreateDBNode(ctx, req.(*CreateDBNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_UpdateDBNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDBNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).UpdateDBNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_UpdateDBNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).UpdateDBNode(ctx, req.(*UpdateDBNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_DeleteDBNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDBNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).DeleteDBNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_DeleteDBNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).DeleteDBNode(ctx, req.(*DeleteDBNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_CountAllEnabledDBNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledDBNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).CountAllEnabledDBNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_CountAllEnabledDBNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).CountAllEnabledDBNodes(ctx, req.(*CountAllEnabledDBNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_ListEnabledDBNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledDBNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).ListEnabledDBNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_ListEnabledDBNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).ListEnabledDBNodes(ctx, req.(*ListEnabledDBNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_FindEnabledDBNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledDBNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).FindEnabledDBNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_FindEnabledDBNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).FindEnabledDBNode(ctx, req.(*FindEnabledDBNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_FindAllDBNodeTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDBNodeTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).FindAllDBNodeTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_FindAllDBNodeTables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).FindAllDBNodeTables(ctx, req.(*FindAllDBNodeTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_DeleteDBNodeTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDBNodeTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).DeleteDBNodeTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_DeleteDBNodeTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).DeleteDBNodeTable(ctx, req.(*DeleteDBNodeTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_TruncateDBNodeTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateDBNodeTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).TruncateDBNodeTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_TruncateDBNodeTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).TruncateDBNodeTable(ctx, req.(*TruncateDBNodeTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBNodeService_CheckDBNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDBNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBNodeServiceServer).CheckDBNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBNodeService_CheckDBNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBNodeServiceServer).CheckDBNodeStatus(ctx, req.(*CheckDBNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBNodeService_ServiceDesc is the grpc.ServiceDesc for DBNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DBNodeService",
	HandlerType: (*DBNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDBNode",
			Handler:    _DBNodeService_CreateDBNode_Handler,
		},
		{
			MethodName: "updateDBNode",
			Handler:    _DBNodeService_UpdateDBNode_Handler,
		},
		{
			MethodName: "deleteDBNode",
			Handler:    _DBNodeService_DeleteDBNode_Handler,
		},
		{
			MethodName: "countAllEnabledDBNodes",
			Handler:    _DBNodeService_CountAllEnabledDBNodes_Handler,
		},
		{
			MethodName: "listEnabledDBNodes",
			Handler:    _DBNodeService_ListEnabledDBNodes_Handler,
		},
		{
			MethodName: "findEnabledDBNode",
			Handler:    _DBNodeService_FindEnabledDBNode_Handler,
		},
		{
			MethodName: "findAllDBNodeTables",
			Handler:    _DBNodeService_FindAllDBNodeTables_Handler,
		},
		{
			MethodName: "deleteDBNodeTable",
			Handler:    _DBNodeService_DeleteDBNodeTable_Handler,
		},
		{
			MethodName: "truncateDBNodeTable",
			Handler:    _DBNodeService_TruncateDBNodeTable_Handler,
		},
		{
			MethodName: "checkDBNodeStatus",
			Handler:    _DBNodeService_CheckDBNodeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_db_node.proto",
}