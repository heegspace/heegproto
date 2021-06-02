// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: notenode.proto

package notenode

import (
	_ "github.com/heegspace/heegproto/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Api Endpoints for NotenodeService service

func NewNotenodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NotenodeService service

type NotenodeService interface {
	//更新或者添加笔记信息
	UpdateNote(ctx context.Context, in *UpdateNoteReq, opts ...client.CallOption) (*NoteMetaRes, error)
	// 获取笔记列表
	NoteMetaList(ctx context.Context, in *NoteMetaListReq, opts ...client.CallOption) (*NoteMetaListRes, error)
	// 获取用户笔记数量
	NoteListCount(ctx context.Context, in *NoteListCountReq, opts ...client.CallOption) (*NoteListCountRes, error)
	// 获取笔记数据列
	NoteData(ctx context.Context, in *NoteDataReq, opts ...client.CallOption) (*NoteDataRes, error)
	// 获取笔记html数据
	NoteHtml(ctx context.Context, in *NoteHtmlReq, opts ...client.CallOption) (*NoteHtmlRes, error)
	// 更新笔记的协作者、标签、颜色
	NoteCooper(ctx context.Context, in *NoteCooperReq, opts ...client.CallOption) (*NoteCooperRes, error)
	NoteTag(ctx context.Context, in *NoteTagReq, opts ...client.CallOption) (*NoteTagRes, error)
	NoteBgcolor(ctx context.Context, in *NoteBgcolorReq, opts ...client.CallOption) (*NoteBgcolorRes, error)
	NoteTagAdd(ctx context.Context, in *NoteTagAddReq, opts ...client.CallOption) (*NoteTagAddRes, error)
	NoteTagList(ctx context.Context, in *NoteTagListReq, opts ...client.CallOption) (*NoteTagListRes, error)
}

type notenodeService struct {
	c    client.Client
	name string
}

func NewNotenodeService(name string, c client.Client) NotenodeService {
	return &notenodeService{
		c:    c,
		name: name,
	}
}

func (c *notenodeService) UpdateNote(ctx context.Context, in *UpdateNoteReq, opts ...client.CallOption) (*NoteMetaRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.updateNote", in)
	out := new(NoteMetaRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteMetaList(ctx context.Context, in *NoteMetaListReq, opts ...client.CallOption) (*NoteMetaListRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.noteMetaList", in)
	out := new(NoteMetaListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteListCount(ctx context.Context, in *NoteListCountReq, opts ...client.CallOption) (*NoteListCountRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.noteListCount", in)
	out := new(NoteListCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteData(ctx context.Context, in *NoteDataReq, opts ...client.CallOption) (*NoteDataRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.noteData", in)
	out := new(NoteDataRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteHtml(ctx context.Context, in *NoteHtmlReq, opts ...client.CallOption) (*NoteHtmlRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.noteHtml", in)
	out := new(NoteHtmlRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteCooper(ctx context.Context, in *NoteCooperReq, opts ...client.CallOption) (*NoteCooperRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.note_cooper", in)
	out := new(NoteCooperRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteTag(ctx context.Context, in *NoteTagReq, opts ...client.CallOption) (*NoteTagRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.note_tag", in)
	out := new(NoteTagRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteBgcolor(ctx context.Context, in *NoteBgcolorReq, opts ...client.CallOption) (*NoteBgcolorRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.note_bgcolor", in)
	out := new(NoteBgcolorRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteTagAdd(ctx context.Context, in *NoteTagAddReq, opts ...client.CallOption) (*NoteTagAddRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.note_tag_add", in)
	out := new(NoteTagAddRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notenodeService) NoteTagList(ctx context.Context, in *NoteTagListReq, opts ...client.CallOption) (*NoteTagListRes, error) {
	req := c.c.NewRequest(c.name, "NotenodeService.note_tag_list", in)
	out := new(NoteTagListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotenodeService service

type NotenodeServiceHandler interface {
	//更新或者添加笔记信息
	UpdateNote(context.Context, *UpdateNoteReq, *NoteMetaRes) error
	// 获取笔记列表
	NoteMetaList(context.Context, *NoteMetaListReq, *NoteMetaListRes) error
	// 获取用户笔记数量
	NoteListCount(context.Context, *NoteListCountReq, *NoteListCountRes) error
	// 获取笔记数据列
	NoteData(context.Context, *NoteDataReq, *NoteDataRes) error
	// 获取笔记html数据
	NoteHtml(context.Context, *NoteHtmlReq, *NoteHtmlRes) error
	// 更新笔记的协作者、标签、颜色
	NoteCooper(context.Context, *NoteCooperReq, *NoteCooperRes) error
	NoteTag(context.Context, *NoteTagReq, *NoteTagRes) error
	NoteBgcolor(context.Context, *NoteBgcolorReq, *NoteBgcolorRes) error
	NoteTagAdd(context.Context, *NoteTagAddReq, *NoteTagAddRes) error
	NoteTagList(context.Context, *NoteTagListReq, *NoteTagListRes) error
}

func RegisterNotenodeServiceHandler(s server.Server, hdlr NotenodeServiceHandler, opts ...server.HandlerOption) error {
	type notenodeService interface {
		UpdateNote(ctx context.Context, in *UpdateNoteReq, out *NoteMetaRes) error
		NoteMetaList(ctx context.Context, in *NoteMetaListReq, out *NoteMetaListRes) error
		NoteListCount(ctx context.Context, in *NoteListCountReq, out *NoteListCountRes) error
		NoteData(ctx context.Context, in *NoteDataReq, out *NoteDataRes) error
		NoteHtml(ctx context.Context, in *NoteHtmlReq, out *NoteHtmlRes) error
		NoteCooper(ctx context.Context, in *NoteCooperReq, out *NoteCooperRes) error
		NoteTag(ctx context.Context, in *NoteTagReq, out *NoteTagRes) error
		NoteBgcolor(ctx context.Context, in *NoteBgcolorReq, out *NoteBgcolorRes) error
		NoteTagAdd(ctx context.Context, in *NoteTagAddReq, out *NoteTagAddRes) error
		NoteTagList(ctx context.Context, in *NoteTagListReq, out *NoteTagListRes) error
	}
	type NotenodeService struct {
		notenodeService
	}
	h := &notenodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NotenodeService{h}, opts...))
}

type notenodeServiceHandler struct {
	NotenodeServiceHandler
}

func (h *notenodeServiceHandler) UpdateNote(ctx context.Context, in *UpdateNoteReq, out *NoteMetaRes) error {
	return h.NotenodeServiceHandler.UpdateNote(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteMetaList(ctx context.Context, in *NoteMetaListReq, out *NoteMetaListRes) error {
	return h.NotenodeServiceHandler.NoteMetaList(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteListCount(ctx context.Context, in *NoteListCountReq, out *NoteListCountRes) error {
	return h.NotenodeServiceHandler.NoteListCount(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteData(ctx context.Context, in *NoteDataReq, out *NoteDataRes) error {
	return h.NotenodeServiceHandler.NoteData(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteHtml(ctx context.Context, in *NoteHtmlReq, out *NoteHtmlRes) error {
	return h.NotenodeServiceHandler.NoteHtml(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteCooper(ctx context.Context, in *NoteCooperReq, out *NoteCooperRes) error {
	return h.NotenodeServiceHandler.NoteCooper(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteTag(ctx context.Context, in *NoteTagReq, out *NoteTagRes) error {
	return h.NotenodeServiceHandler.NoteTag(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteBgcolor(ctx context.Context, in *NoteBgcolorReq, out *NoteBgcolorRes) error {
	return h.NotenodeServiceHandler.NoteBgcolor(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteTagAdd(ctx context.Context, in *NoteTagAddReq, out *NoteTagAddRes) error {
	return h.NotenodeServiceHandler.NoteTagAdd(ctx, in, out)
}

func (h *notenodeServiceHandler) NoteTagList(ctx context.Context, in *NoteTagListReq, out *NoteTagListRes) error {
	return h.NotenodeServiceHandler.NoteTagList(ctx, in, out)
}