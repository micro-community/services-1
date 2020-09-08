// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/client/client.proto

package go_micro_api_client

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

type Request struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab285fece1c29ddf, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab285fece1c29ddf, []int{1}
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

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab285fece1c29ddf, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.api.client.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.client.Response")
	proto.RegisterType((*Message)(nil), "go.micro.api.client.Message")
}

func init() { proto.RegisterFile("proto/client/client.proto", fileDescriptor_ab285fece1c29ddf) }

var fileDescriptor_ab285fece1c29ddf = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x1b, 0x92, 0x72, 0x74, 0x3a, 0x18, 0x42, 0x05, 0xa8, 0x64, 0xca, 0x64, 0x10,
	0xfc, 0x84, 0x8a, 0x09, 0xb1, 0x04, 0xc4, 0x8a, 0x52, 0xf7, 0x54, 0x59, 0x4a, 0x7d, 0x26, 0x3e,
	0x2a, 0xe5, 0xbf, 0xf0, 0x63, 0x91, 0x9c, 0x40, 0x91, 0xa0, 0x13, 0x93, 0xfd, 0xee, 0x93, 0xee,
	0x3d, 0x3f, 0xc3, 0x99, 0x6f, 0x59, 0xf8, 0xda, 0x34, 0x96, 0x9c, 0x0c, 0x87, 0x8e, 0x33, 0x3c,
	0x59, 0xb3, 0xde, 0x58, 0xd3, 0xb2, 0xae, 0xbd, 0xd5, 0x3d, 0x2a, 0xb6, 0x90, 0x55, 0xf4, 0xf6,
	0x4e, 0x41, 0x30, 0x87, 0x2c, 0x50, 0xbb, 0xb5, 0x86, 0x72, 0x35, 0x57, 0xe5, 0x51, 0xf5, 0x25,
	0x71, 0x06, 0x13, 0x72, 0x2b, 0xcf, 0xd6, 0x49, 0x3e, 0x8a, 0xe8, 0x5b, 0xe3, 0x15, 0x4c, 0x0d,
	0x3b, 0x21, 0x27, 0xaf, 0xd2, 0x79, 0xca, 0xc7, 0x91, 0x1f, 0x0f, 0xb3, 0xe7, 0xce, 0x13, 0x22,
	0x24, 0x4b, 0x5e, 0x75, 0x79, 0x32, 0x57, 0xe5, 0xb4, 0x8a, 0xf7, 0xe2, 0x12, 0x26, 0x15, 0x05,
	0xcf, 0x2e, 0xec, 0xb8, 0xfa, 0xc1, 0x5f, 0x20, 0x7b, 0xa4, 0x10, 0xea, 0x35, 0xe1, 0x29, 0x1c,
	0x0a, 0x7b, 0x6b, 0x86, 0x54, 0xbd, 0xf8, 0xe5, 0x3b, 0xda, 0xef, 0x3b, 0xde, 0xed, 0xbd, 0xfd,
	0x50, 0x90, 0x2e, 0xe2, 0xd3, 0xf1, 0x1e, 0x92, 0x45, 0xdd, 0x34, 0x78, 0xae, 0xff, 0x28, 0x46,
	0x0f, 0xad, 0xcc, 0x2e, 0xf6, 0xd0, 0x3e, 0x7b, 0x71, 0x80, 0x0f, 0x90, 0x3e, 0x49, 0x4b, 0xf5,
	0xe6, 0x9f, 0x8b, 0x4a, 0x75, 0xa3, 0x96, 0x69, 0xfc, 0xaa, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0x47, 0x2c, 0xc5, 0xc7, 0x01, 0x00, 0x00,
}
