// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/product_info.proto

package __

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_proto_product_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_product_info_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_product_info_proto protoreflect.FileDescriptor

const file_proto_product_info_proto_rawDesc = "" +
	"\n" +
	"\x18proto/product_info.proto\x12\tecommerce\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1cgoogle/api/annotations.proto\"e\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price2\xc2\x01\n" +
	"\vProductInfo\x12V\n" +
	"\n" +
	"addProduct\x12\x12.ecommerce.Product\x1a\x1c.google.protobuf.StringValue\"\x16\x82\xd3\xe4\x93\x02\x10:\x01*\"\v/v1/product\x12[\n" +
	"\n" +
	"getProduct\x12\x1c.google.protobuf.StringValue\x1a\x12.ecommerce.Product\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/product/{value}b\x06proto3"

var (
	file_proto_product_info_proto_rawDescOnce sync.Once
	file_proto_product_info_proto_rawDescData []byte
)

func file_proto_product_info_proto_rawDescGZIP() []byte {
	file_proto_product_info_proto_rawDescOnce.Do(func() {
		file_proto_product_info_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_product_info_proto_rawDesc), len(file_proto_product_info_proto_rawDesc)))
	})
	return file_proto_product_info_proto_rawDescData
}

var file_proto_product_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_product_info_proto_goTypes = []any{
	(*Product)(nil),                // 0: ecommerce.Product
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_proto_product_info_proto_depIdxs = []int32{
	0, // 0: ecommerce.ProductInfo.addProduct:input_type -> ecommerce.Product
	1, // 1: ecommerce.ProductInfo.getProduct:input_type -> google.protobuf.StringValue
	1, // 2: ecommerce.ProductInfo.addProduct:output_type -> google.protobuf.StringValue
	0, // 3: ecommerce.ProductInfo.getProduct:output_type -> ecommerce.Product
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_product_info_proto_init() }
func file_proto_product_info_proto_init() {
	if File_proto_product_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_product_info_proto_rawDesc), len(file_proto_product_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_info_proto_goTypes,
		DependencyIndexes: file_proto_product_info_proto_depIdxs,
		MessageInfos:      file_proto_product_info_proto_msgTypes,
	}.Build()
	File_proto_product_info_proto = out.File
	file_proto_product_info_proto_goTypes = nil
	file_proto_product_info_proto_depIdxs = nil
}
