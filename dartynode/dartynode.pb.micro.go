// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dartynode.proto

package dartynode

import (
	_ "github.com/heegspace/heegproto/common"
	_ "github.com/heegspace/heegproto/rescode"
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DartynodeService service

func NewDartynodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DartynodeService service

type DartynodeService interface {
	// 登录微信
	LoginWechat(ctx context.Context, in *LoginWechatReq, opts ...client.CallOption) (*LoginWechatRes, error)
	// 刷新微信的token
	RefreshWechat(ctx context.Context, in *RefreshWechatReq, opts ...client.CallOption) (*RefreshWechatRes, error)
	// 退出微信登录
	LogoutWechat(ctx context.Context, in *LogoutWechatReq, opts ...client.CallOption) (*LogoutWechatRes, error)
	// 获取用户信息
	UserinfoWechat(ctx context.Context, in *UserinfoWechatReq, opts ...client.CallOption) (*UserinfoWechatRes, error)
	// 支付宝登陆
	LoginAlipay(ctx context.Context, in *LoginAlipayReq, opts ...client.CallOption) (*LoginAlipayRes, error)
	// 刷新alipay的token
	RefreshAlipay(ctx context.Context, in *RefreshAlipayReq, opts ...client.CallOption) (*RefreshAlipayRes, error)
	// 退出alipay登录
	LogoutAlipay(ctx context.Context, in *LogoutAlipayReq, opts ...client.CallOption) (*LogoutAlipayRes, error)
	// 获取用户信息
	UserinfoAlipay(ctx context.Context, in *UserinfoAlipayReq, opts ...client.CallOption) (*UserinfoAlipayRes, error)
	// 百度实体
	BaiduEntity(ctx context.Context, in *BaiduEntityReq, opts ...client.CallOption) (*BaiduEntityRes, error)
	// 试卷识别
	BaiduDocAnalysis(ctx context.Context, in *BaiduDocAnalysisReq, opts ...client.CallOption) (*BaiduDocAnalysisRes, error)
	// 身份证识别
	BaiduIdcardIdent(ctx context.Context, in *BaiduIdcardIdentReq, opts ...client.CallOption) (*BaiduIdcardIdentRes, error)
	// 检查文本是否违规
	BaiduTextCensor(ctx context.Context, in *BaiduTextCensorReq, opts ...client.CallOption) (*BaiduTextCensorRes, error)
	// 检查图像是否违规
	BaiduImgCensor(ctx context.Context, in *BaiduImgCensorReq, opts ...client.CallOption) (*BaiduImgCensorRes, error)
	// 微信支付
	WechatPay(ctx context.Context, in *WechatPayReq, opts ...client.CallOption) (*WechatPayRes, error)
	// 支付宝支付
	Alipay(ctx context.Context, in *AlipayReq, opts ...client.CallOption) (*AlipayRes, error)
}

type dartynodeService struct {
	c    client.Client
	name string
}

func NewDartynodeService(name string, c client.Client) DartynodeService {
	return &dartynodeService{
		c:    c,
		name: name,
	}
}

func (c *dartynodeService) LoginWechat(ctx context.Context, in *LoginWechatReq, opts ...client.CallOption) (*LoginWechatRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.LoginWechat", in)
	out := new(LoginWechatRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) RefreshWechat(ctx context.Context, in *RefreshWechatReq, opts ...client.CallOption) (*RefreshWechatRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.RefreshWechat", in)
	out := new(RefreshWechatRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) LogoutWechat(ctx context.Context, in *LogoutWechatReq, opts ...client.CallOption) (*LogoutWechatRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.LogoutWechat", in)
	out := new(LogoutWechatRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) UserinfoWechat(ctx context.Context, in *UserinfoWechatReq, opts ...client.CallOption) (*UserinfoWechatRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.UserinfoWechat", in)
	out := new(UserinfoWechatRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) LoginAlipay(ctx context.Context, in *LoginAlipayReq, opts ...client.CallOption) (*LoginAlipayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.LoginAlipay", in)
	out := new(LoginAlipayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) RefreshAlipay(ctx context.Context, in *RefreshAlipayReq, opts ...client.CallOption) (*RefreshAlipayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.RefreshAlipay", in)
	out := new(RefreshAlipayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) LogoutAlipay(ctx context.Context, in *LogoutAlipayReq, opts ...client.CallOption) (*LogoutAlipayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.LogoutAlipay", in)
	out := new(LogoutAlipayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) UserinfoAlipay(ctx context.Context, in *UserinfoAlipayReq, opts ...client.CallOption) (*UserinfoAlipayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.UserinfoAlipay", in)
	out := new(UserinfoAlipayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) BaiduEntity(ctx context.Context, in *BaiduEntityReq, opts ...client.CallOption) (*BaiduEntityRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.BaiduEntity", in)
	out := new(BaiduEntityRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) BaiduDocAnalysis(ctx context.Context, in *BaiduDocAnalysisReq, opts ...client.CallOption) (*BaiduDocAnalysisRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.BaiduDocAnalysis", in)
	out := new(BaiduDocAnalysisRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) BaiduIdcardIdent(ctx context.Context, in *BaiduIdcardIdentReq, opts ...client.CallOption) (*BaiduIdcardIdentRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.BaiduIdcardIdent", in)
	out := new(BaiduIdcardIdentRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) BaiduTextCensor(ctx context.Context, in *BaiduTextCensorReq, opts ...client.CallOption) (*BaiduTextCensorRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.BaiduTextCensor", in)
	out := new(BaiduTextCensorRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) BaiduImgCensor(ctx context.Context, in *BaiduImgCensorReq, opts ...client.CallOption) (*BaiduImgCensorRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.BaiduImgCensor", in)
	out := new(BaiduImgCensorRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) WechatPay(ctx context.Context, in *WechatPayReq, opts ...client.CallOption) (*WechatPayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.WechatPay", in)
	out := new(WechatPayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dartynodeService) Alipay(ctx context.Context, in *AlipayReq, opts ...client.CallOption) (*AlipayRes, error) {
	req := c.c.NewRequest(c.name, "DartynodeService.Alipay", in)
	out := new(AlipayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DartynodeService service

type DartynodeServiceHandler interface {
	// 登录微信
	LoginWechat(context.Context, *LoginWechatReq, *LoginWechatRes) error
	// 刷新微信的token
	RefreshWechat(context.Context, *RefreshWechatReq, *RefreshWechatRes) error
	// 退出微信登录
	LogoutWechat(context.Context, *LogoutWechatReq, *LogoutWechatRes) error
	// 获取用户信息
	UserinfoWechat(context.Context, *UserinfoWechatReq, *UserinfoWechatRes) error
	// 支付宝登陆
	LoginAlipay(context.Context, *LoginAlipayReq, *LoginAlipayRes) error
	// 刷新alipay的token
	RefreshAlipay(context.Context, *RefreshAlipayReq, *RefreshAlipayRes) error
	// 退出alipay登录
	LogoutAlipay(context.Context, *LogoutAlipayReq, *LogoutAlipayRes) error
	// 获取用户信息
	UserinfoAlipay(context.Context, *UserinfoAlipayReq, *UserinfoAlipayRes) error
	// 百度实体
	BaiduEntity(context.Context, *BaiduEntityReq, *BaiduEntityRes) error
	// 试卷识别
	BaiduDocAnalysis(context.Context, *BaiduDocAnalysisReq, *BaiduDocAnalysisRes) error
	// 身份证识别
	BaiduIdcardIdent(context.Context, *BaiduIdcardIdentReq, *BaiduIdcardIdentRes) error
	// 检查文本是否违规
	BaiduTextCensor(context.Context, *BaiduTextCensorReq, *BaiduTextCensorRes) error
	// 检查图像是否违规
	BaiduImgCensor(context.Context, *BaiduImgCensorReq, *BaiduImgCensorRes) error
	// 微信支付
	WechatPay(context.Context, *WechatPayReq, *WechatPayRes) error
	// 支付宝支付
	Alipay(context.Context, *AlipayReq, *AlipayRes) error
}

func RegisterDartynodeServiceHandler(s server.Server, hdlr DartynodeServiceHandler, opts ...server.HandlerOption) error {
	type dartynodeService interface {
		LoginWechat(ctx context.Context, in *LoginWechatReq, out *LoginWechatRes) error
		RefreshWechat(ctx context.Context, in *RefreshWechatReq, out *RefreshWechatRes) error
		LogoutWechat(ctx context.Context, in *LogoutWechatReq, out *LogoutWechatRes) error
		UserinfoWechat(ctx context.Context, in *UserinfoWechatReq, out *UserinfoWechatRes) error
		LoginAlipay(ctx context.Context, in *LoginAlipayReq, out *LoginAlipayRes) error
		RefreshAlipay(ctx context.Context, in *RefreshAlipayReq, out *RefreshAlipayRes) error
		LogoutAlipay(ctx context.Context, in *LogoutAlipayReq, out *LogoutAlipayRes) error
		UserinfoAlipay(ctx context.Context, in *UserinfoAlipayReq, out *UserinfoAlipayRes) error
		BaiduEntity(ctx context.Context, in *BaiduEntityReq, out *BaiduEntityRes) error
		BaiduDocAnalysis(ctx context.Context, in *BaiduDocAnalysisReq, out *BaiduDocAnalysisRes) error
		BaiduIdcardIdent(ctx context.Context, in *BaiduIdcardIdentReq, out *BaiduIdcardIdentRes) error
		BaiduTextCensor(ctx context.Context, in *BaiduTextCensorReq, out *BaiduTextCensorRes) error
		BaiduImgCensor(ctx context.Context, in *BaiduImgCensorReq, out *BaiduImgCensorRes) error
		WechatPay(ctx context.Context, in *WechatPayReq, out *WechatPayRes) error
		Alipay(ctx context.Context, in *AlipayReq, out *AlipayRes) error
	}
	type DartynodeService struct {
		dartynodeService
	}
	h := &dartynodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DartynodeService{h}, opts...))
}

type dartynodeServiceHandler struct {
	DartynodeServiceHandler
}

func (h *dartynodeServiceHandler) LoginWechat(ctx context.Context, in *LoginWechatReq, out *LoginWechatRes) error {
	return h.DartynodeServiceHandler.LoginWechat(ctx, in, out)
}

func (h *dartynodeServiceHandler) RefreshWechat(ctx context.Context, in *RefreshWechatReq, out *RefreshWechatRes) error {
	return h.DartynodeServiceHandler.RefreshWechat(ctx, in, out)
}

func (h *dartynodeServiceHandler) LogoutWechat(ctx context.Context, in *LogoutWechatReq, out *LogoutWechatRes) error {
	return h.DartynodeServiceHandler.LogoutWechat(ctx, in, out)
}

func (h *dartynodeServiceHandler) UserinfoWechat(ctx context.Context, in *UserinfoWechatReq, out *UserinfoWechatRes) error {
	return h.DartynodeServiceHandler.UserinfoWechat(ctx, in, out)
}

func (h *dartynodeServiceHandler) LoginAlipay(ctx context.Context, in *LoginAlipayReq, out *LoginAlipayRes) error {
	return h.DartynodeServiceHandler.LoginAlipay(ctx, in, out)
}

func (h *dartynodeServiceHandler) RefreshAlipay(ctx context.Context, in *RefreshAlipayReq, out *RefreshAlipayRes) error {
	return h.DartynodeServiceHandler.RefreshAlipay(ctx, in, out)
}

func (h *dartynodeServiceHandler) LogoutAlipay(ctx context.Context, in *LogoutAlipayReq, out *LogoutAlipayRes) error {
	return h.DartynodeServiceHandler.LogoutAlipay(ctx, in, out)
}

func (h *dartynodeServiceHandler) UserinfoAlipay(ctx context.Context, in *UserinfoAlipayReq, out *UserinfoAlipayRes) error {
	return h.DartynodeServiceHandler.UserinfoAlipay(ctx, in, out)
}

func (h *dartynodeServiceHandler) BaiduEntity(ctx context.Context, in *BaiduEntityReq, out *BaiduEntityRes) error {
	return h.DartynodeServiceHandler.BaiduEntity(ctx, in, out)
}

func (h *dartynodeServiceHandler) BaiduDocAnalysis(ctx context.Context, in *BaiduDocAnalysisReq, out *BaiduDocAnalysisRes) error {
	return h.DartynodeServiceHandler.BaiduDocAnalysis(ctx, in, out)
}

func (h *dartynodeServiceHandler) BaiduIdcardIdent(ctx context.Context, in *BaiduIdcardIdentReq, out *BaiduIdcardIdentRes) error {
	return h.DartynodeServiceHandler.BaiduIdcardIdent(ctx, in, out)
}

func (h *dartynodeServiceHandler) BaiduTextCensor(ctx context.Context, in *BaiduTextCensorReq, out *BaiduTextCensorRes) error {
	return h.DartynodeServiceHandler.BaiduTextCensor(ctx, in, out)
}

func (h *dartynodeServiceHandler) BaiduImgCensor(ctx context.Context, in *BaiduImgCensorReq, out *BaiduImgCensorRes) error {
	return h.DartynodeServiceHandler.BaiduImgCensor(ctx, in, out)
}

func (h *dartynodeServiceHandler) WechatPay(ctx context.Context, in *WechatPayReq, out *WechatPayRes) error {
	return h.DartynodeServiceHandler.WechatPay(ctx, in, out)
}

func (h *dartynodeServiceHandler) Alipay(ctx context.Context, in *AlipayReq, out *AlipayRes) error {
	return h.DartynodeServiceHandler.Alipay(ctx, in, out)
}
