// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: likenode.proto

package likenode

import (
	common "github.com/heegspace/heegproto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	rescode "github.com/heegspace/heegproto/rescode"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LikesCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Uid  int64             `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Mid  string            `protobuf:"bytes,3,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *LikesCountReq) Reset() {
	*x = LikesCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesCountReq) ProtoMessage() {}

func (x *LikesCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesCountReq.ProtoReflect.Descriptor instead.
func (*LikesCountReq) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{0}
}

func (x *LikesCountReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *LikesCountReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LikesCountReq) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

type LikesCountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Count   int32             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Resmsg  string            `protobuf:"bytes,3,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Extra   map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LikesCountRes) Reset() {
	*x = LikesCountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesCountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesCountRes) ProtoMessage() {}

func (x *LikesCountRes) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesCountRes.ProtoReflect.Descriptor instead.
func (*LikesCountRes) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{1}
}

func (x *LikesCountRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *LikesCountRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LikesCountRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *LikesCountRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type LikesAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Uid  int64             `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Mid  string            `protobuf:"bytes,3,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *LikesAddReq) Reset() {
	*x = LikesAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesAddReq) ProtoMessage() {}

func (x *LikesAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesAddReq.ProtoReflect.Descriptor instead.
func (*LikesAddReq) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{2}
}

func (x *LikesAddReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *LikesAddReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LikesAddReq) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

type LikesAddRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Extra   map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LikesAddRes) Reset() {
	*x = LikesAddRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesAddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesAddRes) ProtoMessage() {}

func (x *LikesAddRes) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesAddRes.ProtoReflect.Descriptor instead.
func (*LikesAddRes) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{3}
}

func (x *LikesAddRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *LikesAddRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *LikesAddRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type LikesListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Uid  int64             `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Mid  string            `protobuf:"bytes,3,opt,name=mid,proto3" json:"mid,omitempty"`
	Page int32             `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Size int32             `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *LikesListReq) Reset() {
	*x = LikesListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesListReq) ProtoMessage() {}

func (x *LikesListReq) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesListReq.ProtoReflect.Descriptor instead.
func (*LikesListReq) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{4}
}

func (x *LikesListReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *LikesListReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LikesListReq) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *LikesListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LikesListReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Likes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *Likes) Reset() {
	*x = Likes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Likes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Likes) ProtoMessage() {}

func (x *Likes) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Likes.ProtoReflect.Descriptor instead.
func (*Likes) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{5}
}

func (x *Likes) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Likes) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

type LikesListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Data    []*Likes          `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Extra   map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LikesListRes) Reset() {
	*x = LikesListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likenode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesListRes) ProtoMessage() {}

func (x *LikesListRes) ProtoReflect() protoreflect.Message {
	mi := &file_likenode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesListRes.ProtoReflect.Descriptor instead.
func (*LikesListRes) Descriptor() ([]byte, []int) {
	return file_likenode_proto_rawDescGZIP(), []int{6}
}

func (x *LikesListRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *LikesListRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *LikesListRes) GetData() []*Likes {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LikesListRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_likenode_proto protoreflect.FileDescriptor

var file_likenode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x69, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x0f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12,
	0x3a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x0d, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61,
	0x64, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69,
	0x64, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38,
	0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x36,
	0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0xeb, 0x01, 0x0a, 0x0e, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xdb, 0x01, 0x0a, 0x10, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x41, 0x64, 0x64, 0x12, 0x17, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61, 0x64, 0x64,
	0x5f, 0x72, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_likenode_proto_rawDescOnce sync.Once
	file_likenode_proto_rawDescData = file_likenode_proto_rawDesc
)

func file_likenode_proto_rawDescGZIP() []byte {
	file_likenode_proto_rawDescOnce.Do(func() {
		file_likenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_likenode_proto_rawDescData)
	})
	return file_likenode_proto_rawDescData
}

var file_likenode_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_likenode_proto_goTypes = []interface{}{
	(*LikesCountReq)(nil),    // 0: likenode.likes_count_req
	(*LikesCountRes)(nil),    // 1: likenode.likes_count_res
	(*LikesAddReq)(nil),      // 2: likenode.likes_add_req
	(*LikesAddRes)(nil),      // 3: likenode.likes_add_res
	(*LikesListReq)(nil),     // 4: likenode.likes_list_req
	(*Likes)(nil),            // 5: likenode.likes
	(*LikesListRes)(nil),     // 6: likenode.likes_list_res
	nil,                      // 7: likenode.likes_count_res.ExtraEntry
	nil,                      // 8: likenode.likes_add_res.ExtraEntry
	nil,                      // 9: likenode.likes_list_res.ExtraEntry
	(*common.Authorize)(nil), // 10: common.authorize
	(rescode.Code)(0),        // 11: rescode.code
}
var file_likenode_proto_depIdxs = []int32{
	10, // 0: likenode.likes_count_req.auth:type_name -> common.authorize
	11, // 1: likenode.likes_count_res.rescode:type_name -> rescode.code
	7,  // 2: likenode.likes_count_res.extra:type_name -> likenode.likes_count_res.ExtraEntry
	10, // 3: likenode.likes_add_req.auth:type_name -> common.authorize
	11, // 4: likenode.likes_add_res.rescode:type_name -> rescode.code
	8,  // 5: likenode.likes_add_res.extra:type_name -> likenode.likes_add_res.ExtraEntry
	10, // 6: likenode.likes_list_req.auth:type_name -> common.authorize
	11, // 7: likenode.likes_list_res.rescode:type_name -> rescode.code
	5,  // 8: likenode.likes_list_res.data:type_name -> likenode.likes
	9,  // 9: likenode.likes_list_res.extra:type_name -> likenode.likes_list_res.ExtraEntry
	0,  // 10: likenode.likenode_service.likesCount:input_type -> likenode.likes_count_req
	2,  // 11: likenode.likenode_service.likesAdd:input_type -> likenode.likes_add_req
	4,  // 12: likenode.likenode_service.likesList:input_type -> likenode.likes_list_req
	1,  // 13: likenode.likenode_service.likesCount:output_type -> likenode.likes_count_res
	3,  // 14: likenode.likenode_service.likesAdd:output_type -> likenode.likes_add_res
	6,  // 15: likenode.likenode_service.likesList:output_type -> likenode.likes_list_res
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_likenode_proto_init() }
func file_likenode_proto_init() {
	if File_likenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_likenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesCountReq); i {
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
		file_likenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesCountRes); i {
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
		file_likenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesAddReq); i {
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
		file_likenode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesAddRes); i {
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
		file_likenode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesListReq); i {
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
		file_likenode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Likes); i {
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
		file_likenode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesListRes); i {
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
			RawDescriptor: file_likenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_likenode_proto_goTypes,
		DependencyIndexes: file_likenode_proto_depIdxs,
		MessageInfos:      file_likenode_proto_msgTypes,
	}.Build()
	File_likenode_proto = out.File
	file_likenode_proto_rawDesc = nil
	file_likenode_proto_goTypes = nil
	file_likenode_proto_depIdxs = nil
}