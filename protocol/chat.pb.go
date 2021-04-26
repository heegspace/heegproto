// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// 聊天数据类型
type Chat_Type int32

const (
	Chat_Type_NONE          Chat_Type = 0
	Chat_Type_TEXT          Chat_Type = 1
	Chat_Type_IMAGE         Chat_Type = 2
	Chat_Type_IMAGE_URL     Chat_Type = 3
	Chat_Type_IMAGE_BASE64  Chat_Type = 4
	Chat_Type_SOUND         Chat_Type = 5
	Chat_Type_SOUND_URL     Chat_Type = 6
	Chat_Type_VIDEO         Chat_Type = 7
	Chat_Type_VIDEO_URL     Chat_Type = 8
	Chat_Type_HTML          Chat_Type = 9
	Chat_Type_HTML_URL      Chat_Type = 10
	Chat_Type_EMOJII        Chat_Type = 11
	Chat_Type_EMOJII_URL    Chat_Type = 12
	Chat_Type_EMOJII_BASE64 Chat_Type = 13
)

var Chat_Type_name = map[int32]string{
	0:  "NONE",
	1:  "TEXT",
	2:  "IMAGE",
	3:  "IMAGE_URL",
	4:  "IMAGE_BASE64",
	5:  "SOUND",
	6:  "SOUND_URL",
	7:  "VIDEO",
	8:  "VIDEO_URL",
	9:  "HTML",
	10: "HTML_URL",
	11: "EMOJII",
	12: "EMOJII_URL",
	13: "EMOJII_BASE64",
}

var Chat_Type_value = map[string]int32{
	"NONE":          0,
	"TEXT":          1,
	"IMAGE":         2,
	"IMAGE_URL":     3,
	"IMAGE_BASE64":  4,
	"SOUND":         5,
	"SOUND_URL":     6,
	"VIDEO":         7,
	"VIDEO_URL":     8,
	"HTML":          9,
	"HTML_URL":      10,
	"EMOJII":        11,
	"EMOJII_URL":    12,
	"EMOJII_BASE64": 13,
}

func (x Chat_Type) String() string {
	return proto.EnumName(Chat_Type_name, int32(x))
}

func (Chat_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

// 发送消息
type SendChat struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Touid                int64    `protobuf:"varint,2,opt,name=touid,proto3" json:"touid,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendChat) Reset()         { *m = SendChat{} }
func (m *SendChat) String() string { return proto.CompactTextString(m) }
func (*SendChat) ProtoMessage()    {}
func (*SendChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *SendChat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendChat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendChat.Merge(m, src)
}
func (m *SendChat) XXX_Size() int {
	return m.Size()
}
func (m *SendChat) XXX_DiscardUnknown() {
	xxx_messageInfo_SendChat.DiscardUnknown(m)
}

var xxx_messageInfo_SendChat proto.InternalMessageInfo

func (m *SendChat) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SendChat) GetTouid() int64 {
	if m != nil {
		return m.Touid
	}
	return 0
}

func (m *SendChat) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendChat) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 接收消息
type RecvChat struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Fromuid              int64    `protobuf:"varint,2,opt,name=fromuid,proto3" json:"fromuid,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecvChat) Reset()         { *m = RecvChat{} }
func (m *RecvChat) String() string { return proto.CompactTextString(m) }
func (*RecvChat) ProtoMessage()    {}
func (*RecvChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *RecvChat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecvChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecvChat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecvChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvChat.Merge(m, src)
}
func (m *RecvChat) XXX_Size() int {
	return m.Size()
}
func (m *RecvChat) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvChat.DiscardUnknown(m)
}

var xxx_messageInfo_RecvChat proto.InternalMessageInfo

func (m *RecvChat) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RecvChat) GetFromuid() int64 {
	if m != nil {
		return m.Fromuid
	}
	return 0
}

