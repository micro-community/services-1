// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secrets/proto/secrets.proto

package go_micro_service_secrets

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

type CreateRequest struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *CreateRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReadRequest struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

type UpdateRequest struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *UpdateRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DeleteRequest struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type ReadResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa057a26a9bc300, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.secrets.CreateRequest")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.service.secrets.ReadRequest")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.service.secrets.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.service.secrets.DeleteRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.secrets.CreateResponse")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.service.secrets.ReadResponse")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.service.secrets.UpdateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.service.secrets.DeleteResponse")
}

func init() { proto.RegisterFile("secrets/proto/secrets.proto", fileDescriptor_8fa057a26a9bc300) }

var fileDescriptor_8fa057a26a9bc300 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x9b, 0x52, 0x82, 0x7a, 0xd0, 0x28, 0x3a, 0x31, 0x44, 0x65, 0x29, 0xe6, 0x5f, 0x26,
	0x57, 0x82, 0x89, 0x19, 0x3e, 0x41, 0x10, 0x62, 0xea, 0x10, 0xd2, 0x13, 0x54, 0x2a, 0x38, 0xd8,
	0x6e, 0xbf, 0x3e, 0xa8, 0x39, 0x07, 0x6a, 0xa4, 0x3a, 0x52, 0xb7, 0xbc, 0xc4, 0xef, 0x77, 0x2f,
	0xef, 0x0c, 0x67, 0x86, 0x2a, 0x4d, 0xd6, 0x4c, 0x6b, 0xad, 0xac, 0x9a, 0x3a, 0x25, 0x1b, 0x85,
	0xd9, 0x9b, 0x92, 0x1f, 0x8b, 0x4a, 0x2b, 0x69, 0x48, 0xaf, 0x17, 0x15, 0x49, 0xf7, 0x5d, 0xdc,
	0xc3, 0xe8, 0x41, 0x53, 0x69, 0xa9, 0xa0, 0xaf, 0x15, 0x19, 0x8b, 0x08, 0x83, 0xba, 0xb4, 0xef,
	0x59, 0x34, 0x39, 0xc8, 0x87, 0x45, 0xf3, 0x8c, 0xa7, 0x70, 0xb8, 0x2e, 0x97, 0x2b, 0xca, 0xfa,
	0x93, 0x28, 0x1f, 0x16, 0x2c, 0xc4, 0x39, 0x1c, 0x17, 0x54, 0xce, 0x03, 0xc6, 0x0d, 0xfd, 0xb9,
	0x9e, 0xef, 0x45, 0xbf, 0x80, 0xd1, 0x23, 0x2d, 0x29, 0x68, 0x15, 0x29, 0x24, 0x6d, 0x7a, 0x53,
	0xab, 0x4f, 0x43, 0xe2, 0x12, 0x4e, 0x38, 0x14, 0xeb, 0x3f, 0x78, 0xb4, 0x0d, 0x4f, 0x21, 0x69,
	0x73, 0x39, 0x5f, 0x0a, 0x49, 0x3b, 0x8e, 0xdf, 0xdc, 0x7e, 0xf7, 0xe1, 0xe8, 0x89, 0x5b, 0xc2,
	0x19, 0xc4, 0x3c, 0x07, 0x6f, 0xe4, 0xae, 0x2a, 0xa5, 0xd7, 0xe3, 0x38, 0xef, 0x3e, 0xe8, 0x46,
	0xf7, 0xf0, 0x05, 0x06, 0x9b, 0xd0, 0x78, 0xb5, 0xdb, 0xb3, 0xd5, 0xf4, 0xf8, 0xba, 0xeb, 0xd8,
	0x2f, 0x78, 0x06, 0x31, 0xff, 0x67, 0x28, 0xb7, 0xb7, 0xa1, 0x50, 0xee, 0x7f, 0x95, 0x35, 0x78,
	0x2e, 0x2d, 0x84, 0xf7, 0xb6, 0x18, 0xc2, 0xfb, 0xfd, 0x8b, 0xde, 0x6b, 0xdc, 0x5c, 0xde, 0xbb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x76, 0xb4, 0xcd, 0xdb, 0x02, 0x00, 0x00,
}
