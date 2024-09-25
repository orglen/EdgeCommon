// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_http_fastcgi.proto

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

// HTTP Fastcgi定义
type HTTPFastcgi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn            bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Address         string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ParamsJSON      []byte `protobuf:"bytes,4,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	ReadTimeoutJSON []byte `protobuf:"bytes,5,opt,name=readTimeoutJSON,proto3" json:"readTimeoutJSON,omitempty"`
	ConnTimeoutJSON []byte `protobuf:"bytes,6,opt,name=connTimeoutJSON,proto3" json:"connTimeoutJSON,omitempty"`
	PoolSize        int32  `protobuf:"varint,7,opt,name=poolSize,proto3" json:"poolSize,omitempty"`
	PathInfoPattern string `protobuf:"bytes,8,opt,name=pathInfoPattern,proto3" json:"pathInfoPattern,omitempty"`
}

func (x *HTTPFastcgi) Reset() {
	*x = HTTPFastcgi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_fastcgi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFastcgi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFastcgi) ProtoMessage() {}

func (x *HTTPFastcgi) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_fastcgi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFastcgi.ProtoReflect.Descriptor instead.
func (*HTTPFastcgi) Descriptor() ([]byte, []int) {
	return file_models_model_http_fastcgi_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPFastcgi) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HTTPFastcgi) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *HTTPFastcgi) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HTTPFastcgi) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *HTTPFastcgi) GetReadTimeoutJSON() []byte {
	if x != nil {
		return x.ReadTimeoutJSON
	}
	return nil
}

func (x *HTTPFastcgi) GetConnTimeoutJSON() []byte {
	if x != nil {
		return x.ConnTimeoutJSON
	}
	return nil
}

func (x *HTTPFastcgi) GetPoolSize() int32 {
	if x != nil {
		return x.PoolSize
	}
	return 0
}

func (x *HTTPFastcgi) GetPathInfoPattern() string {
	if x != nil {
		return x.PathInfoPattern
	}
	return ""
}

var File_models_model_http_fastcgi_proto protoreflect.FileDescriptor

var file_models_model_http_fastcgi_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x66, 0x61, 0x73, 0x74, 0x63, 0x67, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x85, 0x02, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50, 0x46, 0x61,
	0x73, 0x74, 0x63, 0x67, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x65,
	0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_http_fastcgi_proto_rawDescOnce sync.Once
	file_models_model_http_fastcgi_proto_rawDescData = file_models_model_http_fastcgi_proto_rawDesc
)

func file_models_model_http_fastcgi_proto_rawDescGZIP() []byte {
	file_models_model_http_fastcgi_proto_rawDescOnce.Do(func() {
		file_models_model_http_fastcgi_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_http_fastcgi_proto_rawDescData)
	})
	return file_models_model_http_fastcgi_proto_rawDescData
}

var file_models_model_http_fastcgi_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_http_fastcgi_proto_goTypes = []interface{}{
	(*HTTPFastcgi)(nil), // 0: pb.HTTPFastcgi
}
var file_models_model_http_fastcgi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_http_fastcgi_proto_init() }
func file_models_model_http_fastcgi_proto_init() {
	if File_models_model_http_fastcgi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_http_fastcgi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFastcgi); i {
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
			RawDescriptor: file_models_model_http_fastcgi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_http_fastcgi_proto_goTypes,
		DependencyIndexes: file_models_model_http_fastcgi_proto_depIdxs,
		MessageInfos:      file_models_model_http_fastcgi_proto_msgTypes,
	}.Build()
	File_models_model_http_fastcgi_proto = out.File
	file_models_model_http_fastcgi_proto_rawDesc = nil
	file_models_model_http_fastcgi_proto_goTypes = nil
	file_models_model_http_fastcgi_proto_depIdxs = nil
}