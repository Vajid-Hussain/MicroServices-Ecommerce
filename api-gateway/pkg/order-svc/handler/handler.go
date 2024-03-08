package handler_order_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/pb"
	requestmodel_order_svc_clind "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/requestmodel"
)

type WebhookEvent struct {
	Type string `json:"type"`
	Data struct {
		Object struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"object"`
	} `json:"data"`
}

type OrderHandler struct {
	Clind pb.OrderServiceClient
}

func NewOrderHandler(clind pb.OrderServiceClient) *OrderHandler {
	return &OrderHandler{Clind: clind}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exist := c.MustGet("userID").(string)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't fetch user id from context"})
		return
	}

	var order requestmodel_order_svc_clind.OrderReq
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.Clind.OrderCreation(context.Background(), &pb.OrderRequest{
		UserID:   userID,
		OderType: order.PaymentType,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusBadRequest, result)
}

func (h *OrderHandler) Payment(c *gin.Context) {
	userID := c.Query("userID")
	orderID := c.Query("orderID")

	result, err := h.Clind.OnlinePayment(context.Background(), &pb.PaymentRequest{
		UserID:  userID,
		OrderID: orderID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.HTML(http.StatusOK, "stripe.html", gin.H{"paymentSecret": result.OrderIDSecret})
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	userID, exist := c.MustGet("userID").(string)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't fetch user id from context"})
		return
	}

	result, err := h.Clind.GetAllOrders(context.Background(), &pb.OrderShowCaseRequest{
		UserID: userID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusBadRequest, result)
}

func (h *OrderHandler) StripeWebHook(c *gin.Context) {

	var event WebhookEvent
	if err := json.NewDecoder(c.Request.Body).Decode(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode webhook event"})
		return
	}

	eventType := event.Type
	paymentID := event.Data.Object.ID
	paymentStatus := event.Data.Object.Status

	if eventType == "payment_intent.succeeded" {
		fmt.Printf("Payment succeeded. ID: %s\n", paymentID)
		result, err := h.Clind.UpdataPaymentStatus(context.Background(), &pb.UpdataPaymentStatusRequest{
			IntentPaymentID: paymentID,
		})
		if err != nil {
			fmt.Println("error at payment", paymentID, err)
			return
		}

		fmt.Println("payment succesfully", result)
	}
	// else {
	// 	fmt.Printf("Payment failed. ID: %s\n", paymentID)
	// 	// Handle failed payment (e.g., log error, notify user)
	// }

	fmt.Println("Webhook event processed successfully", paymentStatus, paymentID, "evenType :=", eventType)
}
