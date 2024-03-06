package handler_order_svc

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/pb"
)

type OrderHandler struct {
	Clind pb.OrderServiceClient
}

func NewOrderHandler (clind pb.OrderServiceClient) *OrderHandler{
	return &OrderHandler{Clind: clind}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exist := c.MustGet("userID").(string)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't fetch user id from context"})
		return
	}

	result, err := h.Clind.OrderCreation(context.Background(), &pb.OrderRequest{
		UserID: userID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusBadRequest, result)
}
