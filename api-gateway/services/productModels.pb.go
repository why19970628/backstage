// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: productModels.proto

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

type ProductModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	// @inject_tag: json:"category_id"
	CategoryID uint32 `protobuf:"varint,3,opt,name=CategoryID,proto3" json:"category_id"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,4,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"info"
	Info string `protobuf:"bytes,5,opt,name=Info,proto3" json:"info"`
	// @inject_tag: json:"img_path"
	ImgPath string `protobuf:"bytes,6,opt,name=ImgPath,proto3" json:"img_path"`
	// @inject_tag: json:"price"
	Price string `protobuf:"bytes,7,opt,name=Price,proto3" json:"price"`
	// @inject_tag: json:"discount_price"
	DiscountPrice string `protobuf:"bytes,8,opt,name=DiscountPrice,proto3" json:"discount_price"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,9,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,10,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @inject_tag: json:"deleted_at"
	DeletedAt int64 `protobuf:"varint,11,opt,name=DeletedAt,proto3" json:"deleted_at"`
}

func (x *ProductModel) Reset() {
	*x = ProductModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductModel) ProtoMessage() {}

func (x *ProductModel) ProtoReflect() protoreflect.Message {
	mi := &file_productModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductModel.ProtoReflect.Descriptor instead.
func (*ProductModel) Descriptor() ([]byte, []int) {
	return file_productModels_proto_rawDescGZIP(), []int{0}
}

func (x *ProductModel) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductModel) GetCategoryID() uint32 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

func (x *ProductModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductModel) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *ProductModel) GetImgPath() string {
	if x != nil {
		return x.ImgPath
	}
	return ""
}

func (x *ProductModel) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductModel) GetDiscountPrice() string {
	if x != nil {
		return x.DiscountPrice
	}
	return ""
}

func (x *ProductModel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProductModel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ProductModel) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_productModels_proto protoreflect.FileDescriptor

var file_productModels_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22,
	0xac, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x49, 0x6d, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_productModels_proto_rawDescOnce sync.Once
	file_productModels_proto_rawDescData = file_productModels_proto_rawDesc
)

func file_productModels_proto_rawDescGZIP() []byte {
	file_productModels_proto_rawDescOnce.Do(func() {
		file_productModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_productModels_proto_rawDescData)
	})
	return file_productModels_proto_rawDescData
}

var file_productModels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_productModels_proto_goTypes = []interface{}{
	(*ProductModel)(nil), // 0: services.ProductModel
}
var file_productModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_productModels_proto_init() }
func file_productModels_proto_init() {
	if File_productModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_productModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductModel); i {
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
			RawDescriptor: file_productModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_productModels_proto_goTypes,
		DependencyIndexes: file_productModels_proto_depIdxs,
		MessageInfos:      file_productModels_proto_msgTypes,
	}.Build()
	File_productModels_proto = out.File
	file_productModels_proto_rawDesc = nil
	file_productModels_proto_goTypes = nil
	file_productModels_proto_depIdxs = nil
}
