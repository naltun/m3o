// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ip.proto

package ip

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

// Api Endpoints for Ip service

func NewIpEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Ip service

type IpService interface {
	Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error)
}

type ipService struct {
	c    client.Client
	name string
}

func NewIpService(name string, c client.Client) IpService {
	return &ipService{
		c:    c,
		name: name,
	}
}

func (c *ipService) Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error) {
	req := c.c.NewRequest(c.name, "Ip.Lookup", in)
	out := new(LookupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ip service

type IpHandler interface {
	Lookup(context.Context, *LookupRequest, *LookupResponse) error
}

func RegisterIpHandler(s server.Server, hdlr IpHandler, opts ...server.HandlerOption) error {
	type ip interface {
		Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error
	}
	type Ip struct {
		ip
	}
	h := &ipHandler{hdlr}
	return s.Handle(s.NewHandler(&Ip{h}, opts...))
}

type ipHandler struct {
	IpHandler
}

func (h *ipHandler) Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error {
	return h.IpHandler.Lookup(ctx, in, out)
}
