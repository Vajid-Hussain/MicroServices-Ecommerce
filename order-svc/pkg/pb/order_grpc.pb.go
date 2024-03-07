// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrderService_OrderCreation_FullMethodName       = "/order.orderService/OrderCreation"
	OrderService_GetAllOrders_FullMethodName        = "/order.orderService/GetAllOrders"
	OrderService_OnlinePayment_FullMethodName       = "/order.orderService/OnlinePayment"
	OrderService_UpdataPaymentStatus_FullMethodName = "/order.orderService/UpdataPaymentStatus"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	OrderCreation(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetAllOrders(ctx context.Context, in *OrderShowCaseRequest, opts ...grpc.CallOption) (*OrderShowCaseResponse, error)
	OnlinePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	UpdataPaymentStatus(ctx context.Context, in *UpdataPaymentStatusRequest, opts ...grpc.CallOption) (*UpdataPaymentStatusResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) OrderCreation(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_OrderCreation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetAllOrders(ctx context.Context, in *OrderShowCaseRequest, opts ...grpc.CallOption) (*OrderShowCaseResponse, error) {
	out := new(OrderShowCaseResponse)
	err := c.cc.Invoke(ctx, OrderService_GetAllOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OnlinePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, OrderService_OnlinePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdataPaymentStatus(ctx context.Context, in *UpdataPaymentStatusRequest, opts ...grpc.CallOption) (*UpdataPaymentStatusResponse, error) {
	out := new(UpdataPaymentStatusResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdataPaymentStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	OrderCreation(context.Context, *OrderRequest) (*OrderResponse, error)
	GetAllOrders(context.Context, *OrderShowCaseRequest) (*OrderShowCaseResponse, error)
	OnlinePayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	UpdataPaymentStatus(context.Context, *UpdataPaymentStatusRequest) (*UpdataPaymentStatusResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) OrderCreation(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCreation not implemented")
}
func (UnimplementedOrderServiceServer) GetAllOrders(context.Context, *OrderShowCaseRequest) (*OrderShowCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedOrderServiceServer) OnlinePayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlinePayment not implemented")
}
func (UnimplementedOrderServiceServer) UpdataPaymentStatus(context.Context, *UpdataPaymentStatusRequest) (*UpdataPaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdataPaymentStatus not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_OrderCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderCreation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderCreation(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderShowCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetAllOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetAllOrders(ctx, req.(*OrderShowCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OnlinePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OnlinePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OnlinePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OnlinePayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdataPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdataPaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdataPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdataPaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdataPaymentStatus(ctx, req.(*UpdataPaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.orderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderCreation",
			Handler:    _OrderService_OrderCreation_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _OrderService_GetAllOrders_Handler,
		},
		{
			MethodName: "OnlinePayment",
			Handler:    _OrderService_OnlinePayment_Handler,
		},
		{
			MethodName: "UpdataPaymentStatus",
			Handler:    _OrderService_UpdataPaymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/order.proto",
}
