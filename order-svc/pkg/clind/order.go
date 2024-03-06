package product_clind_in_order

import (
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	"github.com/vajid-hussain/mobile-mart-microservice-order/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClind struct {
	Clind pb.ProductServiceClient
}

func InitProductClind(config *config_order_svc.Config) (*OrderClind, error) {
	cc, err := grpc.Dial(config.ProductPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &OrderClind{Clind: pb.NewProductServiceClient(cc)}, nil
}
