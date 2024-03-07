package repository_order_svc

import (
	"errors"
	"time"

	requestmodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/requestModel"
	responsemodel_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/model/responseModel"
	interfaceRepository_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/repository/interface"
	resCustomError_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel/custom_error"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaceRepository_order_svc.IOrderRepository {
	return &OrderRepository{DB: db}
}

func (d *OrderRepository) CreateOrder(order *requestmodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error) {

	var orderSucess responsemodel_order_svc.UserCart
	query := "INSERT INTO orders (user_id, payment_method, order_id_online, intent_payment_id) VALUES(?, ?, ?, ?) RETURNING*"
	result := d.DB.Raw(query, order.UserID, order.PaymentMethod, order.PaymentID, order.IntentUniquePaymentID).Scan(&orderSucess)
	if result.Error != nil {
		return nil, errors.New("face some issue while creating order")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError_product_svc.ErrNoRowAffected
	}

	return &orderSucess, nil
}

func (d *OrderRepository) AddProdutToOrderProductTable(order *requestmodel_order_svc.UserCart, orderCreated *responsemodel_order_svc.UserCart) (*responsemodel_order_svc.UserCart, error) {

	today := time.Now().Format("2006-01-02 15:04:05")

	for _, data := range order.Cart {
		var orderProduct responsemodel_order_svc.CartInventory
		query := "INSERT INTO order_products (order_id, inventory_id, quantity, order_date, order_status,payment_status, price, discount,payable_amount, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *"
		d.DB.Raw(query, orderCreated.OrderID, data.InventoryID, data.Quantity, today, order.OrderStatus, order.PaymentStatus, data.Price, data.Discount, data.FinalPrice, orderCreated.UserID).Scan(&orderProduct)
		orderCreated.Cart = append(orderCreated.Cart, orderProduct)
	}
	return orderCreated, nil
}

func (d *OrderRepository) GetOrderShowcase(userID string) (*[]responsemodel_order_svc.OrderShowcase, error) {

	var OrderShowcase []responsemodel_order_svc.OrderShowcase
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id  WHERE orders.user_id=? ORDER BY order_products.item_id DESC "
	result := d.DB.Raw(query, userID).Scan(&OrderShowcase)
	if result.Error != nil {
		return nil, errors.New("face some issue while order showcase")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError_product_svc.ErrNoRowAffected
	}
	return &OrderShowcase, nil
}

func (d *OrderRepository) GetOrderSecret(userID string, orderID string) (string, error) {
	var paymentSecret string
	query := "select order_id_online  from orders where id = $1 and user_id= $2"
	result := d.DB.Raw(query, orderID, userID).Scan(&paymentSecret)
	if result.Error != nil {
		return "", errors.New("face some issue while order showcase")
	}
	if result.RowsAffected == 0 {
		return "", resCustomError_product_svc.ErrNoRowAffected
	}

	return paymentSecret, nil
}

func (d *OrderRepository) UpdatePaymentStatus(IntentPaymentID string) error {
	query := "UPDATE order_products SET payment_status = 'success', order_status = 'processing' FROM orders WHERE orders.id = order_products.order_id AND orders.intent_payment_id=$1"
	result := d.DB.Exec(query, IntentPaymentID)
	if result.Error != nil {
		return errors.New("face some issue while update payment status")
	}
	if result.RowsAffected == 0 {
		return resCustomError_product_svc.ErrNoRowAffected
	}

	return nil
}
