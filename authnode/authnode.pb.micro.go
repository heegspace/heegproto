// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authnode.proto

package authnode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "github.com/heegspace/heegproto/common"
	_ "github.com/heegspace/heegproto/rescode"
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

// Api Endpoints for AuthnodeService service

func NewAuthnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthnodeService service

type AuthnodeService interface {
	// 验证token
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...client.CallOption) (*VerifyTokenRes, error)
	// 验证是否是管理用户
	AdminRole(ctx context.Context, in *AdminRoleReq, opts ...client.CallOption) (*AdminRoleRes, error)
	// 验证是否是合作用户
	CooperRole(ctx context.Context, in *CooperRoleReq, opts ...client.CallOption) (*CooperRoleRes, error)
}

type authnodeService struct {
	c    client.Client
	name string
}

func NewAuthnodeService(name string, c client.Client) AuthnodeService {
	return &authnodeService{
		c:    c,
		name: name,
	}
}

func (c *authnodeService) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...client.CallOption) (*VerifyTokenRes, error) {
	req := c.c.NewRequest(c.name, "AuthnodeService.verify_token", in)
	out := new(VerifyTokenRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnodeService) AdminRole(ctx context.Context, in *AdminRoleReq, opts ...client.CallOption) (*AdminRoleRes, error) {
	req := c.c.NewRequest(c.name, "AuthnodeService.admin_role", in)
	out := new(AdminRoleRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnodeService) CooperRole(ctx context.Context, in *CooperRoleReq, opts ...client.CallOption) (*CooperRoleRes, error) {
	req := c.c.NewRequest(c.name, "AuthnodeService.cooper_role", in)
	out := new(CooperRoleRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthnodeService service

type AuthnodeServiceHandler interface {
	// 验证token
	VerifyToken(context.Context, *VerifyTokenReq, *VerifyTokenRes) error
	// 验证是否是管理用户
	AdminRole(context.Context, *AdminRoleReq, *AdminRoleRes) error
	// 验证是否是合作用户
	CooperRole(context.Context, *CooperRoleReq, *CooperRoleRes) error
}

func RegisterAuthnodeServiceHandler(s server.Server, hdlr AuthnodeServiceHandler, opts ...server.HandlerOption) error {
	type authnodeService interface {
		VerifyToken(ctx context.Context, in *VerifyTokenReq, out *VerifyTokenRes) error
		AdminRole(ctx context.Context, in *AdminRoleReq, out *AdminRoleRes) error
		CooperRole(ctx context.Context, in *CooperRoleReq, out *CooperRoleRes) error
	}
	type AuthnodeService struct {
		authnodeService
	}
	h := &authnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthnodeService{h}, opts...))
}

type authnodeServiceHandler struct {
	AuthnodeServiceHandler
}

func (h *authnodeServiceHandler) VerifyToken(ctx context.Context, in *VerifyTokenReq, out *VerifyTokenRes) error {
	return h.AuthnodeServiceHandler.VerifyToken(ctx, in, out)
}

func (h *authnodeServiceHandler) AdminRole(ctx context.Context, in *AdminRoleReq, out *AdminRoleRes) error {
	return h.AuthnodeServiceHandler.AdminRole(ctx, in, out)
}

func (h *authnodeServiceHandler) CooperRole(ctx context.Context, in *CooperRoleReq, out *CooperRoleRes) error {
	return h.AuthnodeServiceHandler.CooperRole(ctx, in, out)
}