func (m *RecvChat) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RecvChat) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.Chat_Type", Chat_Type_name, Chat_Type_value)
	proto.RegisterType((*SendChat)(nil), "protocol.SendChat")
	proto.RegisterType((*RecvChat)(nil), "protocol.RecvChat")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x18, 0x84, 0xd9, 0xd2, 0x42, 0xfb, 0x5b, 0xc8, 0xfa, 0xc7, 0x43, 0x4f, 0x4d, 0xc3, 0xa9, 0xf1,
	0xe0, 0x45, 0xe3, 0x1d, 0x64, 0xa3, 0x35, 0x40, 0x93, 0xa5, 0x18, 0x6e, 0x58, 0x01, 0x83, 0x89,
	0xba, 0x84, 0xac, 0x26, 0xbe, 0x89, 0x4f, 0x64, 0x3c, 0xfa, 0x08, 0x06, 0x5f, 0xc4, 0xfc, 0x3f,
	0x35, 0x5e, 0xf4, 0xb4, 0xdf, 0xec, 0xcc, 0xce, 0x1c, 0x16, 0x60, 0xbe, 0x2a, 0xed, 0xd1, 0x7a,
	0x63, 0xac, 0x41, 0x9f, 0x8f, 0xb9, 0xb9, 0xef, 0x4c, 0xc1, 0x1f, 0x2f, 0x1f, 0x17, 0x67, 0xab,
	0xd2, 0x22, 0x82, 0x6b, 0x5f, 0xd6, 0xcb, 0x48, 0x24, 0x22, 0xf5, 0x34, 0x33, 0x1e, 0x80, 0x67,
	0xcd, 0xd3, 0xdd, 0x22, 0x72, 0x12, 0x91, 0xd6, 0xf5, 0x4e, 0x60, 0x1b, 0x1c, 0x6b, 0xa2, 0x7a,
	0x22, 0xd2, 0x40, 0x3b, 0xd6, 0xd0, 0xcb, 0x45, 0x69, 0xcb, 0xc8, 0x4d, 0x44, 0x1a, 0x6a, 0xe6,
	0xce, 0x35, 0xf8, 0x7a, 0x39, 0x7f, 0xfe, 0xb7, 0x39, 0x82, 0xe6, 0xed, 0xc6, 0x3c, 0xfc, 0x76,
	0xff, 0x48, 0x4a, 0x13, 0x56, 0xfd, 0xcc, 0x7f, 0x2d, 0x1c, 0xbe, 0x09, 0x08, 0xa8, 0x7e, 0x56,
	0x50, 0x9f, 0x0f, 0xee, 0x28, 0x1f, 0x29, 0x59, 0x23, 0x2a, 0xd4, 0xb4, 0x90, 0x02, 0x03, 0xf0,
	0xb2, 0x61, 0xf7, 0x5c, 0x49, 0x07, 0x5b, 0x10, 0x30, 0xce, 0x26, 0x7a, 0x20, 0xeb, 0x28, 0x21,
	0xdc, 0xc9, 0x5e, 0x77, 0xac, 0x4e, 0x4f, 0xa4, 0x4b, 0xd9, 0x71, 0x3e, 0x19, 0xf5, 0xa5, 0x47,
	0x59, 0x46, 0xce, 0x36, 0xc8, 0xb9, 0xca, 0xfa, 0x2a, 0x97, 0x4d, 0x72, 0x18, 0xd9, 0xf1, 0x69,
	0xe9, 0xa2, 0x18, 0x0e, 0x64, 0x80, 0x21, 0xf8, 0x44, 0x7c, 0x0f, 0x08, 0xd0, 0x50, 0xc3, 0xfc,
	0x32, 0xcb, 0xe4, 0x1e, 0xb6, 0x01, 0x76, 0xcc, 0x5e, 0x88, 0xfb, 0xd0, 0xaa, 0x74, 0x35, 0xdd,
	0xea, 0xc9, 0xf7, 0x6d, 0x2c, 0x3e, 0xb6, 0xb1, 0xf8, 0xdc, 0xc6, 0xe2, 0xf5, 0x2b, 0xae, 0xdd,
	0x34, 0xf8, 0x83, 0x8e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x1c, 0x67, 0x28, 0xb5, 0x01,
	0x00, 0x00,
}

func (m *SendChat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendChat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendChat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintChat(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Touid != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Touid))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RecvChat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecvChat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecvChat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintChat(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Fromuid != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Fromuid))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendChat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovChat(uint64(m.Type))
	}
	if m.Touid != 0 {
		n += 1 + sovChat(uint64(m.Touid))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RecvChat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovChat(uint64(m.Type))
	}
	if m.Fromuid != 0 {
		n += 1 + sovChat(uint64(m.Fromuid))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendChat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendChat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendChat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Touid", wireType)
			}
			m.Touid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Touid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RecvChat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecvChat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecvChat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fromuid", wireType)
			}
			m.Fromuid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fromuid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowChat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowChat
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChat = fmt.Errorf("proto: unexpected end of group")
)