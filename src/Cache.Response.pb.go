// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Cache.Response.proto

package Cache

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

type Response struct {
	LinkedinID           string   `protobuf:"bytes,1,opt,name=LinkedinID,proto3" json:"LinkedinID,omitempty"`
	BingBidder           string   `protobuf:"bytes,2,opt,name=BingBidder,proto3" json:"BingBidder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe2c8224a41c59ec, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetLinkedinID() string {
	if m != nil {
		return m.LinkedinID
	}
	return ""
}

func (m *Response) GetBingBidder() string {
	if m != nil {
		return m.BingBidder
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "Cache.Response")
}

func init() { proto.RegisterFile("Cache.Response.proto", fileDescriptor_fe2c8224a41c59ec) }

var fileDescriptor_fe2c8224a41c59ec = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x0b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x8b, 0x2a, 0x79, 0x71, 0x71, 0xc0, 0x24, 0x84, 0xe4, 0xb8, 0xb8, 0x7c, 0x32,
	0xf3, 0xb2, 0x53, 0x53, 0x32, 0xf3, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90,
	0x44, 0x40, 0xf2, 0x4e, 0x99, 0x79, 0xe9, 0x4e, 0x99, 0x29, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x10,
	0x79, 0x84, 0x48, 0x12, 0x1b, 0xd8, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xed,
	0x0f, 0x91, 0x71, 0x00, 0x00, 0x00,
}
