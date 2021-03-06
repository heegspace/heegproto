// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: searchnode.proto

package searchnode

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

// Api Endpoints for SearchnodeService service

func NewSearchnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SearchnodeService service

type SearchnodeService interface {
	// 搜索试题
	SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...client.CallOption) (*SearchQuestionRes, error)
	// 搜索关键字补全
	SearchItem(ctx context.Context, in *SearchItemReq, opts ...client.CallOption) (*SearchItemRes, error)
	// 获取搜索记录
	SearchHistory(ctx context.Context, in *SearchHistoryReq, opts ...client.CallOption) (*SearchHistoryRes, error)
}

type searchnodeService struct {
	c    client.Client
	name string
}

func NewSearchnodeService(name string, c client.Client) SearchnodeService {
	return &searchnodeService{
		c:    c,
		name: name,
	}
}

func (c *searchnodeService) SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...client.CallOption) (*SearchQuestionRes, error) {
	req := c.c.NewRequest(c.name, "SearchnodeService.SearchQuestion", in)
	out := new(SearchQuestionRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchnodeService) SearchItem(ctx context.Context, in *SearchItemReq, opts ...client.CallOption) (*SearchItemRes, error) {
	req := c.c.NewRequest(c.name, "SearchnodeService.SearchItem", in)
	out := new(SearchItemRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchnodeService) SearchHistory(ctx context.Context, in *SearchHistoryReq, opts ...client.CallOption) (*SearchHistoryRes, error) {
	req := c.c.NewRequest(c.name, "SearchnodeService.SearchHistory", in)
	out := new(SearchHistoryRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchnodeService service

type SearchnodeServiceHandler interface {
	// 搜索试题
	SearchQuestion(context.Context, *SearchQuestionReq, *SearchQuestionRes) error
	// 搜索关键字补全
	SearchItem(context.Context, *SearchItemReq, *SearchItemRes) error
	// 获取搜索记录
	SearchHistory(context.Context, *SearchHistoryReq, *SearchHistoryRes) error
}

func RegisterSearchnodeServiceHandler(s server.Server, hdlr SearchnodeServiceHandler, opts ...server.HandlerOption) error {
	type searchnodeService interface {
		SearchQuestion(ctx context.Context, in *SearchQuestionReq, out *SearchQuestionRes) error
		SearchItem(ctx context.Context, in *SearchItemReq, out *SearchItemRes) error
		SearchHistory(ctx context.Context, in *SearchHistoryReq, out *SearchHistoryRes) error
	}
	type SearchnodeService struct {
		searchnodeService
	}
	h := &searchnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SearchnodeService{h}, opts...))
}

type searchnodeServiceHandler struct {
	SearchnodeServiceHandler
}

func (h *searchnodeServiceHandler) SearchQuestion(ctx context.Context, in *SearchQuestionReq, out *SearchQuestionRes) error {
	return h.SearchnodeServiceHandler.SearchQuestion(ctx, in, out)
}

func (h *searchnodeServiceHandler) SearchItem(ctx context.Context, in *SearchItemReq, out *SearchItemRes) error {
	return h.SearchnodeServiceHandler.SearchItem(ctx, in, out)
}

func (h *searchnodeServiceHandler) SearchHistory(ctx context.Context, in *SearchHistoryReq, out *SearchHistoryRes) error {
	return h.SearchnodeServiceHandler.SearchHistory(ctx, in, out)
}
