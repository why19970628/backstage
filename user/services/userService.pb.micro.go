// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userService.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	AdminLogin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminDetailResponse, error)
	AdminRegister(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminDetailResponse, error)
	GetUsersList(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UsersListResponse, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserDetailResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AdminLogin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminDetailResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.AdminLogin", in)
	out := new(AdminDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AdminRegister(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminDetailResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.AdminRegister", in)
	out := new(AdminDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUsersList(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UsersListResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUsersList", in)
	out := new(UsersListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserDetailResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(UserDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	AdminLogin(context.Context, *AdminRequest, *AdminDetailResponse) error
	AdminRegister(context.Context, *AdminRequest, *AdminDetailResponse) error
	GetUsersList(context.Context, *UserRequest, *UsersListResponse) error
	GetUser(context.Context, *UserRequest, *UserDetailResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		AdminLogin(ctx context.Context, in *AdminRequest, out *AdminDetailResponse) error
		AdminRegister(ctx context.Context, in *AdminRequest, out *AdminDetailResponse) error
		GetUsersList(ctx context.Context, in *UserRequest, out *UsersListResponse) error
		GetUser(ctx context.Context, in *UserRequest, out *UserDetailResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) AdminLogin(ctx context.Context, in *AdminRequest, out *AdminDetailResponse) error {
	return h.UserServiceHandler.AdminLogin(ctx, in, out)
}

func (h *userServiceHandler) AdminRegister(ctx context.Context, in *AdminRequest, out *AdminDetailResponse) error {
	return h.UserServiceHandler.AdminRegister(ctx, in, out)
}

func (h *userServiceHandler) GetUsersList(ctx context.Context, in *UserRequest, out *UsersListResponse) error {
	return h.UserServiceHandler.GetUsersList(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *UserRequest, out *UserDetailResponse) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}
