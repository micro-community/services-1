// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/usage.proto

package usage

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

// Api Endpoints for Usage service

func NewUsageEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Usage service

type UsageService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Sweep(ctx context.Context, in *SweepRequest, opts ...client.CallOption) (*SweepResponse, error)
}

type usageService struct {
	c    client.Client
	name string
}

func NewUsageService(name string, c client.Client) UsageService {
	return &usageService{
		c:    c,
		name: name,
	}
}

func (c *usageService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Usage.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageService) Sweep(ctx context.Context, in *SweepRequest, opts ...client.CallOption) (*SweepResponse, error) {
	req := c.c.NewRequest(c.name, "Usage.Sweep", in)
	out := new(SweepResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Usage service

type UsageHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Sweep(context.Context, *SweepRequest, *SweepResponse) error
}

func RegisterUsageHandler(s server.Server, hdlr UsageHandler, opts ...server.HandlerOption) error {
	type usage interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Sweep(ctx context.Context, in *SweepRequest, out *SweepResponse) error
	}
	type Usage struct {
		usage
	}
	h := &usageHandler{hdlr}
	return s.Handle(s.NewHandler(&Usage{h}, opts...))
}

type usageHandler struct {
	UsageHandler
}

func (h *usageHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.UsageHandler.Read(ctx, in, out)
}

func (h *usageHandler) Sweep(ctx context.Context, in *SweepRequest, out *SweepResponse) error {
	return h.UsageHandler.Sweep(ctx, in, out)
}
