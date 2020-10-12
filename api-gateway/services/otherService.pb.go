// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: otherService.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CarouselRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"start" form:"start" uri:"start"
	Start uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start" form:"start" uri:"start"`
	// @inject_tag: json:"limit" form:"limit" uri:"limit"
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" form:"limit" uri:"limit"`
	// @inject_tag: json:"carousel_id" form:"carousel_id" uri:"carousel_id"
	CarouselID uint32 `protobuf:"varint,3,opt,name=CarouselID,proto3" json:"carousel_id" form:"carousel_id" uri:"carousel_id"`
	// @inject_tag: json:"img_path" form:"img_path" uri:"img_path"
	ImgPath string `protobuf:"bytes,4,opt,name=ImgPath,proto3" json:"img_path" form:"img_path" uri:"img_path"`
	// @inject_tag: json:"product_id" form:"product_id" uri:"product_id"
	ProductID uint32 `protobuf:"varint,5,opt,name=ProductID,proto3" json:"product_id" form:"product_id" uri:"product_id"`
}

func (x *CarouselRequest) Reset() {
	*x = CarouselRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselRequest) ProtoMessage() {}

func (x *CarouselRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselRequest.ProtoReflect.Descriptor instead.
func (*CarouselRequest) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{0}
}

func (x *CarouselRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *CarouselRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CarouselRequest) GetCarouselID() uint32 {
	if x != nil {
		return x.CarouselID
	}
	return 0
}

func (x *CarouselRequest) GetImgPath() string {
	if x != nil {
		return x.ImgPath
	}
	return ""
}

func (x *CarouselRequest) GetProductID() uint32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type CarouselsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselsList []*CarouselModel `protobuf:"bytes,1,rep,name=CarouselsList,proto3" json:"CarouselsList,omitempty"`
	// @inject_tag: json:"count"
	Count uint32 `protobuf:"varint,2,opt,name=Count,proto3" json:"count"`
}

func (x *CarouselsListResponse) Reset() {
	*x = CarouselsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselsListResponse) ProtoMessage() {}

func (x *CarouselsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselsListResponse.ProtoReflect.Descriptor instead.
func (*CarouselsListResponse) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{1}
}

func (x *CarouselsListResponse) GetCarouselsList() []*CarouselModel {
	if x != nil {
		return x.CarouselsList
	}
	return nil
}

func (x *CarouselsListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CarouselDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarouselDetail *CarouselModel `protobuf:"bytes,1,opt,name=CarouselDetail,proto3" json:"CarouselDetail,omitempty"`
}

func (x *CarouselDetailResponse) Reset() {
	*x = CarouselDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarouselDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarouselDetailResponse) ProtoMessage() {}

func (x *CarouselDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarouselDetailResponse.ProtoReflect.Descriptor instead.
func (*CarouselDetailResponse) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{2}
}

func (x *CarouselDetailResponse) GetCarouselDetail() *CarouselModel {
	if x != nil {
		return x.CarouselDetail
	}
	return nil
}

type NoticeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"notice_id" form:"notice_id" uri:"notice_id"
	NoticeID uint32 `protobuf:"varint,1,opt,name=NoticeID,proto3" json:"notice_id" form:"notice_id" uri:"notice_id"`
	// @inject_tag: json:"text" form:"text" uri:"text"
	Text string `protobuf:"bytes,2,opt,name=Text,proto3" json:"text" form:"text" uri:"text"`
}

func (x *NoticeRequest) Reset() {
	*x = NoticeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeRequest) ProtoMessage() {}

func (x *NoticeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeRequest.ProtoReflect.Descriptor instead.
func (*NoticeRequest) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{3}
}

func (x *NoticeRequest) GetNoticeID() uint32 {
	if x != nil {
		return x.NoticeID
	}
	return 0
}

func (x *NoticeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type NoticeDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoticeDetail *NoticeModel `protobuf:"bytes,1,opt,name=NoticeDetail,proto3" json:"NoticeDetail,omitempty"`
}

func (x *NoticeDetailResponse) Reset() {
	*x = NoticeDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeDetailResponse) ProtoMessage() {}

