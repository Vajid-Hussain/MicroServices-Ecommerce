package usecase_order_svc

import (
	"fmt"

	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	responsemodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/responseModel"
	interfaceRepository_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/repository/interface"
	interfaceUsecase_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/usecase/interface"
)

type OrderUsecase struct {
	repo interfaceRepository_order_svc.IOrderRepository
}

func NewOrderUseCase(repository interfaceRepository_order_svc.IOrderRepository) interfaceUsecase_order_svc.IOrderUseCase {
	return &OrderUsecase{repo: repository}
}

func (r *OrderUsecase) CreateOrder(order *requestmodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error) {
	orderCreated, err := r.repo.CreateOrder(order)
	fmt.Println("---", orderCreated, err)
	order.OrderStatus="pending"
	order.PaymentStatus="pending"

	finalResult,err:=r.repo.AddProdutToOrderProductTable(order, orderCreated)
	if err!=nil{
		return nil,err
	}

	return finalResult,nil
}
