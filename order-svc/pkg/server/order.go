package server_order_svc

import (
	"context"

	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	"github.com/vajid-hussain/mobile-mart-microservice-order/pkg/pb"
	interfaceUsecase_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/usecase/interface"
)

type OrderServer struct {
	clind   pb.ProductServiceClient
	usecase interfaceUsecase_order_svc.IOrderUseCase
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(clind pb.ProductServiceClient, userCase interfaceUsecase_order_svc.IOrderUseCase) *OrderServer {
	return &OrderServer{clind: clind, usecase: userCase}
}

func (u *OrderServer) OrderCreation(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	var orders = requestmodel_order_svc.UserCart{}
	result, err := u.clind.FetchCartForOrder(context.Background(), &pb.GetCartForOrderRequest{
		UserID: req.UserID,
	})

	orders.InventoryCount = uint(result.InventoryCount)
	orders.TotalPrice = uint(result.TotalPrice)
	orders.UserID = result.UserId
	orders.PaymentMethod = req.OderType

	for _, val := range result.Cart {
		var order requestmodel_order_svc.CartInventory
		order.CategoryDiscount = uint(val.Discount)
		order.CategoryID = val.CategoryId
		order.Discount = uint(val.Discount)
		order.FinalPrice = uint(val.FinalPrice)
		order.InventoryID = val.InventoryId
		order.Price = uint(val.Price)
		order.Productname = val.Productname
		order.Quantity = uint(val.Quantity)
		order.Saleprice = uint(val.Saleprice)
		order.Title = val.Title
		order.Units = val.Units
		orders.Cart = append(orders.Cart, order)
	}

	_, err = u.usecase.CreateOrder(&orders)
	if err != nil {
		return &pb.OrderResponse{
			Message: "no order created something wend wrong",
		}, err
	}

	return &pb.OrderResponse{
		Message: "order craete successfully please pay the bill",
	}, nil
}

func (u *OrderServer) GetAllOrders(ctx context.Context, req *pb.OrderShowCaseRequest) (*pb.OrderShowCaseResponse, error) {
	result, err := u.usecase.OrderShowcase(req.UserID)
	if err != nil {
		return nil, err
	}

	var orders []*pb.OrderShowcase
	for _, o := range *result {
		order := &pb.OrderShowcase{
			SingleOrderId: o.SingleOrderID,
			OrderId:       o.ID,
			UserId:        o.UserID,
			InventoryId:   o.InventoryID,
			Price:         uint32(o.Price),
			SalePrice:     uint32(o.Saleprice),
			OrderStatus:   o.OrderStatus,
			PaymentStatus: o.PaymentStatus,
			Quantity:      uint32(o.Quantity),
		}
		orders = append(orders, order)
	}

	return &pb.OrderShowCaseResponse{Orders: orders}, nil

}

func (u *OrderServer) OnlinePayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	result, err := u.usecase.GetOrderSecret(req.UserID, req.OrderID)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentResponse{
		OrderIDSecret: result,
	}, nil
}

func (u *OrderServer) UpdataPaymentStatus(ctx context.Context, req *pb.UpdataPaymentStatusRequest) (*pb.UpdataPaymentStatusResponse, error) {
	err := u.usecase.UpdatePaymentStatus(req.IntentPaymentID)
	if err != nil {
		return nil, err
	}

	return &pb.UpdataPaymentStatusResponse{
		Message: "payment succesfully updated. you order shortly deliverd",
	}, nil
}

