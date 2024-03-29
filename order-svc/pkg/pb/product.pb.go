// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/product.proto

package pb

import (
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

type CartInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Productname string `protobuf:"bytes,1,opt,name=productname,proto3" json:"productname,omitempty"`
	InventoryId string `protobuf:"bytes,2,opt,name=inventory_id,json=inventoryId,proto3" json:"inventory_id,omitempty"`
	CategoryId  string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Quantity    uint32 `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price       uint32 `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Discount    uint32 `protobuf:"varint,7,opt,name=discount,proto3" json:"discount,omitempty"`
	Saleprice   uint32 `protobuf:"varint,8,opt,name=saleprice,proto3" json:"saleprice,omitempty"`
	FinalPrice  uint32 `protobuf:"varint,10,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	Title       string `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	Units       uint64 `protobuf:"varint,12,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *CartInventory) Reset() {
	*x = CartInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartInventory) ProtoMessage() {}

func (x *CartInventory) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartInventory.ProtoReflect.Descriptor instead.
func (*CartInventory) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{0}
}

func (x *CartInventory) GetProductname() string {
	if x != nil {
		return x.Productname
	}
	return ""
}

func (x *CartInventory) GetInventoryId() string {
	if x != nil {
		return x.InventoryId
	}
	return ""
}

func (x *CartInventory) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CartInventory) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartInventory) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartInventory) GetDiscount() uint32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *CartInventory) GetSaleprice() uint32 {
	if x != nil {
		return x.Saleprice
	}
	return 0
}

func (x *CartInventory) GetFinalPrice() uint32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

func (x *CartInventory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CartInventory) GetUnits() uint64 {
	if x != nil {
		return x.Units
	}
	return 0
}

type GetCartForOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetCartForOrderRequest) Reset() {
	*x = GetCartForOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartForOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartForOrderRequest) ProtoMessage() {}

func (x *GetCartForOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartForOrderRequest.ProtoReflect.Descriptor instead.
func (*GetCartForOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetCartForOrderRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetCartForOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice     uint32           `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	InventoryCount uint32           `protobuf:"varint,3,opt,name=inventory_count,json=inventoryCount,proto3" json:"inventory_count,omitempty"`
	Cart           []*CartInventory `protobuf:"bytes,4,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *GetCartForOrderResponse) Reset() {
	*x = GetCartForOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartForOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartForOrderResponse) ProtoMessage() {}

func (x *GetCartForOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartForOrderResponse.ProtoReflect.Descriptor instead.
func (*GetCartForOrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetCartForOrderResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCartForOrderResponse) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *GetCartForOrderResponse) GetInventoryCount() uint32 {
	if x != nil {
		return x.InventoryCount
	}
	return 0
}

func (x *GetCartForOrderResponse) GetCart() []*CartInventory {
	if x != nil {
		return x.Cart
	}
	return nil
}

var File_pkg_pb_product_proto protoreflect.FileDescriptor

var file_pkg_pb_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0xae, 0x02, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0xa8, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x46, 0x6f,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x32, 0x6a, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x58, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x61, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_product_proto_rawDescOnce sync.Once
	file_pkg_pb_product_proto_rawDescData = file_pkg_pb_product_proto_rawDesc
)

func file_pkg_pb_product_proto_rawDescGZIP() []byte {
	file_pkg_pb_product_proto_rawDescOnce.Do(func() {
		file_pkg_pb_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_product_proto_rawDescData)
	})
	return file_pkg_pb_product_proto_rawDescData
}

var file_pkg_pb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_pb_product_proto_goTypes = []interface{}{
	(*CartInventory)(nil),           // 0: product.CartInventory
	(*GetCartForOrderRequest)(nil),  // 1: product.GetCartForOrderRequest
	(*GetCartForOrderResponse)(nil), // 2: product.GetCartForOrderResponse
}
var file_pkg_pb_product_proto_depIdxs = []int32{
	0, // 0: product.GetCartForOrderResponse.cart:type_name -> product.CartInventory
	1, // 1: product.ProductService.FetchCartForOrder:input_type -> product.GetCartForOrderRequest
	2, // 2: product.ProductService.FetchCartForOrder:output_type -> product.GetCartForOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_product_proto_init() }
func file_pkg_pb_product_proto_init() {
	if File_pkg_pb_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartInventory); i {
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
		file_pkg_pb_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartForOrderRequest); i {
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
		file_pkg_pb_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartForOrderResponse); i {
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
			RawDescriptor: file_pkg_pb_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_product_proto_goTypes,
		DependencyIndexes: file_pkg_pb_product_proto_depIdxs,
		MessageInfos:      file_pkg_pb_product_proto_msgTypes,
	}.Build()
	File_pkg_pb_product_proto = out.File
	file_pkg_pb_product_proto_rawDesc = nil
	file_pkg_pb_product_proto_goTypes = nil
	file_pkg_pb_product_proto_depIdxs = nil
}
