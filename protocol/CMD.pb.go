// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: CMD.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 流协议号定义
type Stream_Procol_Num int32

const (
	// 心跳协议
	Stream_Procol_Num_HEART Stream_Procol_Num = 0
	// 接受聊天信息
	Stream_Procol_Num_CHAT_RECV Stream_Procol_Num = 50
	// 发送聊天信息
	Stream_Procol_Num_CHAT_SEND Stream_Procol_Num = 51
	// 系统消息通知
	Stream_Procol_Num_NOTIFY_SYS Stream_Procol_Num = 100
	// 广告消息通知
	Stream_Procol_Num_NOTIFY_AD Stream_Procol_Num = 101
	// 用户vip通知
	Stream_Procol_Num_NOTIFY_VIP Stream_Procol_Num = 102
	// 用户动态通知
	Stream_Procol_Num_NOTIFY_USER Stream_Procol_Num = 103
)

var Stream_Procol_Num_name = map[int32]string{
	0:   "HEART",
	50:  "CHAT_RECV",
	51:  "CHAT_SEND",
	100: "NOTIFY_SYS",
	101: "NOTIFY_AD",
	102: "NOTIFY_VIP",
	103: "NOTIFY_USER",
}

var Stream_Procol_Num_value = map[string]int32{
	"HEART":       0,
	"CHAT_RECV":   50,
	"CHAT_SEND":   51,
	"NOTIFY_SYS":  100,
	"NOTIFY_AD":   101,
	"NOTIFY_VIP":  102,
	"NOTIFY_USER": 103,
}

func (x Stream_Procol_Num) String() string {
	return proto.EnumName(Stream_Procol_Num_name, int32(x))
}

func (Stream_Procol_Num) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51ccfcfc406062df, []int{0}
}

func init() {
	proto.RegisterEnum("protocol.Stream_Procol_Num", Stream_Procol_Num_name, Stream_Procol_Num_value)
}

func init() { proto.RegisterFile("CMD.proto", fileDescriptor_51ccfcfc406062df) }

var fileDescriptor_51ccfcfc406062df = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x74, 0xf6, 0x75, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x5a, 0x35, 0x5c, 0x82,
	0xc1, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0xf1, 0x01, 0x45, 0x20, 0x81, 0x78, 0xbf, 0xd2, 0x5c, 0x21,
	0x4e, 0x2e, 0x56, 0x0f, 0x57, 0xc7, 0xa0, 0x10, 0x01, 0x06, 0x21, 0x5e, 0x2e, 0x4e, 0x67, 0x0f,
	0xc7, 0x90, 0xf8, 0x20, 0x57, 0xe7, 0x30, 0x01, 0x23, 0x38, 0x37, 0xd8, 0xd5, 0xcf, 0x45, 0xc0,
	0x58, 0x88, 0x8f, 0x8b, 0xcb, 0xcf, 0x3f, 0xc4, 0xd3, 0x2d, 0x32, 0x3e, 0x38, 0x32, 0x58, 0x20,
	0x05, 0x24, 0x0d, 0xe5, 0x3b, 0xba, 0x08, 0xa4, 0x22, 0x49, 0x87, 0x79, 0x06, 0x08, 0xa4, 0x09,
	0xf1, 0x73, 0x71, 0x43, 0xf9, 0xa1, 0xc1, 0xae, 0x41, 0x02, 0xe9, 0x4e, 0x02, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c,
	0x60, 0x97, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x64, 0xc2, 0x50, 0xad, 0x00, 0x00,
	0x00,
}