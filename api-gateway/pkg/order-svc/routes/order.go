package routes_order

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	handler_order_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/handler"
)

func OrderRoutes(engin *gin.Engine, authMiddlewire handler.Middlewire, order handler_order_svc.OrderHandler) {
	orderManagement := engin.Group("/order")
	{
		orderManagement.Use(authMiddlewire.UserAuthMiddlewire)
		{
			engin.POST("/", order.CreateOrder)
		}
	}
}