func (x *NoticeDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeDetailResponse.ProtoReflect.Descriptor instead.
func (*NoticeDetailResponse) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{4}
}

func (x *NoticeDetailResponse) GetNoticeDetail() *NoticeModel {
	if x != nil {
		return x.NoticeDetail
	}
	return nil
}

type CategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"start" form:"start" uri:"start"
	Start uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start" form:"start" uri:"start"`
	// @inject_tag: json:"limit" form:"limit" uri:"limit"
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" form:"limit" uri:"limit"`
	// @inject_tag: json:"id" form:"id" uri:"id"
	ID uint32 `protobuf:"varint,3,opt,name=ID,proto3" json:"id" form:"id" uri:"id"`
	// @inject_tag: json:"category_id" form:"category_id" uri:"category_id"
	CategoryID uint32 `protobuf:"varint,4,opt,name=CategoryID,proto3" json:"category_id" form:"category_id" uri:"category_id"`
	// @inject_tag: json:"category_name" form:"category_name" uri:"category_name"
	CategoryName string `protobuf:"bytes,5,opt,name=CategoryName,proto3" json:"category_name" form:"category_name" uri:"category_name"`
	// @inject_tag: json:"num" form:"num" uri:"num"
	Num uint32 `protobuf:"varint,6,opt,name=Num,proto3" json:"num" form:"num" uri:"num"`
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *CategoryRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CategoryRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CategoryRequest) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

func (x *CategoryRequest) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryRequest) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CategoriesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoriesList []*CategoryModel `protobuf:"bytes,1,rep,name=CategoriesList,proto3" json:"CategoriesList,omitempty"`
	// @inject_tag: json:"count"
	Count uint32 `protobuf:"varint,2,opt,name=Count,proto3" json:"count"`
}

func (x *CategoriesListResponse) Reset() {
	*x = CategoriesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesListResponse) ProtoMessage() {}

func (x *CategoriesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesListResponse.ProtoReflect.Descriptor instead.
func (*CategoriesListResponse) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{6}
}

func (x *CategoriesListResponse) GetCategoriesList() []*CategoryModel {
	if x != nil {
		return x.CategoriesList
	}
	return nil
}

func (x *CategoriesListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CategoryDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryDetail *CategoryModel `protobuf:"bytes,1,opt,name=CategoryDetail,proto3" json:"CategoryDetail,omitempty"`
}

func (x *CategoryDetailResponse) Reset() {
	*x = CategoryDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otherService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDetailResponse) ProtoMessage() {}

func (x *CategoryDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otherService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDetailResponse.ProtoReflect.Descriptor instead.
func (*CategoryDetailResponse) Descriptor() ([]byte, []int) {
	return file_otherService_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryDetailResponse) GetCategoryDetail() *CategoryModel {
	if x != nil {
		return x.CategoryDetail
	}
	return nil
}

var File_otherService_proto protoreflect.FileDescriptor

var file_otherService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x6c, 0x0a, 0x15, 0x43, 0x61, 0x72,
	0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x0d, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x61, 0x72, 0x6f, 0x75,
	0x73, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x0e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x3f, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x51, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0x6f, 0x0a, 0x16,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a,
	0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32, 0xa5, 0x07, 0x0a, 0x0c, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otherService_proto_rawDescOnce sync.Once
	file_otherService_proto_rawDescData = file_otherService_proto_rawDesc
)

func file_otherService_proto_rawDescGZIP() []byte {
	file_otherService_proto_rawDescOnce.Do(func() {
		file_otherService_proto_rawDescData = protoimpl.X.CompressGZIP(file_otherService_proto_rawDescData)
	})
	return file_otherService_proto_rawDescData
}

