// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_node_login.proto

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
	NodeLoginService_FindNodeLoginSuggestPorts_FullMethodName = "/pb.NodeLoginService/findNodeLoginSuggestPorts"
)

// NodeLoginServiceClient is the client API for NodeLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeLoginServiceClient interface {
	// 读取建议的端口
	FindNodeLoginSuggestPorts(ctx context.Context, in *FindNodeLoginSuggestPortsRequest, opts ...grpc.CallOption) (*FindNodeLoginSuggestPortsResponse, error)
}

type nodeLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeLoginServiceClient(cc grpc.ClientConnInterface) NodeLoginServiceClient {
	return &nodeLoginServiceClient{cc}
}

func (c *nodeLoginServiceClient) FindNodeLoginSuggestPorts(ctx context.Context, in *FindNodeLoginSuggestPortsRequest, opts ...grpc.CallOption) (*FindNodeLoginSuggestPortsResponse, error) {
	out := new(FindNodeLoginSuggestPortsResponse)
	err := c.cc.Invoke(ctx, NodeLoginService_FindNodeLoginSuggestPorts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeLoginServiceServer is the server API for NodeLoginService service.
// All implementations should embed UnimplementedNodeLoginServiceServer
// for forward compatibility
type NodeLoginServiceServer interface {
	// 读取建议的端口
	FindNodeLoginSuggestPorts(context.Context, *FindNodeLoginSuggestPortsRequest) (*FindNodeLoginSuggestPortsResponse, error)
}

// UnimplementedNodeLoginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNodeLoginServiceServer struct {
}

func (UnimplementedNodeLoginServiceServer) FindNodeLoginSuggestPorts(context.Context, *FindNodeLoginSuggestPortsRequest) (*FindNodeLoginSuggestPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNodeLoginSuggestPorts not implemented")
}

// UnsafeNodeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeLoginServiceServer will
// result in compilation errors.
type UnsafeNodeLoginServiceServer interface {
	mustEmbedUnimplementedNodeLoginServiceServer()
}

func RegisterNodeLoginServiceServer(s grpc.ServiceRegistrar, srv NodeLoginServiceServer) {
	s.RegisterService(&NodeLoginService_ServiceDesc, srv)
}

func _NodeLoginService_FindNodeLoginSuggestPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodeLoginSuggestPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeLoginServiceServer).FindNodeLoginSuggestPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeLoginService_FindNodeLoginSuggestPorts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeLoginServiceServer).FindNodeLoginSuggestPorts(ctx, req.(*FindNodeLoginSuggestPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeLoginService_ServiceDesc is the grpc.ServiceDesc for NodeLoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeLoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeLoginService",
	HandlerType: (*NodeLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findNodeLoginSuggestPorts",
			Handler:    _NodeLoginService_FindNodeLoginSuggestPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_login.proto",
}
