package server_order_svc

import (
	"context"
	"fmt"

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

	fmt.Println("---", result, err)
	orders.InventoryCount = uint(result.InventoryCount)
	orders.TotalPrice = uint(result.TotalPrice)
	orders.UserID = result.UserId
	orders.PaymentMethod = req.OderType
	fmt.Println("****", orders)

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

	order, err:=u.usecase.CreateOrder(&orders)
	if err != nil {
		return &pb.OrderResponse{
			Message: "no order created something wend wrong",
		}, err
	}

	fmt.Println("---", order, err)

	return &pb.OrderResponse{
		Message: "order craete successtully please pay the bill",
	}, nil
}
