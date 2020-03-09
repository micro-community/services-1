// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sprints/sprints.proto

package go_micro_srv_distributed_sprints

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Sprints service

type SprintsService interface {
	CreateSprint(ctx context.Context, in *CreateSprintRequest, opts ...client.CallOption) (*CreateSprintResponse, error)
	ListSprints(ctx context.Context, in *ListSprintsRequest, opts ...client.CallOption) (*ListSprintsResponse, error)
	ReadSprint(ctx context.Context, in *ReadSprintRequest, opts ...client.CallOption) (*ReadSprintResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...client.CallOption) (*DeleteTaskResponse, error)
	CreateObjective(ctx context.Context, in *CreateObjectiveRequest, opts ...client.CallOption) (*CreateObjectiveResponse, error)
	UpdateObjective(ctx context.Context, in *UpdateObjectiveRequest, opts ...client.CallOption) (*UpdateObjectiveResponse, error)
	DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, opts ...client.CallOption) (*DeleteObjectiveResponse, error)
}

type sprintsService struct {
	c    client.Client
	name string
}

func NewSprintsService(name string, c client.Client) SprintsService {
	return &sprintsService{
		c:    c,
		name: name,
	}
}

func (c *sprintsService) CreateSprint(ctx context.Context, in *CreateSprintRequest, opts ...client.CallOption) (*CreateSprintResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.CreateSprint", in)
	out := new(CreateSprintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) ListSprints(ctx context.Context, in *ListSprintsRequest, opts ...client.CallOption) (*ListSprintsResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.ListSprints", in)
	out := new(ListSprintsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) ReadSprint(ctx context.Context, in *ReadSprintRequest, opts ...client.CallOption) (*ReadSprintResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.ReadSprint", in)
	out := new(ReadSprintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.CreateTask", in)
	out := new(CreateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.UpdateTask", in)
	out := new(UpdateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...client.CallOption) (*DeleteTaskResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.DeleteTask", in)
	out := new(DeleteTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) CreateObjective(ctx context.Context, in *CreateObjectiveRequest, opts ...client.CallOption) (*CreateObjectiveResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.CreateObjective", in)
	out := new(CreateObjectiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) UpdateObjective(ctx context.Context, in *UpdateObjectiveRequest, opts ...client.CallOption) (*UpdateObjectiveResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.UpdateObjective", in)
	out := new(UpdateObjectiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintsService) DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, opts ...client.CallOption) (*DeleteObjectiveResponse, error) {
	req := c.c.NewRequest(c.name, "Sprints.DeleteObjective", in)
	out := new(DeleteObjectiveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sprints service

type SprintsHandler interface {
	CreateSprint(context.Context, *CreateSprintRequest, *CreateSprintResponse) error
	ListSprints(context.Context, *ListSprintsRequest, *ListSprintsResponse) error
	ReadSprint(context.Context, *ReadSprintRequest, *ReadSprintResponse) error
	CreateTask(context.Context, *CreateTaskRequest, *CreateTaskResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *UpdateTaskResponse) error
	DeleteTask(context.Context, *DeleteTaskRequest, *DeleteTaskResponse) error
	CreateObjective(context.Context, *CreateObjectiveRequest, *CreateObjectiveResponse) error
	UpdateObjective(context.Context, *UpdateObjectiveRequest, *UpdateObjectiveResponse) error
	DeleteObjective(context.Context, *DeleteObjectiveRequest, *DeleteObjectiveResponse) error
}

func RegisterSprintsHandler(s server.Server, hdlr SprintsHandler, opts ...server.HandlerOption) error {
	type sprints interface {
		CreateSprint(ctx context.Context, in *CreateSprintRequest, out *CreateSprintResponse) error
		ListSprints(ctx context.Context, in *ListSprintsRequest, out *ListSprintsResponse) error
		ReadSprint(ctx context.Context, in *ReadSprintRequest, out *ReadSprintResponse) error
		CreateTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error
		DeleteTask(ctx context.Context, in *DeleteTaskRequest, out *DeleteTaskResponse) error
		CreateObjective(ctx context.Context, in *CreateObjectiveRequest, out *CreateObjectiveResponse) error
		UpdateObjective(ctx context.Context, in *UpdateObjectiveRequest, out *UpdateObjectiveResponse) error
		DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, out *DeleteObjectiveResponse) error
	}
	type Sprints struct {
		sprints
	}
	h := &sprintsHandler{hdlr}
	return s.Handle(s.NewHandler(&Sprints{h}, opts...))
}

type sprintsHandler struct {
	SprintsHandler
}

func (h *sprintsHandler) CreateSprint(ctx context.Context, in *CreateSprintRequest, out *CreateSprintResponse) error {
	return h.SprintsHandler.CreateSprint(ctx, in, out)
}

func (h *sprintsHandler) ListSprints(ctx context.Context, in *ListSprintsRequest, out *ListSprintsResponse) error {
	return h.SprintsHandler.ListSprints(ctx, in, out)
}

func (h *sprintsHandler) ReadSprint(ctx context.Context, in *ReadSprintRequest, out *ReadSprintResponse) error {
	return h.SprintsHandler.ReadSprint(ctx, in, out)
}

func (h *sprintsHandler) CreateTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error {
	return h.SprintsHandler.CreateTask(ctx, in, out)
}

func (h *sprintsHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error {
	return h.SprintsHandler.UpdateTask(ctx, in, out)
}

func (h *sprintsHandler) DeleteTask(ctx context.Context, in *DeleteTaskRequest, out *DeleteTaskResponse) error {
	return h.SprintsHandler.DeleteTask(ctx, in, out)
}

func (h *sprintsHandler) CreateObjective(ctx context.Context, in *CreateObjectiveRequest, out *CreateObjectiveResponse) error {
	return h.SprintsHandler.CreateObjective(ctx, in, out)
}

func (h *sprintsHandler) UpdateObjective(ctx context.Context, in *UpdateObjectiveRequest, out *UpdateObjectiveResponse) error {
	return h.SprintsHandler.UpdateObjective(ctx, in, out)
}

func (h *sprintsHandler) DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, out *DeleteObjectiveResponse) error {
	return h.SprintsHandler.DeleteObjective(ctx, in, out)
}
