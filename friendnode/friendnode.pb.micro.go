// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: friendnode.proto

package friendnode

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

// Api Endpoints for FriendnodeService service

func NewFriendnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FriendnodeService service

type FriendnodeService interface {
	// --------------- 好友接口 ----------- //
	// 添加好友
	AddFriends(ctx context.Context, in *AddFriendReq, opts ...client.CallOption) (*AddFriendRes, error)
	// 同意好友
	AgreeFriends(ctx context.Context, in *AgreeFriendReq, opts ...client.CallOption) (*AgreeFriendRes, error)
	//  请求好友列表
	FriendsList(ctx context.Context, in *FriendListReq, opts ...client.CallOption) (*FriendListRes, error)
	// 添加组
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...client.CallOption) (*CreateGroupRes, error)
	// 重命名组
	RenameGroup(ctx context.Context, in *RenameGroupReq, opts ...client.CallOption) (*RenameGroupRes, error)
	// 添加好友备注
	AddNoteFriend(ctx context.Context, in *AddFriendNoteReq, opts ...client.CallOption) (*AddFriendNoteRes, error)
	// 移动到新的组
	MoveToNewGroup(ctx context.Context, in *MoveGroupReq, opts ...client.CallOption) (*MoveGroupRes, error)
	// 删除好友
	RemoveFriend(ctx context.Context, in *RemoveFriendReq, opts ...client.CallOption) (*RemoveFriendRes, error)
	// 共享
	Share(ctx context.Context, in *ShareReq, opts ...client.CallOption) (*ShareRes, error)
	// 获取共享信息
	GetShare(ctx context.Context, in *GetShareReq, opts ...client.CallOption) (*GetShareRes, error)
	// 获取共享列表信息
	ShareList(ctx context.Context, in *ShareListReq, opts ...client.CallOption) (*ShareListRes, error)
	// 获取共享数量
	ShareCount(ctx context.Context, in *ShareCountReq, opts ...client.CallOption) (*ShareCountRes, error)
}

type friendnodeService struct {
	c    client.Client
	name string
}

func NewFriendnodeService(name string, c client.Client) FriendnodeService {
	return &friendnodeService{
		c:    c,
		name: name,
	}
}

