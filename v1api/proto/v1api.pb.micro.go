// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/v1api.proto

package v1api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for V1 service

func NewV1Endpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for V1 service

type V1Service interface {
	GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...client.CallOption) (*GenerateKeyResponse, error)
	ListKeys(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	RevokeKey(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error)
	UnblockKey(ctx context.Context, in *UnblockKeyRequest, opts ...client.CallOption) (*UnblockKeyResponse, error)
	BlockKey(ctx context.Context, in *BlockKeyRequest, opts ...client.CallOption) (*BlockKeyResponse, error)
}

type v1Service struct {
	c    client.Client
	name string
}

func NewV1Service(name string, c client.Client) V1Service {
	return &v1Service{
		c:    c,
		name: name,
	}
}

func (c *v1Service) GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...client.CallOption) (*GenerateKeyResponse, error) {
	req := c.c.NewRequest(c.name, "V1.GenerateKey", in)
	out := new(GenerateKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Service) ListKeys(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "V1.ListKeys", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Service) RevokeKey(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error) {
	req := c.c.NewRequest(c.name, "V1.RevokeKey", in)
	out := new(RevokeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Service) UnblockKey(ctx context.Context, in *UnblockKeyRequest, opts ...client.CallOption) (*UnblockKeyResponse, error) {
	req := c.c.NewRequest(c.name, "V1.UnblockKey", in)
	out := new(UnblockKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Service) BlockKey(ctx context.Context, in *BlockKeyRequest, opts ...client.CallOption) (*BlockKeyResponse, error) {
	req := c.c.NewRequest(c.name, "V1.BlockKey", in)
	out := new(BlockKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for V1 service

type V1Handler interface {
	GenerateKey(context.Context, *GenerateKeyRequest, *GenerateKeyResponse) error
	ListKeys(context.Context, *ListRequest, *ListResponse) error
	RevokeKey(context.Context, *RevokeRequest, *RevokeResponse) error
	UnblockKey(context.Context, *UnblockKeyRequest, *UnblockKeyResponse) error
	BlockKey(context.Context, *BlockKeyRequest, *BlockKeyResponse) error
}

func RegisterV1Handler(s server.Server, hdlr V1Handler, opts ...server.HandlerOption) error {
	type v1 interface {
		GenerateKey(ctx context.Context, in *GenerateKeyRequest, out *GenerateKeyResponse) error
		ListKeys(ctx context.Context, in *ListRequest, out *ListResponse) error
		RevokeKey(ctx context.Context, in *RevokeRequest, out *RevokeResponse) error
		UnblockKey(ctx context.Context, in *UnblockKeyRequest, out *UnblockKeyResponse) error
		BlockKey(ctx context.Context, in *BlockKeyRequest, out *BlockKeyResponse) error
	}
	type V1 struct {
		v1
	}
	h := &v1Handler{hdlr}
	return s.Handle(s.NewHandler(&V1{h}, opts...))
}

type v1Handler struct {
	V1Handler
}

func (h *v1Handler) GenerateKey(ctx context.Context, in *GenerateKeyRequest, out *GenerateKeyResponse) error {
	return h.V1Handler.GenerateKey(ctx, in, out)
}

func (h *v1Handler) ListKeys(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.V1Handler.ListKeys(ctx, in, out)
}

func (h *v1Handler) RevokeKey(ctx context.Context, in *RevokeRequest, out *RevokeResponse) error {
	return h.V1Handler.RevokeKey(ctx, in, out)
}

func (h *v1Handler) UnblockKey(ctx context.Context, in *UnblockKeyRequest, out *UnblockKeyResponse) error {
	return h.V1Handler.UnblockKey(ctx, in, out)
}

func (h *v1Handler) BlockKey(ctx context.Context, in *BlockKeyRequest, out *BlockKeyResponse) error {
	return h.V1Handler.BlockKey(ctx, in, out)
}
