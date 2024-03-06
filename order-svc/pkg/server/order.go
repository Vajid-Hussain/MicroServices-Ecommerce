package server_order_svc

import (
	"context"
	"fmt"

	"github.com/vajid-hussain/mobile-mart-microservice-order/pkg/pb"
)

type OrderServer struct {
	clind pb.ProductServiceClient
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(clind pb.ProductServiceClient) *OrderServer {
	return &OrderServer{clind: clind}
}

func (u *OrderServer) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	result, err := u.clind.FetchCartForOrder(context.Background(), &pb.GetCartForOrderRequest{
		UserID: req.UserID,
	})

	fmt.Println("---", result, err)

	if err != nil {
		return &pb.OrderResponse{
			Message: "no order created something wend wrong",
		}, err
	}

	return &pb.OrderResponse{
		Message: "order craete successtully please pay the bill",
	}, nil
}
