// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_user_traffic_package.proto

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

// 用户流量包
type UserTrafficPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                        int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId                    int64           `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	TrafficPackageId          int64           `protobuf:"varint,3,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	TotalBytes                int64           `protobuf:"varint,4,opt,name=totalBytes,proto3" json:"totalBytes,omitempty"`
	UsedBytes                 int64           `protobuf:"varint,5,opt,name=usedBytes,proto3" json:"usedBytes,omitempty"`
	NodeRegionId              int64           `protobuf:"varint,6,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId    int64           `protobuf:"varint,7,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	TrafficPackagePeriodCount int32           `protobuf:"varint,8,opt,name=trafficPackagePeriodCount,proto3" json:"trafficPackagePeriodCount,omitempty"`
	TrafficPackagePeriodUnit  string          `protobuf:"bytes,9,opt,name=trafficPackagePeriodUnit,proto3" json:"trafficPackagePeriodUnit,omitempty"`
	DayFrom                   string          `protobuf:"bytes,10,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`
	DayTo                     string          `protobuf:"bytes,11,opt,name=dayTo,proto3" json:"dayTo,omitempty"`
	CreatedAt                 int64           `protobuf:"varint,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	TrafficPackage            *TrafficPackage `protobuf:"bytes,30,opt,name=trafficPackage,proto3" json:"trafficPackage,omitempty"`
	NodeRegion                *NodeRegion     `protobuf:"bytes,31,opt,name=nodeRegion,proto3" json:"nodeRegion,omitempty"`
	User                      *User           `protobuf:"bytes,32,opt,name=user,proto3" json:"user,omitempty"`
	CanDelete                 bool            `protobuf:"varint,33,opt,name=canDelete,proto3" json:"canDelete,omitempty"`
}

func (x *UserTrafficPackage) Reset() {
	*x = UserTrafficPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_traffic_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTrafficPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTrafficPackage) ProtoMessage() {}

func (x *UserTrafficPackage) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_traffic_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTrafficPackage.ProtoReflect.Descriptor instead.
func (*UserTrafficPackage) Descriptor() ([]byte, []int) {
	return file_models_model_user_traffic_package_proto_rawDescGZIP(), []int{0}
}

func (x *UserTrafficPackage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTrafficPackage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserTrafficPackage) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *UserTrafficPackage) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *UserTrafficPackage) GetUsedBytes() int64 {
	if x != nil {
		return x.UsedBytes
	}
	return 0
}

func (x *UserTrafficPackage) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *UserTrafficPackage) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *UserTrafficPackage) GetTrafficPackagePeriodCount() int32 {
	if x != nil {
		return x.TrafficPackagePeriodCount
	}
	return 0
}

func (x *UserTrafficPackage) GetTrafficPackagePeriodUnit() string {
	if x != nil {
		return x.TrafficPackagePeriodUnit
	}
	return ""
}

func (x *UserTrafficPackage) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *UserTrafficPackage) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *UserTrafficPackage) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserTrafficPackage) GetTrafficPackage() *TrafficPackage {
	if x != nil {
		return x.TrafficPackage
	}
	return nil
}

func (x *UserTrafficPackage) GetNodeRegion() *NodeRegion {
	if x != nil {
		return x.NodeRegion
	}
	return nil
}

func (x *UserTrafficPackage) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserTrafficPackage) GetCanDelete() bool {
	if x != nil {
		return x.CanDelete
	}
	return false
}

var File_models_model_user_traffic_package_proto protoreflect.FileDescriptor

var file_models_model_user_traffic_package_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x04, 0x0a, 0x12, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x19, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x19, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a,
	0x18, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_traffic_package_proto_rawDescOnce sync.Once
	file_models_model_user_traffic_package_proto_rawDescData = file_models_model_user_traffic_package_proto_rawDesc
)

func file_models_model_user_traffic_package_proto_rawDescGZIP() []byte {
	file_models_model_user_traffic_package_proto_rawDescOnce.Do(func() {
		file_models_model_user_traffic_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_traffic_package_proto_rawDescData)
	})
	return file_models_model_user_traffic_package_proto_rawDescData
}

var file_models_model_user_traffic_package_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_traffic_package_proto_goTypes = []interface{}{
	(*UserTrafficPackage)(nil), // 0: pb.UserTrafficPackage
	(*TrafficPackage)(nil),     // 1: pb.TrafficPackage
	(*NodeRegion)(nil),         // 2: pb.NodeRegion
	(*User)(nil),               // 3: pb.User
}
var file_models_model_user_traffic_package_proto_depIdxs = []int32{
	1, // 0: pb.UserTrafficPackage.trafficPackage:type_name -> pb.TrafficPackage
	2, // 1: pb.UserTrafficPackage.nodeRegion:type_name -> pb.NodeRegion
	3, // 2: pb.UserTrafficPackage.user:type_name -> pb.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_user_traffic_package_proto_init() }
func file_models_model_user_traffic_package_proto_init() {
	if File_models_model_user_traffic_package_proto != nil {
		return
	}
	file_models_model_node_region_proto_init()
	file_models_model_traffic_package_proto_init()
	file_models_model_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_traffic_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTrafficPackage); i {
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
			RawDescriptor: file_models_model_user_traffic_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_traffic_package_proto_goTypes,
		DependencyIndexes: file_models_model_user_traffic_package_proto_depIdxs,
		MessageInfos:      file_models_model_user_traffic_package_proto_msgTypes,
	}.Build()
	File_models_model_user_traffic_package_proto = out.File
	file_models_model_user_traffic_package_proto_rawDesc = nil
	file_models_model_user_traffic_package_proto_goTypes = nil
	file_models_model_user_traffic_package_proto_depIdxs = nil
}
