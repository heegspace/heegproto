// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: macipnode.proto

package macipnode

import (
	common "github.com/heegspace/heegproto/common"
	rescode "github.com/heegspace/heegproto/rescode"
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

type IpToAddressReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth  *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Ip    string            `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Extra map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IpToAddressReq) Reset() {
	*x = IpToAddressReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_macipnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpToAddressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpToAddressReq) ProtoMessage() {}

func (x *IpToAddressReq) ProtoReflect() protoreflect.Message {
	mi := &file_macipnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpToAddressReq.ProtoReflect.Descriptor instead.
func (*IpToAddressReq) Descriptor() ([]byte, []int) {
	return file_macipnode_proto_rawDescGZIP(), []int{0}
}

func (x *IpToAddressReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *IpToAddressReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IpToAddressReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type AddressItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip           string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Province     string `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City         string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	District     string `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Organization string `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
	Isp          string `protobuf:"bytes,7,opt,name=isp,proto3" json:"isp,omitempty"`
	CountryCode  string `protobuf:"bytes,8,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *AddressItem) Reset() {
	*x = AddressItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_macipnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressItem) ProtoMessage() {}

func (x *AddressItem) ProtoReflect() protoreflect.Message {
	mi := &file_macipnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressItem.ProtoReflect.Descriptor instead.
func (*AddressItem) Descriptor() ([]byte, []int) {
	return file_macipnode_proto_rawDescGZIP(), []int{1}
}

func (x *AddressItem) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AddressItem) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AddressItem) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *AddressItem) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressItem) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *AddressItem) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *AddressItem) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *AddressItem) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type IpToAddressRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Address *AddressItem      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Extra   map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IpToAddressRes) Reset() {
	*x = IpToAddressRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_macipnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpToAddressRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpToAddressRes) ProtoMessage() {}

func (x *IpToAddressRes) ProtoReflect() protoreflect.Message {
	mi := &file_macipnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpToAddressRes.ProtoReflect.Descriptor instead.
func (*IpToAddressRes) Descriptor() ([]byte, []int) {
	return file_macipnode_proto_rawDescGZIP(), []int{2}
}

func (x *IpToAddressRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *IpToAddressRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *IpToAddressRes) GetAddress() *AddressItem {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *IpToAddressRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_macipnode_proto protoreflect.FileDescriptor

var file_macipnode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x72, 0x65,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0e, 0x49, 0x70, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x3a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x70,
	0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38,
	0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x0e, 0x49, 0x70, 0x54, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d,
	0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x70, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x59, 0x0a, 0x10, 0x4d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x49, 0x70, 0x54, 0x6f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x49, 0x70, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x63, 0x69, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x70,
	0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_macipnode_proto_rawDescOnce sync.Once
	file_macipnode_proto_rawDescData = file_macipnode_proto_rawDesc
)

func file_macipnode_proto_rawDescGZIP() []byte {
	file_macipnode_proto_rawDescOnce.Do(func() {
		file_macipnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_macipnode_proto_rawDescData)
	})
	return file_macipnode_proto_rawDescData
}

var file_macipnode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_macipnode_proto_goTypes = []interface{}{
	(*IpToAddressReq)(nil),   // 0: macipnode.IpToAddressReq
	(*AddressItem)(nil),      // 1: macipnode.AddressItem
	(*IpToAddressRes)(nil),   // 2: macipnode.IpToAddressRes
	nil,                      // 3: macipnode.IpToAddressReq.ExtraEntry
	nil,                      // 4: macipnode.IpToAddressRes.ExtraEntry
	(*common.Authorize)(nil), // 5: common.Authorize
	(rescode.Code)(0),        // 6: rescode.Code
}
var file_macipnode_proto_depIdxs = []int32{
	5, // 0: macipnode.IpToAddressReq.auth:type_name -> common.Authorize
	3, // 1: macipnode.IpToAddressReq.extra:type_name -> macipnode.IpToAddressReq.ExtraEntry
	6, // 2: macipnode.IpToAddressRes.rescode:type_name -> rescode.Code
	1, // 3: macipnode.IpToAddressRes.address:type_name -> macipnode.AddressItem
	4, // 4: macipnode.IpToAddressRes.extra:type_name -> macipnode.IpToAddressRes.ExtraEntry
	0, // 5: macipnode.MacipnodeService.IpToAddress:input_type -> macipnode.IpToAddressReq
	2, // 6: macipnode.MacipnodeService.IpToAddress:output_type -> macipnode.IpToAddressRes
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_macipnode_proto_init() }
func file_macipnode_proto_init() {
	if File_macipnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_macipnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpToAddressReq); i {
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
		file_macipnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressItem); i {
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
		file_macipnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpToAddressRes); i {
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
			RawDescriptor: file_macipnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_macipnode_proto_goTypes,
		DependencyIndexes: file_macipnode_proto_depIdxs,
		MessageInfos:      file_macipnode_proto_msgTypes,
	}.Build()
	File_macipnode_proto = out.File
	file_macipnode_proto_rawDesc = nil
	file_macipnode_proto_goTypes = nil
	file_macipnode_proto_depIdxs = nil
}
