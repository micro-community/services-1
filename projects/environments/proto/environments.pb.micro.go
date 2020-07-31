// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/environments.proto

package go_micro_service_projects_environments

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
	microClient "github.com/micro/micro/v3/service/client"
	microServer "github.com/micro/micro/v3/service/server"
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
var _ = microServer.Handle
var _ = microClient.Call

// Api Endpoints for Environments service

func NewEnvironmentsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Environments service

type EnvironmentsService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type environmentsService struct {
	name string
}

func NewEnvironmentsService(name string) EnvironmentsService {
	return &environmentsService{name: name}
}

func (c *environmentsService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := microClient.NewRequest(c.name, "Environments.Create", in)
	out := new(CreateResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := microClient.NewRequest(c.name, "Environments.Read", in)
	out := new(ReadResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := microClient.NewRequest(c.name, "Environments.Update", in)
	out := new(UpdateResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := microClient.NewRequest(c.name, "Environments.Delete", in)
	out := new(DeleteResponse)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Environments service

type EnvironmentsHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterEnvironmentsHandler(hdlr EnvironmentsHandler, opts ...server.HandlerOption) error {
	type environments interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type Environments struct {
		environments
	}
	h := &environmentsHandler{hdlr}
	return microServer.Handle(microServer.NewHandler(&Environments{h}, opts...))
}

type environmentsHandler struct {
	EnvironmentsHandler
}

func (h *environmentsHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.EnvironmentsHandler.Create(ctx, in, out)
}

func (h *environmentsHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.EnvironmentsHandler.Read(ctx, in, out)
}

func (h *environmentsHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.EnvironmentsHandler.Update(ctx, in, out)
}

func (h *environmentsHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.EnvironmentsHandler.Delete(ctx, in, out)
}
