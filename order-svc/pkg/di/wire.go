package di_order_svc

import (
	product_clind_in_order "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/clind"
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	server_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/server"
)

func InitializeOrderServer(config *config_order_svc.Config) (*server_order_svc.OrderServer, error) {

	clind, err := product_clind_in_order.InitProductClind(config)
	if err != nil {
		return nil, err
	}

	orderHandler := server_order_svc.NewOrderService(clind.Clind)

	return orderHandler, nil
}
