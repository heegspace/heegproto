// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: strategynode.proto

package strategynode

import (
	_ "github.com/heegspace/heegproto/common"
	_ "github.com/heegspace/heegproto/rescode"
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for StrategynodeService service

func NewStrategynodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StrategynodeService service

type StrategynodeService interface {
	// 更新修改试题的奖励积分
	QueryStrategy(ctx context.Context, in *QueryStrategyReq, opts ...client.CallOption) (*QueryStrategyRes, error)
}

type strategynodeService struct {
	c    client.Client
	name string
}

func NewStrategynodeService(name string, c client.Client) StrategynodeService {
	return &strategynodeService{
		c:    c,
		name: name,
	}
}

func (c *strategynodeService) QueryStrategy(ctx context.Context, in *QueryStrategyReq, opts ...client.CallOption) (*QueryStrategyRes, error) {
	req := c.c.NewRequest(c.name, "StrategynodeService.QueryStrategy", in)
	out := new(QueryStrategyRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StrategynodeService service

type StrategynodeServiceHandler interface {
	// 更新修改试题的奖励积分
	QueryStrategy(context.Context, *QueryStrategyReq, *QueryStrategyRes) error
}

func RegisterStrategynodeServiceHandler(s server.Server, hdlr StrategynodeServiceHandler, opts ...server.HandlerOption) error {
	type strategynodeService interface {
		QueryStrategy(ctx context.Context, in *QueryStrategyReq, out *QueryStrategyRes) error
	}
	type StrategynodeService struct {
		strategynodeService
	}
	h := &strategynodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StrategynodeService{h}, opts...))
}

type strategynodeServiceHandler struct {
	StrategynodeServiceHandler
}

func (h *strategynodeServiceHandler) QueryStrategy(ctx context.Context, in *QueryStrategyReq, out *QueryStrategyRes) error {
	return h.StrategynodeServiceHandler.QueryStrategy(ctx, in, out)
}
