// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: entry/entry.proto

package entry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Entry service

type EntryService interface {
	CreateApp(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	ListApps(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	GetApp(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	CreateCluster(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	GetCluster(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	ListClusters(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	PullInstances(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	CreateNamespace(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
	UpdateItems(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error)
}

type entryService struct {
	c    client.Client
	name string
}

func NewEntryService(name string, c client.Client) EntryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "entry"
	}
	return &entryService{
		c:    c,
		name: name,
	}
}

func (c *entryService) CreateApp(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.CreateApp", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) ListApps(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.ListApps", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) GetApp(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.GetApp", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) CreateCluster(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.CreateCluster", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) GetCluster(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.GetCluster", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) ListClusters(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.ListClusters", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) PullInstances(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.PullInstances", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) CreateNamespace(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.CreateNamespace", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryService) UpdateItems(ctx context.Context, in *EntryRequest, opts ...client.CallOption) (*EntryResponse, error) {
	req := c.c.NewRequest(c.name, "Entry.UpdateItems", in)
	out := new(EntryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Entry service

type EntryHandler interface {
	CreateApp(context.Context, *EntryRequest, *EntryResponse) error
	ListApps(context.Context, *EntryRequest, *EntryResponse) error
	GetApp(context.Context, *EntryRequest, *EntryResponse) error
	CreateCluster(context.Context, *EntryRequest, *EntryResponse) error
	GetCluster(context.Context, *EntryRequest, *EntryResponse) error
	ListClusters(context.Context, *EntryRequest, *EntryResponse) error
	PullInstances(context.Context, *EntryRequest, *EntryResponse) error
	CreateNamespace(context.Context, *EntryRequest, *EntryResponse) error
	UpdateItems(context.Context, *EntryRequest, *EntryResponse) error
}

func RegisterEntryHandler(s server.Server, hdlr EntryHandler, opts ...server.HandlerOption) error {
	type entry interface {
		CreateApp(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		ListApps(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		GetApp(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		CreateCluster(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		GetCluster(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		ListClusters(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		PullInstances(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		CreateNamespace(ctx context.Context, in *EntryRequest, out *EntryResponse) error
		UpdateItems(ctx context.Context, in *EntryRequest, out *EntryResponse) error
	}
	type Entry struct {
		entry
	}
	h := &entryHandler{hdlr}
	return s.Handle(s.NewHandler(&Entry{h}, opts...))
}

type entryHandler struct {
	EntryHandler
}

func (h *entryHandler) CreateApp(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.CreateApp(ctx, in, out)
}

func (h *entryHandler) ListApps(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.ListApps(ctx, in, out)
}

func (h *entryHandler) GetApp(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.GetApp(ctx, in, out)
}

func (h *entryHandler) CreateCluster(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.CreateCluster(ctx, in, out)
}

func (h *entryHandler) GetCluster(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.GetCluster(ctx, in, out)
}

func (h *entryHandler) ListClusters(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.ListClusters(ctx, in, out)
}

func (h *entryHandler) PullInstances(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.PullInstances(ctx, in, out)
}

func (h *entryHandler) CreateNamespace(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.CreateNamespace(ctx, in, out)
}

func (h *entryHandler) UpdateItems(ctx context.Context, in *EntryRequest, out *EntryResponse) error {
	return h.EntryHandler.UpdateItems(ctx, in, out)
}
