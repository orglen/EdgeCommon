// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_server_bandwidth_stat.proto

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

// 带宽统计数据
type ServerBandwidthStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                // ID
	UserId                    int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`                                        // 用户ID
	ServerId                  int64  `protobuf:"varint,3,opt,name=serverId,proto3" json:"serverId,omitempty"`                                    //服务ID
	Day                       string `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`                                               // 日期，格式YYYYMMDD
	TimeAt                    string `protobuf:"bytes,5,opt,name=timeAt,proto3" json:"timeAt,omitempty"`                                         // 时间，格式HHII
	Bytes                     int64  `protobuf:"varint,6,opt,name=bytes,proto3" json:"bytes,omitempty"`                                          // 峰值带宽字节
	TotalBytes                int64  `protobuf:"varint,9,opt,name=totalBytes,proto3" json:"totalBytes,omitempty"`                                // 总流量
	Bits                      int64  `protobuf:"varint,7,opt,name=bits,proto3" json:"bits,omitempty"`                                            // 峰值带宽比特
	NodeRegionId              int64  `protobuf:"varint,8,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`                            // 节点所在区域ID
	CachedBytes               int64  `protobuf:"varint,10,opt,name=cachedBytes,proto3" json:"cachedBytes,omitempty"`                             // 总缓存流量
	AttackBytes               int64  `protobuf:"varint,11,opt,name=attackBytes,proto3" json:"attackBytes,omitempty"`                             // 总攻击流量
	CountRequests             int64  `protobuf:"varint,12,opt,name=countRequests,proto3" json:"countRequests,omitempty"`                         // 总请求数
	CountCachedRequests       int64  `protobuf:"varint,13,opt,name=countCachedRequests,proto3" json:"countCachedRequests,omitempty"`             // 总缓存请求数
	CountAttackRequests       int64  `protobuf:"varint,14,opt,name=countAttackRequests,proto3" json:"countAttackRequests,omitempty"`             // 总攻击请求数
	UserPlanId                int64  `protobuf:"varint,15,opt,name=userPlanId,proto3" json:"userPlanId,omitempty"`                               // 绑定的用户套餐ID
	CountWebsocketConnections int64  `protobuf:"varint,16,opt,name=countWebsocketConnections,proto3" json:"countWebsocketConnections,omitempty"` // Websocket连接数
	CountIPs                  int64  `protobuf:"varint,17,opt,name=countIPs,proto3" json:"countIPs,omitempty"`                                   // 总IP数
}

func (x *ServerBandwidthStat) Reset() {
	*x = ServerBandwidthStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_server_bandwidth_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerBandwidthStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerBandwidthStat) ProtoMessage() {}

func (x *ServerBandwidthStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_server_bandwidth_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerBandwidthStat.ProtoReflect.Descriptor instead.
func (*ServerBandwidthStat) Descriptor() ([]byte, []int) {
	return file_models_model_server_bandwidth_stat_proto_rawDescGZIP(), []int{0}
}

func (x *ServerBandwidthStat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServerBandwidthStat) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ServerBandwidthStat) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *ServerBandwidthStat) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ServerBandwidthStat) GetTimeAt() string {
	if x != nil {
		return x.TimeAt
	}
	return ""
}

func (x *ServerBandwidthStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *ServerBandwidthStat) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *ServerBandwidthStat) GetBits() int64 {
	if x != nil {
		return x.Bits
	}
	return 0
}

func (x *ServerBandwidthStat) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *ServerBandwidthStat) GetCachedBytes() int64 {
	if x != nil {
		return x.CachedBytes
	}
	return 0
}

func (x *ServerBandwidthStat) GetAttackBytes() int64 {
	if x != nil {
		return x.AttackBytes
	}
	return 0
}

func (x *ServerBandwidthStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

func (x *ServerBandwidthStat) GetCountCachedRequests() int64 {
	if x != nil {
		return x.CountCachedRequests
	}
	return 0
}

func (x *ServerBandwidthStat) GetCountAttackRequests() int64 {
	if x != nil {
		return x.CountAttackRequests
	}
	return 0
}

func (x *ServerBandwidthStat) GetUserPlanId() int64 {
	if x != nil {
		return x.UserPlanId
	}
	return 0
}

func (x *ServerBandwidthStat) GetCountWebsocketConnections() int64 {
	if x != nil {
		return x.CountWebsocketConnections
	}
	return 0
}

func (x *ServerBandwidthStat) GetCountIPs() int64 {
	if x != nil {
		return x.CountIPs
	}
	return 0
}

var File_models_model_server_bandwidth_stat_proto protoreflect.FileDescriptor

var file_models_model_server_bandwidth_stat_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb9,
	0x04, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x69, 0x6d, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69,
	0x6d, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x13,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x30,
	0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x12, 0x3c, 0x0a, 0x19, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x19, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x50, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x50, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_server_bandwidth_stat_proto_rawDescOnce sync.Once
	file_models_model_server_bandwidth_stat_proto_rawDescData = file_models_model_server_bandwidth_stat_proto_rawDesc
)

func file_models_model_server_bandwidth_stat_proto_rawDescGZIP() []byte {
	file_models_model_server_bandwidth_stat_proto_rawDescOnce.Do(func() {
		file_models_model_server_bandwidth_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_server_bandwidth_stat_proto_rawDescData)
	})
	return file_models_model_server_bandwidth_stat_proto_rawDescData
}

var file_models_model_server_bandwidth_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_server_bandwidth_stat_proto_goTypes = []interface{}{
	(*ServerBandwidthStat)(nil), // 0: pb.ServerBandwidthStat
}
var file_models_model_server_bandwidth_stat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_server_bandwidth_stat_proto_init() }
func file_models_model_server_bandwidth_stat_proto_init() {
	if File_models_model_server_bandwidth_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_server_bandwidth_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerBandwidthStat); i {
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
			RawDescriptor: file_models_model_server_bandwidth_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_server_bandwidth_stat_proto_goTypes,
		DependencyIndexes: file_models_model_server_bandwidth_stat_proto_depIdxs,
		MessageInfos:      file_models_model_server_bandwidth_stat_proto_msgTypes,
	}.Build()
	File_models_model_server_bandwidth_stat_proto = out.File
	file_models_model_server_bandwidth_stat_proto_rawDesc = nil
	file_models_model_server_bandwidth_stat_proto_goTypes = nil
	file_models_model_server_bandwidth_stat_proto_depIdxs = nil
}
