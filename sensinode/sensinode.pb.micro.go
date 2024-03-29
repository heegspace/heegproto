// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sensinode.proto

package sensinode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/heegspace/heegproto/common"
	_ "github.com/heegspace/heegproto/rescode"
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

// Api Endpoints for SensinodeService service

func NewSensinodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SensinodeService service

type SensinodeService interface {
	// 更新修改试题的奖励积分
	RefreshModifyReward(ctx context.Context, in *RefreshModifyRewardReq, opts ...client.CallOption) (*RefreshModifyRewardRes, error)
	// 更新添加试题的奖励积分
	RefreshAddReward(ctx context.Context, in *RefreshAddRewardReq, opts ...client.CallOption) (*RefreshAddRewardRes, error)
	// 更新用户的coin数值
	RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, opts ...client.CallOption) (*RefreshUserCoinRes, error)
	// 刷新识别奖励
	RefreshIdentReward(ctx context.Context, in *RefreshIdentRewardReq, opts ...client.CallOption) (*RefreshIdentRewardRes, error)
	// 更新用户vip
	RefreshUserVip(ctx context.Context, in *RefreshUserVipReq, opts ...client.CallOption) (*RefreshUserVipRes, error)
	// 更新用户积分
	RefreshUserScore(ctx context.Context, in *RefreshUserScoreReq, opts ...client.CallOption) (*RefreshUserScoreRes, error)
}

type sensinodeService struct {
	c    client.Client
	name string
}

func NewSensinodeService(name string, c client.Client) SensinodeService {
	return &sensinodeService{
		c:    c,
		name: name,
	}
}

func (c *sensinodeService) RefreshModifyReward(ctx context.Context, in *RefreshModifyRewardReq, opts ...client.CallOption) (*RefreshModifyRewardRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshModifyReward", in)
	out := new(RefreshModifyRewardRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensinodeService) RefreshAddReward(ctx context.Context, in *RefreshAddRewardReq, opts ...client.CallOption) (*RefreshAddRewardRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshAddReward", in)
	out := new(RefreshAddRewardRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensinodeService) RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, opts ...client.CallOption) (*RefreshUserCoinRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshUserCoin", in)
	out := new(RefreshUserCoinRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensinodeService) RefreshIdentReward(ctx context.Context, in *RefreshIdentRewardReq, opts ...client.CallOption) (*RefreshIdentRewardRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshIdentReward", in)
	out := new(RefreshIdentRewardRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensinodeService) RefreshUserVip(ctx context.Context, in *RefreshUserVipReq, opts ...client.CallOption) (*RefreshUserVipRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshUserVip", in)
	out := new(RefreshUserVipRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensinodeService) RefreshUserScore(ctx context.Context, in *RefreshUserScoreReq, opts ...client.CallOption) (*RefreshUserScoreRes, error) {
	req := c.c.NewRequest(c.name, "SensinodeService.RefreshUserScore", in)
	out := new(RefreshUserScoreRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SensinodeService service

type SensinodeServiceHandler interface {
	// 更新修改试题的奖励积分
	RefreshModifyReward(context.Context, *RefreshModifyRewardReq, *RefreshModifyRewardRes) error
	// 更新添加试题的奖励积分
	RefreshAddReward(context.Context, *RefreshAddRewardReq, *RefreshAddRewardRes) error
	// 更新用户的coin数值
	RefreshUserCoin(context.Context, *RefreshUserCoinReq, *RefreshUserCoinRes) error
	// 刷新识别奖励
	RefreshIdentReward(context.Context, *RefreshIdentRewardReq, *RefreshIdentRewardRes) error
	// 更新用户vip
	RefreshUserVip(context.Context, *RefreshUserVipReq, *RefreshUserVipRes) error
	// 更新用户积分
	RefreshUserScore(context.Context, *RefreshUserScoreReq, *RefreshUserScoreRes) error
}

func RegisterSensinodeServiceHandler(s server.Server, hdlr SensinodeServiceHandler, opts ...server.HandlerOption) error {
	type sensinodeService interface {
		RefreshModifyReward(ctx context.Context, in *RefreshModifyRewardReq, out *RefreshModifyRewardRes) error
		RefreshAddReward(ctx context.Context, in *RefreshAddRewardReq, out *RefreshAddRewardRes) error
		RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, out *RefreshUserCoinRes) error
		RefreshIdentReward(ctx context.Context, in *RefreshIdentRewardReq, out *RefreshIdentRewardRes) error
		RefreshUserVip(ctx context.Context, in *RefreshUserVipReq, out *RefreshUserVipRes) error
		RefreshUserScore(ctx context.Context, in *RefreshUserScoreReq, out *RefreshUserScoreRes) error
	}
	type SensinodeService struct {
		sensinodeService
	}
	h := &sensinodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SensinodeService{h}, opts...))
}

type sensinodeServiceHandler struct {
	SensinodeServiceHandler
}

func (h *sensinodeServiceHandler) RefreshModifyReward(ctx context.Context, in *RefreshModifyRewardReq, out *RefreshModifyRewardRes) error {
	return h.SensinodeServiceHandler.RefreshModifyReward(ctx, in, out)
}

func (h *sensinodeServiceHandler) RefreshAddReward(ctx context.Context, in *RefreshAddRewardReq, out *RefreshAddRewardRes) error {
	return h.SensinodeServiceHandler.RefreshAddReward(ctx, in, out)
}

func (h *sensinodeServiceHandler) RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, out *RefreshUserCoinRes) error {
	return h.SensinodeServiceHandler.RefreshUserCoin(ctx, in, out)
}

func (h *sensinodeServiceHandler) RefreshIdentReward(ctx context.Context, in *RefreshIdentRewardReq, out *RefreshIdentRewardRes) error {
	return h.SensinodeServiceHandler.RefreshIdentReward(ctx, in, out)
}

func (h *sensinodeServiceHandler) RefreshUserVip(ctx context.Context, in *RefreshUserVipReq, out *RefreshUserVipRes) error {
	return h.SensinodeServiceHandler.RefreshUserVip(ctx, in, out)
}

func (h *sensinodeServiceHandler) RefreshUserScore(ctx context.Context, in *RefreshUserScoreReq, out *RefreshUserScoreRes) error {
	return h.SensinodeServiceHandler.RefreshUserScore(ctx, in, out)
}
