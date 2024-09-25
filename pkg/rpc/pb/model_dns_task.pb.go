// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_dns_task.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DNS相关同步任务
type DNSTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string       `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	IsDone      bool         `protobuf:"varint,3,opt,name=isDone,proto3" json:"isDone,omitempty"`
	IsOk        bool         `protobuf:"varint,4,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Error       string       `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	UpdatedAt   int64        `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Node        *Node        `protobuf:"bytes,30,opt,name=node,proto3" json:"node,omitempty"`
	NodeCluster *NodeCluster `protobuf:"bytes,31,opt,name=nodeCluster,proto3" json:"nodeCluster,omitempty"`
	Server      *Server      `protobuf:"bytes,32,opt,name=server,proto3" json:"server,omitempty"`
	DnsDomain   *DNSDomain   `protobuf:"bytes,33,opt,name=dnsDomain,proto3" json:"dnsDomain,omitempty"`
}

func (x *DNSTask) Reset() {
	*x = DNSTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_dns_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSTask) ProtoMessage() {}

func (x *DNSTask) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_dns_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSTask.ProtoReflect.Descriptor instead.
func (*DNSTask) Descriptor() ([]byte, []int) {
	return file_models_model_dns_task_proto_rawDescGZIP(), []int{0}
}

func (x *DNSTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DNSTask) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DNSTask) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *DNSTask) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *DNSTask) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DNSTask) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DNSTask) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DNSTask) GetNodeCluster() *NodeCluster {
	if x != nil {
		return x.NodeCluster
	}
	return nil
}

func (x *DNSTask) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *DNSTask) GetDnsDomain() *DNSDomain {
	if x != nil {
		return x.DnsDomain
	}
	return nil
}

var File_models_model_dns_task_proto protoreflect.FileDescriptor

var file_models_model_dns_task_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64,
	0x6e, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x07, 0x44, 0x4e, 0x53, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x64, 0x6e,
	0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x09, 0x64, 0x6e,
	0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_dns_task_proto_rawDescOnce sync.Once
	file_models_model_dns_task_proto_rawDescData = file_models_model_dns_task_proto_rawDesc
)

func file_models_model_dns_task_proto_rawDescGZIP() []byte {
	file_models_model_dns_task_proto_rawDescOnce.Do(func() {
		file_models_model_dns_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_dns_task_proto_rawDescData)
	})
	return file_models_model_dns_task_proto_rawDescData
}

var file_models_model_dns_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_dns_task_proto_goTypes = []interface{}{
	(*DNSTask)(nil),     // 0: pb.DNSTask
	(*Node)(nil),        // 1: pb.Node
	(*NodeCluster)(nil), // 2: pb.NodeCluster
	(*Server)(nil),      // 3: pb.Server
	(*DNSDomain)(nil),   // 4: pb.DNSDomain
}
var file_models_model_dns_task_proto_depIdxs = []int32{
	1, // 0: pb.DNSTask.node:type_name -> pb.Node
	2, // 1: pb.DNSTask.nodeCluster:type_name -> pb.NodeCluster
	3, // 2: pb.DNSTask.server:type_name -> pb.Server
	4, // 3: pb.DNSTask.dnsDomain:type_name -> pb.DNSDomain
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_models_model_dns_task_proto_init() }
func file_models_model_dns_task_proto_init() {
	if File_models_model_dns_task_proto != nil {
		return
	}
	file_models_model_node_proto_init()
	file_models_model_node_cluster_proto_init()
	file_models_model_server_proto_init()
	file_models_model_dns_domain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_dns_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_model_dns_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_dns_task_proto_goTypes,
		DependencyIndexes: file_models_model_dns_task_proto_depIdxs,
		MessageInfos:      file_models_model_dns_task_proto_msgTypes,
	}.Build()
	File_models_model_dns_task_proto = out.File
	file_models_model_dns_task_proto_rawDesc = nil
	file_models_model_dns_task_proto_goTypes = nil
	file_models_model_dns_task_proto_depIdxs = nil
}