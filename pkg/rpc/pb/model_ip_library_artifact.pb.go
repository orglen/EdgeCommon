// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_ip_library_artifact.proto

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

type IPLibraryArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileId    int64  `protobuf:"varint,2,opt,name=fileId,proto3" json:"fileId,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MetaJSON  []byte `protobuf:"bytes,4,opt,name=metaJSON,proto3" json:"metaJSON,omitempty"`
	IsPublic  bool   `protobuf:"varint,5,opt,name=isPublic,proto3" json:"isPublic,omitempty"` // 是否公开
	Name      string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Code      string `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	File      *File  `protobuf:"bytes,30,opt,name=file,proto3" json:"file,omitempty"` // 文件信息
}

func (x *IPLibraryArtifact) Reset() {
	*x = IPLibraryArtifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ip_library_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPLibraryArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPLibraryArtifact) ProtoMessage() {}

func (x *IPLibraryArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ip_library_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPLibraryArtifact.ProtoReflect.Descriptor instead.
func (*IPLibraryArtifact) Descriptor() ([]byte, []int) {
	return file_models_model_ip_library_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *IPLibraryArtifact) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IPLibraryArtifact) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *IPLibraryArtifact) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *IPLibraryArtifact) GetMetaJSON() []byte {
	if x != nil {
		return x.MetaJSON
	}
	return nil
}

func (x *IPLibraryArtifact) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *IPLibraryArtifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IPLibraryArtifact) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *IPLibraryArtifact) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_models_model_ip_library_artifact_proto protoreflect.FileDescriptor

var file_models_model_ip_library_artifact_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69,
	0x70, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ip_library_artifact_proto_rawDescOnce sync.Once
	file_models_model_ip_library_artifact_proto_rawDescData = file_models_model_ip_library_artifact_proto_rawDesc
)

func file_models_model_ip_library_artifact_proto_rawDescGZIP() []byte {
	file_models_model_ip_library_artifact_proto_rawDescOnce.Do(func() {
		file_models_model_ip_library_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ip_library_artifact_proto_rawDescData)
	})
	return file_models_model_ip_library_artifact_proto_rawDescData
}

var file_models_model_ip_library_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ip_library_artifact_proto_goTypes = []interface{}{
	(*IPLibraryArtifact)(nil), // 0: pb.IPLibraryArtifact
	(*File)(nil),              // 1: pb.File
}
var file_models_model_ip_library_artifact_proto_depIdxs = []int32{
	1, // 0: pb.IPLibraryArtifact.file:type_name -> pb.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_ip_library_artifact_proto_init() }
func file_models_model_ip_library_artifact_proto_init() {
	if File_models_model_ip_library_artifact_proto != nil {
		return
	}
	file_models_model_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ip_library_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPLibraryArtifact); i {
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
			RawDescriptor: file_models_model_ip_library_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ip_library_artifact_proto_goTypes,
		DependencyIndexes: file_models_model_ip_library_artifact_proto_depIdxs,
		MessageInfos:      file_models_model_ip_library_artifact_proto_msgTypes,
	}.Build()
	File_models_model_ip_library_artifact_proto = out.File
	file_models_model_ip_library_artifact_proto_rawDesc = nil
	file_models_model_ip_library_artifact_proto_goTypes = nil
	file_models_model_ip_library_artifact_proto_depIdxs = nil
}