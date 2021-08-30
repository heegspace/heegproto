// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: paynode.proto

package paynode

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

// Api Endpoints for SponsornodeService service

func NewSponsornodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SponsornodeService service

type SponsornodeService interface {
	// 添加赞助信息
	SponsorAdd(ctx context.Context, in *SponsorAddReq, opts ...client.CallOption) (*SponsorAddRes, error)
	// 获取赞助列表
	SponsorList(ctx context.Context, in *SponsorListReq, opts ...client.CallOption) (*SponsorListRes, error)
}

type sponsornodeService struct {
	c    client.Client
	name string
}

func NewSponsornodeService(name string, c client.Client) SponsornodeService {
	return &sponsornodeService{
		c:    c,
		name: name,
	}
}

func (c *sponsornodeService) SponsorAdd(ctx context.Context, in *SponsorAddReq, opts ...client.CallOption) (*SponsorAddRes, error) {
	req := c.c.NewRequest(c.name, "SponsornodeService.SponsorAdd", in)
	out := new(SponsorAddRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sponsornodeService) SponsorList(ctx context.Context, in *SponsorListReq, opts ...client.CallOption) (*SponsorListRes, error) {
	req := c.c.NewRequest(c.name, "SponsornodeService.SponsorList", in)
	out := new(SponsorListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SponsornodeService service

type SponsornodeServiceHandler interface {
	// 添加赞助信息
	SponsorAdd(context.Context, *SponsorAddReq, *SponsorAddRes) error
	// 获取赞助列表
	SponsorList(context.Context, *SponsorListReq, *SponsorListRes) error
}

func RegisterSponsornodeServiceHandler(s server.Server, hdlr SponsornodeServiceHandler, opts ...server.HandlerOption) error {
	type sponsornodeService interface {
		SponsorAdd(ctx context.Context, in *SponsorAddReq, out *SponsorAddRes) error
		SponsorList(ctx context.Context, in *SponsorListReq, out *SponsorListRes) error
	}
	type SponsornodeService struct {
		sponsornodeService
	}
	h := &sponsornodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SponsornodeService{h}, opts...))
}

type sponsornodeServiceHandler struct {
	SponsornodeServiceHandler
}

func (h *sponsornodeServiceHandler) SponsorAdd(ctx context.Context, in *SponsorAddReq, out *SponsorAddRes) error {
	return h.SponsornodeServiceHandler.SponsorAdd(ctx, in, out)
}

func (h *sponsornodeServiceHandler) SponsorList(ctx context.Context, in *SponsorListReq, out *SponsorListRes) error {
	return h.SponsornodeServiceHandler.SponsorList(ctx, in, out)
}
