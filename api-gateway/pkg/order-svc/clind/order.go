package order_clind

import (
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClind struct {
	Clind pb.OrderServiceClient
}

func InitOrderClind(config config.Config) (*OrderClind, error) {
	cc, err := grpc.Dial(config.OrderPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := OrderClind{Clind: pb.NewOrderServiceClient(cc)}
	return &c, nil
}
