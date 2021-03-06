// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: notify.proto

package protocol

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

// 系统通知
type SysNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`       // 推送者
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 推送的消息
	Url     string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`         // 消息的url
	Images  []string `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`   // 消息的图片
}

func (x *SysNotify) Reset() {
	*x = SysNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysNotify) ProtoMessage() {}

func (x *SysNotify) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysNotify.ProtoReflect.Descriptor instead.
func (*SysNotify) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{0}
}

func (x *SysNotify) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SysNotify) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SysNotify) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SysNotify) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

// 广告通知
type AdNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 推送的消息
	Url     string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`         // 消息的url
	Images  []string `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`   // 消息的图片
}

func (x *AdNotify) Reset() {
	*x = AdNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdNotify) ProtoMessage() {}

func (x *AdNotify) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdNotify.ProtoReflect.Descriptor instead.
func (*AdNotify) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{1}
}

func (x *AdNotify) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AdNotify) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AdNotify) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

// vip通知
type VipNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`    // vip消息类型；    0: vip过期  1:开通成功  2：续费成功
	Level  int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`  // vip等级
	Expire string `protobuf:"bytes,3,opt,name=expire,proto3" json:"expire,omitempty"` // 过期时间
	Desc   string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`     // 描述
}

func (x *VipNotify) Reset() {
	*x = VipNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipNotify) ProtoMessage() {}

func (x *VipNotify) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipNotify.ProtoReflect.Descriptor instead.
func (*VipNotify) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{2}
}

func (x *VipNotify) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *VipNotify) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *VipNotify) GetExpire() string {
	if x != nil {
		return x.Expire
	}
	return ""
}

func (x *VipNotify) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// 用户动态通知
type UserNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UserNotify) Reset() {
	*x = UserNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotify) ProtoMessage() {}

func (x *UserNotify) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotify.ProtoReflect.Descriptor instead.
func (*UserNotify) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{3}
}

func (x *UserNotify) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_notify_proto protoreflect.FileDescriptor

var file_notify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x63, 0x0a, 0x09, 0x53, 0x79, 0x73, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4e, 0x0a,
	0x08, 0x41, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x61, 0x0a,
	0x09, 0x56, 0x69, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x22, 0x20, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notify_proto_rawDescOnce sync.Once
	file_notify_proto_rawDescData = file_notify_proto_rawDesc
)

func file_notify_proto_rawDescGZIP() []byte {
	file_notify_proto_rawDescOnce.Do(func() {
		file_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_proto_rawDescData)
	})
	return file_notify_proto_rawDescData
}

var file_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notify_proto_goTypes = []interface{}{
	(*SysNotify)(nil),  // 0: protocol.SysNotify
	(*AdNotify)(nil),   // 1: protocol.AdNotify
	(*VipNotify)(nil),  // 2: protocol.VipNotify
	(*UserNotify)(nil), // 3: protocol.UserNotify
}
var file_notify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notify_proto_init() }
func file_notify_proto_init() {
	if File_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysNotify); i {
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
		file_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdNotify); i {
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
		file_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipNotify); i {
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
		file_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotify); i {
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
			RawDescriptor: file_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notify_proto_goTypes,
		DependencyIndexes: file_notify_proto_depIdxs,
		MessageInfos:      file_notify_proto_msgTypes,
	}.Build()
	File_notify_proto = out.File
	file_notify_proto_rawDesc = nil
	file_notify_proto_goTypes = nil
	file_notify_proto_depIdxs = nil
}
