// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: macipnode.proto

package macipnode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/heegspace/heegproto/common"
	_ "github.com/heegspace/heegproto/rescode"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for MacipnodeService service

func NewMacipnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MacipnodeService service

type MacipnodeService interface {
	// ip地址转换
	IpToAddress(ctx context.Context, in *IpToAddressReq, opts ...client.CallOption) (*IpToAddressRes, error)
}

type macipnodeService struct {
	c    client.Client
	name string
}

func NewMacipnodeService(name string, c client.Client) MacipnodeService {
	return &macipnodeService{
		c:    c,
		name: name,
	}
}

func (c *macipnodeService) IpToAddress(ctx context.Context, in *IpToAddressReq, opts ...client.CallOption) (*IpToAddressRes, error) {
	req := c.c.NewRequest(c.name, "MacipnodeService.IpToAddress", in)
	out := new(IpToAddressRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MacipnodeService service

type MacipnodeServiceHandler interface {
	// ip地址转换
	IpToAddress(context.Context, *IpToAddressReq, *IpToAddressRes) error
}

func RegisterMacipnodeServiceHandler(s server.Server, hdlr MacipnodeServiceHandler, opts ...server.HandlerOption) error {
	type macipnodeService interface {
		IpToAddress(ctx context.Context, in *IpToAddressReq, out *IpToAddressRes) error
	}
	type MacipnodeService struct {
		macipnodeService
	}
	h := &macipnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MacipnodeService{h}, opts...))
}

type macipnodeServiceHandler struct {
	MacipnodeServiceHandler
}

func (h *macipnodeServiceHandler) IpToAddress(ctx context.Context, in *IpToAddressReq, out *IpToAddressRes) error {
	return h.MacipnodeServiceHandler.IpToAddress(ctx, in, out)
}
