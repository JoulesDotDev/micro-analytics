// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/analytics.proto

package analytics

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

// Api Endpoints for Analytics service

func NewAnalyticsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Analytics service

type AnalyticsService interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error)
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...client.CallOption) (*DeleteProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*UpdateProjectResponse, error)
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...client.CallOption) (*CreateActionResponse, error)
	ListActions(ctx context.Context, in *ListActionsRequest, opts ...client.CallOption) (*ListActionsResponse, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...client.CallOption) (*DeleteActionResponse, error)
	UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...client.CallOption) (*UpdateActionResponse, error)
}

type analyticsService struct {
	c    client.Client
	name string
}

func NewAnalyticsService(name string, c client.Client) AnalyticsService {
	return &analyticsService{
		c:    c,
		name: name,
	}
}

func (c *analyticsService) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.CreateProject", in)
	out := new(CreateProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.ListProjects", in)
	out := new(ListProjectsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...client.CallOption) (*DeleteProjectResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.DeleteProject", in)
	out := new(DeleteProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*UpdateProjectResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.UpdateProject", in)
	out := new(UpdateProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...client.CallOption) (*CreateActionResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.CreateAction", in)
	out := new(CreateActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) ListActions(ctx context.Context, in *ListActionsRequest, opts ...client.CallOption) (*ListActionsResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.ListActions", in)
	out := new(ListActionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...client.CallOption) (*DeleteActionResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.DeleteAction", in)
	out := new(DeleteActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsService) UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...client.CallOption) (*UpdateActionResponse, error) {
	req := c.c.NewRequest(c.name, "Analytics.UpdateAction", in)
	out := new(UpdateActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analytics service

type AnalyticsHandler interface {
	CreateProject(context.Context, *CreateProjectRequest, *CreateProjectResponse) error
	ListProjects(context.Context, *ListProjectsRequest, *ListProjectsResponse) error
	DeleteProject(context.Context, *DeleteProjectRequest, *DeleteProjectResponse) error
	UpdateProject(context.Context, *UpdateProjectRequest, *UpdateProjectResponse) error
	CreateAction(context.Context, *CreateActionRequest, *CreateActionResponse) error
	ListActions(context.Context, *ListActionsRequest, *ListActionsResponse) error
	DeleteAction(context.Context, *DeleteActionRequest, *DeleteActionResponse) error
	UpdateAction(context.Context, *UpdateActionRequest, *UpdateActionResponse) error
}

func RegisterAnalyticsHandler(s server.Server, hdlr AnalyticsHandler, opts ...server.HandlerOption) error {
	type analytics interface {
		CreateProject(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error
		ListProjects(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error
		DeleteProject(ctx context.Context, in *DeleteProjectRequest, out *DeleteProjectResponse) error
		UpdateProject(ctx context.Context, in *UpdateProjectRequest, out *UpdateProjectResponse) error
		CreateAction(ctx context.Context, in *CreateActionRequest, out *CreateActionResponse) error
		ListActions(ctx context.Context, in *ListActionsRequest, out *ListActionsResponse) error
		DeleteAction(ctx context.Context, in *DeleteActionRequest, out *DeleteActionResponse) error
		UpdateAction(ctx context.Context, in *UpdateActionRequest, out *UpdateActionResponse) error
	}
	type Analytics struct {
		analytics
	}
	h := &analyticsHandler{hdlr}
	return s.Handle(s.NewHandler(&Analytics{h}, opts...))
}

type analyticsHandler struct {
	AnalyticsHandler
}

func (h *analyticsHandler) CreateProject(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error {
	return h.AnalyticsHandler.CreateProject(ctx, in, out)
}

func (h *analyticsHandler) ListProjects(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error {
	return h.AnalyticsHandler.ListProjects(ctx, in, out)
}

func (h *analyticsHandler) DeleteProject(ctx context.Context, in *DeleteProjectRequest, out *DeleteProjectResponse) error {
	return h.AnalyticsHandler.DeleteProject(ctx, in, out)
}

func (h *analyticsHandler) UpdateProject(ctx context.Context, in *UpdateProjectRequest, out *UpdateProjectResponse) error {
	return h.AnalyticsHandler.UpdateProject(ctx, in, out)
}

func (h *analyticsHandler) CreateAction(ctx context.Context, in *CreateActionRequest, out *CreateActionResponse) error {
	return h.AnalyticsHandler.CreateAction(ctx, in, out)
}

func (h *analyticsHandler) ListActions(ctx context.Context, in *ListActionsRequest, out *ListActionsResponse) error {
	return h.AnalyticsHandler.ListActions(ctx, in, out)
}

func (h *analyticsHandler) DeleteAction(ctx context.Context, in *DeleteActionRequest, out *DeleteActionResponse) error {
	return h.AnalyticsHandler.DeleteAction(ctx, in, out)
}

func (h *analyticsHandler) UpdateAction(ctx context.Context, in *UpdateActionRequest, out *UpdateActionResponse) error {
	return h.AnalyticsHandler.UpdateAction(ctx, in, out)
}
