package interfaceUsecase_order_svc

import (
	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	responsemodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/responseModel"
)

type IOrderUseCase interface {
	CreateOrder(order *requestmodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error)
	OrderShowcase(string) (*[]responsemodel_order_svc.OrderShowcase, error)
	GetOrderSecret(string, string) (string, error)
	UpdatePaymentStatus( string) error
}
