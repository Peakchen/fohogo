// Code generated by protoc-gen-go.
// source: MSG_Chat.proto
// DO NOT EDIT!

/*
Package MSG_Chat is a generated protocol buffer package.

It is generated from these files:
	MSG_Chat.proto

It has these top-level messages:
	CS_BroadCast_Req
	SC_BroadCast_Rsp
*/
package MSG_Chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// add by stefan
// server
type SUBMSG int32

const (
	SUBMSG_Begin        SUBMSG = 0
	SUBMSG_CS_BroadCast SUBMSG = 1
	SUBMSG_SC_BroadCast SUBMSG = 2
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_BroadCast",
	2: "SC_BroadCast",
}
var SUBMSG_value = map[string]int32{
	"Begin":        0,
	"CS_BroadCast": 1,
	"SC_BroadCast": 2,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorCode int32

const (
	ErrorCode_Invalid ErrorCode = 0
	ErrorCode_Success ErrorCode = 1
	ErrorCode_Fail    ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var ErrorCode_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BroadCastRangeType int32

const (
	BroadCastRangeType_No   BroadCastRangeType = 0
	BroadCastRangeType_Zone BroadCastRangeType = 1
)

var BroadCastRangeType_name = map[int32]string{
	0: "No",
	1: "Zone",
}
var BroadCastRangeType_value = map[string]int32{
	"No":   0,
	"Zone": 1,
}

func (x BroadCastRangeType) String() string {
	return proto.EnumName(BroadCastRangeType_name, int32(x))
}
func (BroadCastRangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// CS_BroadCast
type CS_BroadCast_Req struct {
	SrcData string             `protobuf:"bytes,1,opt,name=SrcData" json:"SrcData,omitempty"`
	Range   BroadCastRangeType `protobuf:"varint,2,opt,name=range,enum=MSG_Chat.BroadCastRangeType" json:"range,omitempty"`
}

func (m *CS_BroadCast_Req) Reset()                    { *m = CS_BroadCast_Req{} }
func (m *CS_BroadCast_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_BroadCast_Req) ProtoMessage()               {}
func (*CS_BroadCast_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// SC_BroadCast
type SC_BroadCast_Rsp struct {
	SrcData string `protobuf:"bytes,1,opt,name=SrcData" json:"SrcData,omitempty"`
}

func (m *SC_BroadCast_Rsp) Reset()                    { *m = SC_BroadCast_Rsp{} }
func (m *SC_BroadCast_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_BroadCast_Rsp) ProtoMessage()               {}
func (*SC_BroadCast_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*CS_BroadCast_Req)(nil), "MSG_Chat.CS_BroadCast_Req")
	proto.RegisterType((*SC_BroadCast_Rsp)(nil), "MSG_Chat.SC_BroadCast_Rsp")
	proto.RegisterEnum("MSG_Chat.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_Chat.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("MSG_Chat.BroadCastRangeType", BroadCastRangeType_name, BroadCastRangeType_value)
}

func init() { proto.RegisterFile("MSG_Chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x9b, 0xe2, 0xba, 0xf5, 0x53, 0xc6, 0xe7, 0xe7, 0xa5, 0x07, 0x0f, 0x63, 0x07, 0x19,
	0x45, 0x36, 0x98, 0x07, 0xef, 0x8d, 0x3a, 0x3c, 0x4c, 0xa4, 0xd1, 0xcb, 0x2e, 0x35, 0xb6, 0xa1,
	0x16, 0x66, 0x53, 0x93, 0x2a, 0xf8, 0xdf, 0x4b, 0x94, 0x96, 0x82, 0xec, 0xf8, 0x92, 0xc7, 0xef,
	0xbd, 0xef, 0xc1, 0x74, 0x2b, 0x36, 0x19, 0x7f, 0x93, 0xed, 0xb2, 0x31, 0xba, 0xd5, 0x34, 0xe9,
	0xf4, 0xfc, 0x05, 0x90, 0x8b, 0x2c, 0x31, 0x5a, 0x16, 0x5c, 0xda, 0x36, 0x4b, 0xd5, 0x07, 0x45,
	0x30, 0x16, 0x26, 0xbf, 0x91, 0xad, 0x8c, 0xd8, 0x8c, 0x2d, 0xc2, 0xb4, 0x93, 0xb4, 0x86, 0x91,
	0x91, 0x75, 0xa9, 0x22, 0x7f, 0xc6, 0x16, 0xd3, 0xf5, 0xf9, 0xb2, 0xe7, 0xf6, 0x84, 0xd4, 0xfd,
	0x3f, 0x7d, 0x37, 0x2a, 0xfd, 0xb3, 0xce, 0x2f, 0x01, 0x05, 0x1f, 0x26, 0xd8, 0xe6, 0x70, 0x42,
	0x7c, 0x0d, 0x81, 0x78, 0x4e, 0xb6, 0x62, 0x43, 0x21, 0x8c, 0x12, 0x55, 0x56, 0x35, 0x7a, 0x84,
	0x70, 0x32, 0x2c, 0x89, 0xcc, 0xbd, 0x0c, 0xa1, 0xe8, 0xc7, 0x2b, 0x08, 0x6f, 0x8d, 0xd1, 0x86,
	0xeb, 0x42, 0xd1, 0x31, 0x8c, 0xef, 0xeb, 0x2f, 0xb9, 0xaf, 0x0a, 0xf4, 0x9c, 0x10, 0x9f, 0x79,
	0xae, 0xac, 0x45, 0x46, 0x13, 0x38, 0xba, 0x93, 0xd5, 0x1e, 0xfd, 0xf8, 0x02, 0xe8, 0x7f, 0x69,
	0x0a, 0xc0, 0x7f, 0xd0, 0xe8, 0x39, 0xdf, 0x4e, 0xd7, 0x0a, 0x59, 0x72, 0xb6, 0x3b, 0x7d, 0xb7,
	0xe5, 0xa3, 0xdb, 0x6d, 0xd5, 0x9d, 0xfb, 0x1a, 0xfc, 0xee, 0x78, 0xf5, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xbc, 0x34, 0x6a, 0x59, 0x01, 0x00, 0x00,
}