func (c *friendnodeService) AddFriends(ctx context.Context, in *AddFriendReq, opts ...client.CallOption) (*AddFriendRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.AddFriends", in)
	out := new(AddFriendRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) AgreeFriends(ctx context.Context, in *AgreeFriendReq, opts ...client.CallOption) (*AgreeFriendRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.AgreeFriends", in)
	out := new(AgreeFriendRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) FriendsList(ctx context.Context, in *FriendListReq, opts ...client.CallOption) (*FriendListRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.FriendsList", in)
	out := new(FriendListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...client.CallOption) (*CreateGroupRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.CreateGroup", in)
	out := new(CreateGroupRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) RenameGroup(ctx context.Context, in *RenameGroupReq, opts ...client.CallOption) (*RenameGroupRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.RenameGroup", in)
	out := new(RenameGroupRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) AddNoteFriend(ctx context.Context, in *AddFriendNoteReq, opts ...client.CallOption) (*AddFriendNoteRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.AddNoteFriend", in)
	out := new(AddFriendNoteRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) MoveToNewGroup(ctx context.Context, in *MoveGroupReq, opts ...client.CallOption) (*MoveGroupRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.MoveToNewGroup", in)
	out := new(MoveGroupRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) RemoveFriend(ctx context.Context, in *RemoveFriendReq, opts ...client.CallOption) (*RemoveFriendRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.RemoveFriend", in)
	out := new(RemoveFriendRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) Share(ctx context.Context, in *ShareReq, opts ...client.CallOption) (*ShareRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.Share", in)
	out := new(ShareRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) GetShare(ctx context.Context, in *GetShareReq, opts ...client.CallOption) (*GetShareRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.GetShare", in)
	out := new(GetShareRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) ShareList(ctx context.Context, in *ShareListReq, opts ...client.CallOption) (*ShareListRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.ShareList", in)
	out := new(ShareListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendnodeService) ShareCount(ctx context.Context, in *ShareCountReq, opts ...client.CallOption) (*ShareCountRes, error) {
	req := c.c.NewRequest(c.name, "FriendnodeService.ShareCount", in)
	out := new(ShareCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FriendnodeService service

type FriendnodeServiceHandler interface {
	// --------------- 好友接口 ----------- //
	// 添加好友
	AddFriends(context.Context, *AddFriendReq, *AddFriendRes) error
	// 同意好友
	AgreeFriends(context.Context, *AgreeFriendReq, *AgreeFriendRes) error
	//  请求好友列表
	FriendsList(context.Context, *FriendListReq, *FriendListRes) error
	// 添加组
	CreateGroup(context.Context, *CreateGroupReq, *CreateGroupRes) error
	// 重命名组
	RenameGroup(context.Context, *RenameGroupReq, *RenameGroupRes) error
	// 添加好友备注
	AddNoteFriend(context.Context, *AddFriendNoteReq, *AddFriendNoteRes) error
	// 移动到新的组
	MoveToNewGroup(context.Context, *MoveGroupReq, *MoveGroupRes) error
	// 删除好友
	RemoveFriend(context.Context, *RemoveFriendReq, *RemoveFriendRes) error
	// 共享
	Share(context.Context, *ShareReq, *ShareRes) error
	// 获取共享信息
	GetShare(context.Context, *GetShareReq, *GetShareRes) error
	// 获取共享列表信息
	ShareList(context.Context, *ShareListReq, *ShareListRes) error
	// 获取共享数量
	ShareCount(context.Context, *ShareCountReq, *ShareCountRes) error
}

func RegisterFriendnodeServiceHandler(s server.Server, hdlr FriendnodeServiceHandler, opts ...server.HandlerOption) error {
	type friendnodeService interface {
		AddFriends(ctx context.Context, in *AddFriendReq, out *AddFriendRes) error
		AgreeFriends(ctx context.Context, in *AgreeFriendReq, out *AgreeFriendRes) error
		FriendsList(ctx context.Context, in *FriendListReq, out *FriendListRes) error
		CreateGroup(ctx context.Context, in *CreateGroupReq, out *CreateGroupRes) error
		RenameGroup(ctx context.Context, in *RenameGroupReq, out *RenameGroupRes) error
		AddNoteFriend(ctx context.Context, in *AddFriendNoteReq, out *AddFriendNoteRes) error
		MoveToNewGroup(ctx context.Context, in *MoveGroupReq, out *MoveGroupRes) error
		RemoveFriend(ctx context.Context, in *RemoveFriendReq, out *RemoveFriendRes) error
		Share(ctx context.Context, in *ShareReq, out *ShareRes) error
		GetShare(ctx context.Context, in *GetShareReq, out *GetShareRes) error
		ShareList(ctx context.Context, in *ShareListReq, out *ShareListRes) error
		ShareCount(ctx context.Context, in *ShareCountReq, out *ShareCountRes) error
	}
	type FriendnodeService struct {
		friendnodeService
	}
	h := &friendnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FriendnodeService{h}, opts...))
}

type friendnodeServiceHandler struct {
	FriendnodeServiceHandler
}

func (h *friendnodeServiceHandler) AddFriends(ctx context.Context, in *AddFriendReq, out *AddFriendRes) error {
	return h.FriendnodeServiceHandler.AddFriends(ctx, in, out)
}

func (h *friendnodeServiceHandler) AgreeFriends(ctx context.Context, in *AgreeFriendReq, out *AgreeFriendRes) error {
	return h.FriendnodeServiceHandler.AgreeFriends(ctx, in, out)
}

func (h *friendnodeServiceHandler) FriendsList(ctx context.Context, in *FriendListReq, out *FriendListRes) error {
	return h.FriendnodeServiceHandler.FriendsList(ctx, in, out)
}

func (h *friendnodeServiceHandler) CreateGroup(ctx context.Context, in *CreateGroupReq, out *CreateGroupRes) error {
	return h.FriendnodeServiceHandler.CreateGroup(ctx, in, out)
}

func (h *friendnodeServiceHandler) RenameGroup(ctx context.Context, in *RenameGroupReq, out *RenameGroupRes) error {
	return h.FriendnodeServiceHandler.RenameGroup(ctx, in, out)
}

func (h *friendnodeServiceHandler) AddNoteFriend(ctx context.Context, in *AddFriendNoteReq, out *AddFriendNoteRes) error {
	return h.FriendnodeServiceHandler.AddNoteFriend(ctx, in, out)
}

func (h *friendnodeServiceHandler) MoveToNewGroup(ctx context.Context, in *MoveGroupReq, out *MoveGroupRes) error {
	return h.FriendnodeServiceHandler.MoveToNewGroup(ctx, in, out)
}

func (h *friendnodeServiceHandler) RemoveFriend(ctx context.Context, in *RemoveFriendReq, out *RemoveFriendRes) error {
	return h.FriendnodeServiceHandler.RemoveFriend(ctx, in, out)
}

func (h *friendnodeServiceHandler) Share(ctx context.Context, in *ShareReq, out *ShareRes) error {
	return h.FriendnodeServiceHandler.Share(ctx, in, out)
}

func (h *friendnodeServiceHandler) GetShare(ctx context.Context, in *GetShareReq, out *GetShareRes) error {
	return h.FriendnodeServiceHandler.GetShare(ctx, in, out)
}

func (h *friendnodeServiceHandler) ShareList(ctx context.Context, in *ShareListReq, out *ShareListRes) error {
	return h.FriendnodeServiceHandler.ShareList(ctx, in, out)
}

func (h *friendnodeServiceHandler) ShareCount(ctx context.Context, in *ShareCountReq, out *ShareCountRes) error {
	return h.FriendnodeServiceHandler.ShareCount(ctx, in, out)
}
