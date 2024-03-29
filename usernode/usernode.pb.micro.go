// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usernode.proto

package usernode

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

// Api Endpoints for UsernodeService service

func NewUsernodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UsernodeService service

type UsernodeService interface {
	// 更新用户信息
	UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, opts ...client.CallOption) (*UpdateUserinfoRes, error)
	// 更新身份证号
	UpdateCardid(ctx context.Context, in *UpdateCardidReq, opts ...client.CallOption) (*UpdateUserRes, error)
	// 更新关注对象
	UpdateAttention(ctx context.Context, in *UpdateAttentionReq, opts ...client.CallOption) (*UpdateUserRes, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserInfoRes, error)
	// 用户缓存
	UserCache(ctx context.Context, in *UserCacheReq, opts ...client.CallOption) (*UserCacheRes, error)
	// 用户积分获取
	UserScore(ctx context.Context, in *UserScoreReq, opts ...client.CallOption) (*UserScoreRes, error)
	// 用户VIP获取
	UserVip(ctx context.Context, in *UserVipReq, opts ...client.CallOption) (*UserVipRes, error)
}

type usernodeService struct {
	c    client.Client
	name string
}

func NewUsernodeService(name string, c client.Client) UsernodeService {
	return &usernodeService{
		c:    c,
		name: name,
	}
}

func (c *usernodeService) UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, opts ...client.CallOption) (*UpdateUserinfoRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UpdateUserInfo", in)
	out := new(UpdateUserinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UpdateCardid(ctx context.Context, in *UpdateCardidReq, opts ...client.CallOption) (*UpdateUserRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UpdateCardid", in)
	out := new(UpdateUserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UpdateAttention(ctx context.Context, in *UpdateAttentionReq, opts ...client.CallOption) (*UpdateUserRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UpdateAttention", in)
	out := new(UpdateUserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserInfoRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UserInfo", in)
	out := new(UserInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UserCache(ctx context.Context, in *UserCacheReq, opts ...client.CallOption) (*UserCacheRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UserCache", in)
	out := new(UserCacheRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UserScore(ctx context.Context, in *UserScoreReq, opts ...client.CallOption) (*UserScoreRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UserScore", in)
	out := new(UserScoreRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usernodeService) UserVip(ctx context.Context, in *UserVipReq, opts ...client.CallOption) (*UserVipRes, error) {
	req := c.c.NewRequest(c.name, "UsernodeService.UserVip", in)
	out := new(UserVipRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UsernodeService service

type UsernodeServiceHandler interface {
	// 更新用户信息
	UpdateUserInfo(context.Context, *UpdateUserinfoReq, *UpdateUserinfoRes) error
	// 更新身份证号
	UpdateCardid(context.Context, *UpdateCardidReq, *UpdateUserRes) error
	// 更新关注对象
	UpdateAttention(context.Context, *UpdateAttentionReq, *UpdateUserRes) error
	// 获取用户信息
	UserInfo(context.Context, *UserInfoReq, *UserInfoRes) error
	// 用户缓存
	UserCache(context.Context, *UserCacheReq, *UserCacheRes) error
	// 用户积分获取
	UserScore(context.Context, *UserScoreReq, *UserScoreRes) error
	// 用户VIP获取
	UserVip(context.Context, *UserVipReq, *UserVipRes) error
}

func RegisterUsernodeServiceHandler(s server.Server, hdlr UsernodeServiceHandler, opts ...server.HandlerOption) error {
	type usernodeService interface {
		UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, out *UpdateUserinfoRes) error
		UpdateCardid(ctx context.Context, in *UpdateCardidReq, out *UpdateUserRes) error
		UpdateAttention(ctx context.Context, in *UpdateAttentionReq, out *UpdateUserRes) error
		UserInfo(ctx context.Context, in *UserInfoReq, out *UserInfoRes) error
		UserCache(ctx context.Context, in *UserCacheReq, out *UserCacheRes) error
		UserScore(ctx context.Context, in *UserScoreReq, out *UserScoreRes) error
		UserVip(ctx context.Context, in *UserVipReq, out *UserVipRes) error
	}
	type UsernodeService struct {
		usernodeService
	}
	h := &usernodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UsernodeService{h}, opts...))
}

type usernodeServiceHandler struct {
	UsernodeServiceHandler
}

func (h *usernodeServiceHandler) UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, out *UpdateUserinfoRes) error {
	return h.UsernodeServiceHandler.UpdateUserInfo(ctx, in, out)
}

func (h *usernodeServiceHandler) UpdateCardid(ctx context.Context, in *UpdateCardidReq, out *UpdateUserRes) error {
	return h.UsernodeServiceHandler.UpdateCardid(ctx, in, out)
}

func (h *usernodeServiceHandler) UpdateAttention(ctx context.Context, in *UpdateAttentionReq, out *UpdateUserRes) error {
	return h.UsernodeServiceHandler.UpdateAttention(ctx, in, out)
}

func (h *usernodeServiceHandler) UserInfo(ctx context.Context, in *UserInfoReq, out *UserInfoRes) error {
	return h.UsernodeServiceHandler.UserInfo(ctx, in, out)
}

func (h *usernodeServiceHandler) UserCache(ctx context.Context, in *UserCacheReq, out *UserCacheRes) error {
	return h.UsernodeServiceHandler.UserCache(ctx, in, out)
}

func (h *usernodeServiceHandler) UserScore(ctx context.Context, in *UserScoreReq, out *UserScoreRes) error {
	return h.UsernodeServiceHandler.UserScore(ctx, in, out)
}

func (h *usernodeServiceHandler) UserVip(ctx context.Context, in *UserVipReq, out *UserVipRes) error {
	return h.UsernodeServiceHandler.UserVip(ctx, in, out)
}
