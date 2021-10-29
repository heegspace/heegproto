// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: certnode.proto

package certnode

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

// Api Endpoints for CertnodeService service

func NewCertnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CertnodeService service

type CertnodeService interface {
	// 提交实名
	SubmitCert(ctx context.Context, in *SubmitCertReq, opts ...client.CallOption) (*SubmitCertRes, error)
	// 获取实名信息
	CertInfo(ctx context.Context, in *CertInfoReq, opts ...client.CallOption) (*CertInfoRes, error)
	// 实名日志记录
	CertFlow(ctx context.Context, in *CertFlowReq, opts ...client.CallOption) (*CertFlowRes, error)
	// 取消实名
	CertCancel(ctx context.Context, in *CertCancelReq, opts ...client.CallOption) (*CertCancelRes, error)
	//--------------需要授权--------------//
	// 审核通过
	CertApproved(ctx context.Context, in *CertApprovedReq, opts ...client.CallOption) (*CertApprovedRes, error)
	// 实名失败
	CertRefuse(ctx context.Context, in *CertRefuseReq, opts ...client.CallOption) (*CertRefuseRes, error)
	// 实名缓存
	CertCache(ctx context.Context, in *CertCacheReq, opts ...client.CallOption) (*CertCacheRes, error)
	// 实名状态
	CertStatus(ctx context.Context, in *CertStatusReq, opts ...client.CallOption) (*CertStatusRes, error)
}

type certnodeService struct {
	c    client.Client
	name string
}

func NewCertnodeService(name string, c client.Client) CertnodeService {
	return &certnodeService{
		c:    c,
		name: name,
	}
}

func (c *certnodeService) SubmitCert(ctx context.Context, in *SubmitCertReq, opts ...client.CallOption) (*SubmitCertRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.SubmitCert", in)
	out := new(SubmitCertRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertInfo(ctx context.Context, in *CertInfoReq, opts ...client.CallOption) (*CertInfoRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertInfo", in)
	out := new(CertInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertFlow(ctx context.Context, in *CertFlowReq, opts ...client.CallOption) (*CertFlowRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertFlow", in)
	out := new(CertFlowRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertCancel(ctx context.Context, in *CertCancelReq, opts ...client.CallOption) (*CertCancelRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertCancel", in)
	out := new(CertCancelRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertApproved(ctx context.Context, in *CertApprovedReq, opts ...client.CallOption) (*CertApprovedRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertApproved", in)
	out := new(CertApprovedRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertRefuse(ctx context.Context, in *CertRefuseReq, opts ...client.CallOption) (*CertRefuseRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertRefuse", in)
	out := new(CertRefuseRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertCache(ctx context.Context, in *CertCacheReq, opts ...client.CallOption) (*CertCacheRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertCache", in)
	out := new(CertCacheRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certnodeService) CertStatus(ctx context.Context, in *CertStatusReq, opts ...client.CallOption) (*CertStatusRes, error) {
	req := c.c.NewRequest(c.name, "CertnodeService.CertStatus", in)
	out := new(CertStatusRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertnodeService service

type CertnodeServiceHandler interface {
	// 提交实名
	SubmitCert(context.Context, *SubmitCertReq, *SubmitCertRes) error
	// 获取实名信息
	CertInfo(context.Context, *CertInfoReq, *CertInfoRes) error
	// 实名日志记录
	CertFlow(context.Context, *CertFlowReq, *CertFlowRes) error
	// 取消实名
	CertCancel(context.Context, *CertCancelReq, *CertCancelRes) error
	//--------------需要授权--------------//
	// 审核通过
	CertApproved(context.Context, *CertApprovedReq, *CertApprovedRes) error
	// 实名失败
	CertRefuse(context.Context, *CertRefuseReq, *CertRefuseRes) error
	// 实名缓存
	CertCache(context.Context, *CertCacheReq, *CertCacheRes) error
	// 实名状态
	CertStatus(context.Context, *CertStatusReq, *CertStatusRes) error
}

func RegisterCertnodeServiceHandler(s server.Server, hdlr CertnodeServiceHandler, opts ...server.HandlerOption) error {
	type certnodeService interface {
		SubmitCert(ctx context.Context, in *SubmitCertReq, out *SubmitCertRes) error
		CertInfo(ctx context.Context, in *CertInfoReq, out *CertInfoRes) error
		CertFlow(ctx context.Context, in *CertFlowReq, out *CertFlowRes) error
		CertCancel(ctx context.Context, in *CertCancelReq, out *CertCancelRes) error
		CertApproved(ctx context.Context, in *CertApprovedReq, out *CertApprovedRes) error
		CertRefuse(ctx context.Context, in *CertRefuseReq, out *CertRefuseRes) error
		CertCache(ctx context.Context, in *CertCacheReq, out *CertCacheRes) error
		CertStatus(ctx context.Context, in *CertStatusReq, out *CertStatusRes) error
	}
	type CertnodeService struct {
		certnodeService
	}
	h := &certnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CertnodeService{h}, opts...))
}

type certnodeServiceHandler struct {
	CertnodeServiceHandler
}

func (h *certnodeServiceHandler) SubmitCert(ctx context.Context, in *SubmitCertReq, out *SubmitCertRes) error {
	return h.CertnodeServiceHandler.SubmitCert(ctx, in, out)
}

func (h *certnodeServiceHandler) CertInfo(ctx context.Context, in *CertInfoReq, out *CertInfoRes) error {
	return h.CertnodeServiceHandler.CertInfo(ctx, in, out)
}

func (h *certnodeServiceHandler) CertFlow(ctx context.Context, in *CertFlowReq, out *CertFlowRes) error {
	return h.CertnodeServiceHandler.CertFlow(ctx, in, out)
}

func (h *certnodeServiceHandler) CertCancel(ctx context.Context, in *CertCancelReq, out *CertCancelRes) error {
	return h.CertnodeServiceHandler.CertCancel(ctx, in, out)
}

func (h *certnodeServiceHandler) CertApproved(ctx context.Context, in *CertApprovedReq, out *CertApprovedRes) error {
	return h.CertnodeServiceHandler.CertApproved(ctx, in, out)
}

func (h *certnodeServiceHandler) CertRefuse(ctx context.Context, in *CertRefuseReq, out *CertRefuseRes) error {
	return h.CertnodeServiceHandler.CertRefuse(ctx, in, out)
}

func (h *certnodeServiceHandler) CertCache(ctx context.Context, in *CertCacheReq, out *CertCacheRes) error {
	return h.CertnodeServiceHandler.CertCache(ctx, in, out)
}

func (h *certnodeServiceHandler) CertStatus(ctx context.Context, in *CertStatusReq, out *CertStatusRes) error {
	return h.CertnodeServiceHandler.CertStatus(ctx, in, out)
}