var file_otherService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_otherService_proto_goTypes = []interface{}{
	(*CarouselRequest)(nil),        // 0: services.CarouselRequest
	(*CarouselsListResponse)(nil),  // 1: services.CarouselsListResponse
	(*CarouselDetailResponse)(nil), // 2: services.CarouselDetailResponse
	(*NoticeRequest)(nil),          // 3: services.NoticeRequest
	(*NoticeDetailResponse)(nil),   // 4: services.NoticeDetailResponse
	(*CategoryRequest)(nil),        // 5: services.CategoryRequest
	(*CategoriesListResponse)(nil), // 6: services.CategoriesListResponse
	(*CategoryDetailResponse)(nil), // 7: services.CategoryDetailResponse
	(*CarouselModel)(nil),          // 8: services.CarouselModel
	(*NoticeModel)(nil),            // 9: services.NoticeModel
	(*CategoryModel)(nil),          // 10: services.CategoryModel
}
var file_otherService_proto_depIdxs = []int32{
	8,  // 0: services.CarouselsListResponse.CarouselsList:type_name -> services.CarouselModel
	8,  // 1: services.CarouselDetailResponse.CarouselDetail:type_name -> services.CarouselModel
	9,  // 2: services.NoticeDetailResponse.NoticeDetail:type_name -> services.NoticeModel
	10, // 3: services.CategoriesListResponse.CategoriesList:type_name -> services.CategoryModel
	10, // 4: services.CategoryDetailResponse.CategoryDetail:type_name -> services.CategoryModel
	0,  // 5: services.OtherService.GetCarouselsList:input_type -> services.CarouselRequest
	0,  // 6: services.OtherService.GetCarousel:input_type -> services.CarouselRequest
	0,  // 7: services.OtherService.UpdateCarousel:input_type -> services.CarouselRequest
	3,  // 8: services.OtherService.CreateNotice:input_type -> services.NoticeRequest
	3,  // 9: services.OtherService.GetNotice:input_type -> services.NoticeRequest
	3,  // 10: services.OtherService.UpdateNotice:input_type -> services.NoticeRequest
	3,  // 11: services.OtherService.DeleteNotice:input_type -> services.NoticeRequest
	5,  // 12: services.OtherService.CreateCategory:input_type -> services.CategoryRequest
	5,  // 13: services.OtherService.GetCategoriesList:input_type -> services.CategoryRequest
	5,  // 14: services.OtherService.GetCategory:input_type -> services.CategoryRequest
	5,  // 15: services.OtherService.UpdateCategory:input_type -> services.CategoryRequest
	5,  // 16: services.OtherService.DeleteCategory:input_type -> services.CategoryRequest
	1,  // 17: services.OtherService.GetCarouselsList:output_type -> services.CarouselsListResponse
	2,  // 18: services.OtherService.GetCarousel:output_type -> services.CarouselDetailResponse
	2,  // 19: services.OtherService.UpdateCarousel:output_type -> services.CarouselDetailResponse
	4,  // 20: services.OtherService.CreateNotice:output_type -> services.NoticeDetailResponse
	4,  // 21: services.OtherService.GetNotice:output_type -> services.NoticeDetailResponse
	4,  // 22: services.OtherService.UpdateNotice:output_type -> services.NoticeDetailResponse
	4,  // 23: services.OtherService.DeleteNotice:output_type -> services.NoticeDetailResponse
	7,  // 24: services.OtherService.CreateCategory:output_type -> services.CategoryDetailResponse
	6,  // 25: services.OtherService.GetCategoriesList:output_type -> services.CategoriesListResponse
	7,  // 26: services.OtherService.GetCategory:output_type -> services.CategoryDetailResponse
	7,  // 27: services.OtherService.UpdateCategory:output_type -> services.CategoryDetailResponse
	7,  // 28: services.OtherService.DeleteCategory:output_type -> services.CategoryDetailResponse
	17, // [17:29] is the sub-list for method output_type
	5,  // [5:17] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_otherService_proto_init() }
func file_otherService_proto_init() {
	if File_otherService_proto != nil {
		return
	}
	file_otherModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_otherService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselsListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarouselDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_otherService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_otherService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_otherService_proto_goTypes,
		DependencyIndexes: file_otherService_proto_depIdxs,
		MessageInfos:      file_otherService_proto_msgTypes,
	}.Build()
	File_otherService_proto = out.File
	file_otherService_proto_rawDesc = nil
	file_otherService_proto_goTypes = nil
	file_otherService_proto_depIdxs = nil
}
