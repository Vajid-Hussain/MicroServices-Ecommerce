package di_order_svc

import (
	product_clind_in_order "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/clind"
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	db_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/db"
	repository_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/repository"
	server_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/server"
	usecase_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/usecase"
)

func InitializeOrderServer(config *config_order_svc.Config) (*server_order_svc.OrderServer, error) {

	clind, err := product_clind_in_order.InitProductClind(config)
	if err != nil {
		return nil, err
	}

	db, err:= db_order_svc.InitDB(config)
	if err != nil {
		return nil, err
	}

	OrderRepository := repository_order_svc.NewOrderRepository(db)
	usecaseOrder := usecase_order_svc.NewOrderUseCase(OrderRepository)
	orderHandler := server_order_svc.NewOrderService(clind.Clind, usecaseOrder)

	return orderHandler, nil
}
