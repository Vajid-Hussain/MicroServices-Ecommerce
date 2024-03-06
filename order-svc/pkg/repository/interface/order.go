package interfaceRepository_order_svc

import (
	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	responsemodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/responseModel"
)

type IOrderRepository interface {
	CreateOrder(order *requestmodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error)
	AddProdutToOrderProductTable( *requestmodel_order_svc.UserCart,  *responsemodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error)
}
