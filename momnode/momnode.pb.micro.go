// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: momnode.proto

package momnode

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

// Api Endpoints for MomnodeService service

func NewMomnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MomnodeService service

type MomnodeService interface {
	// 获取动态的数量
	MomentsCount(ctx context.Context, in *MomentsCountReq, opts ...client.CallOption) (*MomentsCountRes, error)
	// 添加动态
	MomentsAdd(ctx context.Context, in *AddMomentsReq, opts ...client.CallOption) (*AddMomentsRes, error)
	// 获取动态列表
	MomentsList(ctx context.Context, in *MomentsListReq, opts ...client.CallOption) (*MomentsListRes, error)
}

type momnodeService struct {
	c    client.Client
	name string
}

func NewMomnodeService(name string, c client.Client) MomnodeService {
	return &momnodeService{
		c:    c,
		name: name,
	}
}

func (c *momnodeService) MomentsCount(ctx context.Context, in *MomentsCountReq, opts ...client.CallOption) (*MomentsCountRes, error) {
	req := c.c.NewRequest(c.name, "MomnodeService.MomentsCount", in)
	out := new(MomentsCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momnodeService) MomentsAdd(ctx context.Context, in *AddMomentsReq, opts ...client.CallOption) (*AddMomentsRes, error) {
	req := c.c.NewRequest(c.name, "MomnodeService.MomentsAdd", in)
	out := new(AddMomentsRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momnodeService) MomentsList(ctx context.Context, in *MomentsListReq, opts ...client.CallOption) (*MomentsListRes, error) {
	req := c.c.NewRequest(c.name, "MomnodeService.MomentsList", in)
	out := new(MomentsListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MomnodeService service

type MomnodeServiceHandler interface {
	// 获取动态的数量
	MomentsCount(context.Context, *MomentsCountReq, *MomentsCountRes) error
	// 添加动态
	MomentsAdd(context.Context, *AddMomentsReq, *AddMomentsRes) error
	// 获取动态列表
	MomentsList(context.Context, *MomentsListReq, *MomentsListRes) error
}

func RegisterMomnodeServiceHandler(s server.Server, hdlr MomnodeServiceHandler, opts ...server.HandlerOption) error {
	type momnodeService interface {
		MomentsCount(ctx context.Context, in *MomentsCountReq, out *MomentsCountRes) error
		MomentsAdd(ctx context.Context, in *AddMomentsReq, out *AddMomentsRes) error
		MomentsList(ctx context.Context, in *MomentsListReq, out *MomentsListRes) error
	}
	type MomnodeService struct {
		momnodeService
	}
	h := &momnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MomnodeService{h}, opts...))
}

type momnodeServiceHandler struct {
	MomnodeServiceHandler
}

func (h *momnodeServiceHandler) MomentsCount(ctx context.Context, in *MomentsCountReq, out *MomentsCountRes) error {
	return h.MomnodeServiceHandler.MomentsCount(ctx, in, out)
}

func (h *momnodeServiceHandler) MomentsAdd(ctx context.Context, in *AddMomentsReq, out *AddMomentsRes) error {
	return h.MomnodeServiceHandler.MomentsAdd(ctx, in, out)
}

func (h *momnodeServiceHandler) MomentsList(ctx context.Context, in *MomentsListReq, out *MomentsListRes) error {
	return h.MomnodeServiceHandler.MomentsList(ctx, in, out)
}
