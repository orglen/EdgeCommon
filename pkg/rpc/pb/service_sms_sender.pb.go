// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_sms_sender.proto

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

// 发送短信
type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile     string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`         // 手机号
	Body       string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`             // 内容
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`             // 验证码
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`             // 渠道类型：webHook ...
	ParamsJSON []byte `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"` // 参数
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_sms_sender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_sms_sender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_service_sms_sender_proto_rawDescGZIP(), []int{0}
}

func (x *SendSMSRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SendSMSRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SendSMSRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendSMSRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendSMSRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

type SendSMSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk   bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`    // 是否成功
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"` // 发送返回内容，只有失败时才会有数据
}

func (x *SendSMSResponse) Reset() {
	*x = SendSMSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_sms_sender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSResponse) ProtoMessage() {}

func (x *SendSMSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_sms_sender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSResponse.ProtoReflect.Descriptor instead.
func (*SendSMSResponse) Descriptor() ([]byte, []int) {
	return file_service_sms_sender_proto_rawDescGZIP(), []int{1}
}

func (x *SendSMSResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *SendSMSResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_service_sms_sender_proto protoreflect.FileDescriptor

var file_service_sms_sender_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x6d, 0x73, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x84,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x3d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x46, 0x0a, 0x10, 0x53, 0x4d, 0x53, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64,
	0x53, 0x4d, 0x53, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_sms_sender_proto_rawDescOnce sync.Once
	file_service_sms_sender_proto_rawDescData = file_service_sms_sender_proto_rawDesc
)

func file_service_sms_sender_proto_rawDescGZIP() []byte {
	file_service_sms_sender_proto_rawDescOnce.Do(func() {
		file_service_sms_sender_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_sms_sender_proto_rawDescData)
	})
	return file_service_sms_sender_proto_rawDescData
}

var file_service_sms_sender_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_sms_sender_proto_goTypes = []interface{}{
	(*SendSMSRequest)(nil),  // 0: pb.SendSMSRequest
	(*SendSMSResponse)(nil), // 1: pb.SendSMSResponse
}
var file_service_sms_sender_proto_depIdxs = []int32{
	0, // 0: pb.SMSSenderService.sendSMS:input_type -> pb.SendSMSRequest
	1, // 1: pb.SMSSenderService.sendSMS:output_type -> pb.SendSMSResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_sms_sender_proto_init() }
func file_service_sms_sender_proto_init() {
	if File_service_sms_sender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_sms_sender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
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
		file_service_sms_sender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSResponse); i {
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
			RawDescriptor: file_service_sms_sender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_sms_sender_proto_goTypes,
		DependencyIndexes: file_service_sms_sender_proto_depIdxs,
		MessageInfos:      file_service_sms_sender_proto_msgTypes,
	}.Build()
	File_service_sms_sender_proto = out.File
	file_service_sms_sender_proto_rawDesc = nil
	file_service_sms_sender_proto_goTypes = nil
	file_service_sms_sender_proto_depIdxs = nil
}