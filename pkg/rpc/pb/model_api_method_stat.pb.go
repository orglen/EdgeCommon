// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_api_method_stat.proto

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

type APIMethodStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ApiNodeId  int64    `protobuf:"varint,2,opt,name=apiNodeId,proto3" json:"apiNodeId,omitempty"`
	Method     string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Tag        string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	CostMs     float32  `protobuf:"fixed32,5,opt,name=costMs,proto3" json:"costMs,omitempty"`
	PeekMs     float32  `protobuf:"fixed32,6,opt,name=peekMs,proto3" json:"peekMs,omitempty"`
	CountCalls int64    `protobuf:"varint,7,opt,name=countCalls,proto3" json:"countCalls,omitempty"`
	ApiNode    *APINode `protobuf:"bytes,30,opt,name=apiNode,proto3" json:"apiNode,omitempty"`
}

func (x *APIMethodStat) Reset() {
	*x = APIMethodStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_api_method_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIMethodStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIMethodStat) ProtoMessage() {}

func (x *APIMethodStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_api_method_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIMethodStat.ProtoReflect.Descriptor instead.
func (*APIMethodStat) Descriptor() ([]byte, []int) {
	return file_models_model_api_method_stat_proto_rawDescGZIP(), []int{0}
}

func (x *APIMethodStat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *APIMethodStat) GetApiNodeId() int64 {
	if x != nil {
		return x.ApiNodeId
	}
	return 0
}

func (x *APIMethodStat) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *APIMethodStat) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *APIMethodStat) GetCostMs() float32 {
	if x != nil {
		return x.CostMs
	}
	return 0
}

func (x *APIMethodStat) GetPeekMs() float32 {
	if x != nil {
		return x.PeekMs
	}
	return 0
}

func (x *APIMethodStat) GetCountCalls() int64 {
	if x != nil {
		return x.CountCalls
	}
	return 0
}

func (x *APIMethodStat) GetApiNode() *APINode {
	if x != nil {
		return x.ApiNode
	}
	return nil
}

var File_models_model_api_method_stat_proto protoreflect.FileDescriptor

var file_models_model_api_method_stat_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x70, 0x69, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x73, 0x74, 0x4d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x63, 0x6f, 0x73, 0x74, 0x4d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x6b, 0x4d,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x65, 0x65, 0x6b, 0x4d, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12,
	0x25, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x50, 0x49, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x61,
	0x70, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_api_method_stat_proto_rawDescOnce sync.Once
	file_models_model_api_method_stat_proto_rawDescData = file_models_model_api_method_stat_proto_rawDesc
)

func file_models_model_api_method_stat_proto_rawDescGZIP() []byte {
	file_models_model_api_method_stat_proto_rawDescOnce.Do(func() {
		file_models_model_api_method_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_api_method_stat_proto_rawDescData)
	})
	return file_models_model_api_method_stat_proto_rawDescData
}

var file_models_model_api_method_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_api_method_stat_proto_goTypes = []interface{}{
	(*APIMethodStat)(nil), // 0: pb.APIMethodStat
	(*APINode)(nil),       // 1: pb.APINode
}
var file_models_model_api_method_stat_proto_depIdxs = []int32{
	1, // 0: pb.APIMethodStat.apiNode:type_name -> pb.APINode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_api_method_stat_proto_init() }
func file_models_model_api_method_stat_proto_init() {
	if File_models_model_api_method_stat_proto != nil {
		return
	}
	file_models_model_api_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_api_method_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIMethodStat); i {
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
			RawDescriptor: file_models_model_api_method_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_api_method_stat_proto_goTypes,
		DependencyIndexes: file_models_model_api_method_stat_proto_depIdxs,
		MessageInfos:      file_models_model_api_method_stat_proto_msgTypes,
	}.Build()
	File_models_model_api_method_stat_proto = out.File
	file_models_model_api_method_stat_proto_rawDesc = nil
	file_models_model_api_method_stat_proto_goTypes = nil
	file_models_model_api_method_stat_proto_depIdxs = nil
}