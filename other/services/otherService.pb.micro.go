// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: otherService.proto

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

// Client API for OtherService service

type OtherService interface {
	GetCarouselsList(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselsListResponse, error)
	UpdateCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error)
}

type otherService struct {
	c    client.Client
	name string
}

func NewOtherService(name string, c client.Client) OtherService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &otherService{
		c:    c,
		name: name,
	}
}

func (c *otherService) GetCarouselsList(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselsListResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetCarouselsList", in)
	out := new(CarouselsListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) UpdateCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.UpdateCarousel", in)
	out := new(CarouselDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OtherService service

type OtherServiceHandler interface {
	GetCarouselsList(context.Context, *CarouselRequest, *CarouselsListResponse) error
	UpdateCarousel(context.Context, *CarouselRequest, *CarouselDetailResponse) error
}

func RegisterOtherServiceHandler(s server.Server, hdlr OtherServiceHandler, opts ...server.HandlerOption) error {
	type otherService interface {
		GetCarouselsList(ctx context.Context, in *CarouselRequest, out *CarouselsListResponse) error
		UpdateCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error
	}
	type OtherService struct {
		otherService
	}
	h := &otherServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OtherService{h}, opts...))
}

type otherServiceHandler struct {
	OtherServiceHandler
}

func (h *otherServiceHandler) GetCarouselsList(ctx context.Context, in *CarouselRequest, out *CarouselsListResponse) error {
	return h.OtherServiceHandler.GetCarouselsList(ctx, in, out)
}

func (h *otherServiceHandler) UpdateCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error {
	return h.OtherServiceHandler.UpdateCarousel(ctx, in, out)
}
