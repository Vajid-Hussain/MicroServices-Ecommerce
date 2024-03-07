package usecase_order_svc

import (
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	responsemodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/responseModel"
	interfaceRepository_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/repository/interface"
	interfaceUsecase_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/usecase/interface"
	utils_order "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/utils"
)

type OrderUsecase struct {
	repo   interfaceRepository_order_svc.IOrderRepository
	config *config_order_svc.Config
}

func NewOrderUseCase(repository interfaceRepository_order_svc.IOrderRepository, config *config_order_svc.Config) interfaceUsecase_order_svc.IOrderUseCase {
	return &OrderUsecase{repo: repository, config: config}
}

func (r *OrderUsecase) CreateOrder(order *requestmodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error) {

	intentClindSecret, intentPaymentID, err := utils_order.StripePaymentInitiation(int(order.TotalPrice)*100, r.config)
	if err != nil {
		return nil, err
	}

	order.PaymentID = intentClindSecret
	order.IntentUniquePaymentID= intentPaymentID

	orderCreated, err := r.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	order.OrderStatus = "pending"
	order.PaymentStatus = "pending"

	finalResult, err := r.repo.AddProdutToOrderProductTable(order, orderCreated)
	if err != nil {
		return nil, err
	}

	return finalResult, nil
}

func (r *OrderUsecase) OrderShowcase(userID string) (*[]responsemodel_order_svc.OrderShowcase, error) {
	abstractOrder, err := r.repo.GetOrderShowcase(userID)
	if err != nil {
		return nil, err
	}
	return abstractOrder, nil
}

func (r *OrderUsecase) GetOrderSecret(userID, orderID string) (string, error) {
	return r.repo.GetOrderSecret( userID,orderID)
}

func (r *OrderUsecase) UpdatePaymentStatus(intentPaymentID string) error{
	return r.repo.UpdatePaymentStatus(intentPaymentID)
}

