package handler_order_svc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/pb"
	requestmodel_order_svc_clind "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/requestmodel"
)

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
	c.HTML(http.StatusOK, "stripe.html", "")
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
