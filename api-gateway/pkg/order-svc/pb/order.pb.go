// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/order-svc/pb/order.proto

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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OderType string `protobuf:"bytes,2,opt,name=OderType,proto3" json:"OderType,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *OrderRequest) GetOderType() string {
	if x != nil {
		return x.OderType
	}
	return ""
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OrderShowCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *OrderShowCaseRequest) Reset() {
	*x = OrderShowCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShowCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShowCaseRequest) ProtoMessage() {}

func (x *OrderShowCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShowCaseRequest.ProtoReflect.Descriptor instead.
func (*OrderShowCaseRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderShowCaseRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type OrderShowCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderShowcase `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderShowCaseResponse) Reset() {
	*x = OrderShowCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShowCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShowCaseResponse) ProtoMessage() {}

func (x *OrderShowCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShowCaseResponse.ProtoReflect.Descriptor instead.
func (*OrderShowCaseResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderShowCaseResponse) GetOrders() []*OrderShowcase {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderShowcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingleOrderId string `protobuf:"bytes,1,opt,name=single_order_id,json=singleOrderId,proto3" json:"single_order_id,omitempty"`
	OrderId       string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InventoryId   string `protobuf:"bytes,4,opt,name=inventory_id,json=inventoryId,proto3" json:"inventory_id,omitempty"`
	Price         uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	SalePrice     uint32 `protobuf:"varint,6,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
	OrderStatus   string `protobuf:"bytes,7,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	PaymentStatus string `protobuf:"bytes,8,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	Quantity      uint32 `protobuf:"varint,9,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderShowcase) Reset() {
	*x = OrderShowcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderShowcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderShowcase) ProtoMessage() {}

func (x *OrderShowcase) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderShowcase.ProtoReflect.Descriptor instead.
func (*OrderShowcase) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderShowcase) GetSingleOrderId() string {
	if x != nil {
		return x.SingleOrderId
	}
	return ""
}

func (x *OrderShowcase) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderShowcase) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderShowcase) GetInventoryId() string {
	if x != nil {
		return x.InventoryId
	}
	return ""
}

func (x *OrderShowcase) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderShowcase) GetSalePrice() uint32 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *OrderShowcase) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *OrderShowcase) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *OrderShowcase) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID string `protobuf:"bytes,2,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PaymentRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIDSecret string `protobuf:"bytes,1,opt,name=OrderIDSecret,proto3" json:"OrderIDSecret,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_order_svc_pb_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_order_svc_pb_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_pkg_order_svc_pb_order_proto_rawDescGZIP(), []int{6}
}

func (x *PaymentResponse) GetOrderIDSecret() string {
	if x != nil {
		return x.OrderIDSecret
	}
	return ""
}

var File_pkg_order_svc_pb_order_proto protoreflect.FileDescriptor

var file_pkg_order_svc_pb_order_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x76, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x4f, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4f, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x77, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x77, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xa9, 0x02, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x0f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x32, 0xdb, 0x01, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x77, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x77, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_order_svc_pb_order_proto_rawDescOnce sync.Once
	file_pkg_order_svc_pb_order_proto_rawDescData = file_pkg_order_svc_pb_order_proto_rawDesc
)

func file_pkg_order_svc_pb_order_proto_rawDescGZIP() []byte {
	file_pkg_order_svc_pb_order_proto_rawDescOnce.Do(func() {
		file_pkg_order_svc_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_order_svc_pb_order_proto_rawDescData)
	})
	return file_pkg_order_svc_pb_order_proto_rawDescData
}

var file_pkg_order_svc_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_order_svc_pb_order_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),          // 0: order.orderRequest
	(*OrderResponse)(nil),         // 1: order.orderResponse
	(*OrderShowCaseRequest)(nil),  // 2: order.orderShowCaseRequest
	(*OrderShowCaseResponse)(nil), // 3: order.orderShowCaseResponse
	(*OrderShowcase)(nil),         // 4: order.OrderShowcase
	(*PaymentRequest)(nil),        // 5: order.PaymentRequest
	(*PaymentResponse)(nil),       // 6: order.paymentResponse
}
var file_pkg_order_svc_pb_order_proto_depIdxs = []int32{
	4, // 0: order.orderShowCaseResponse.orders:type_name -> order.OrderShowcase
	0, // 1: order.orderService.OrderCreation:input_type -> order.orderRequest
	2, // 2: order.orderService.GetAllOrders:input_type -> order.orderShowCaseRequest
	5, // 3: order.orderService.OnlinePayment:input_type -> order.PaymentRequest
	1, // 4: order.orderService.OrderCreation:output_type -> order.orderResponse
	3, // 5: order.orderService.GetAllOrders:output_type -> order.orderShowCaseResponse
	6, // 6: order.orderService.OnlinePayment:output_type -> order.paymentResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_order_svc_pb_order_proto_init() }
func file_pkg_order_svc_pb_order_proto_init() {
	if File_pkg_order_svc_pb_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_order_svc_pb_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShowCaseRequest); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShowCaseResponse); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderShowcase); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
		file_pkg_order_svc_pb_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
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
			RawDescriptor: file_pkg_order_svc_pb_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_order_svc_pb_order_proto_goTypes,
		DependencyIndexes: file_pkg_order_svc_pb_order_proto_depIdxs,
		MessageInfos:      file_pkg_order_svc_pb_order_proto_msgTypes,
	}.Build()
	File_pkg_order_svc_pb_order_proto = out.File
	file_pkg_order_svc_pb_order_proto_rawDesc = nil
	file_pkg_order_svc_pb_order_proto_goTypes = nil
	file_pkg_order_svc_pb_order_proto_depIdxs = nil
}
