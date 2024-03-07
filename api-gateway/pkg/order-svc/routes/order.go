package routes_order

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	handler_order_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/handler"
)

func OrderRoutes(engin *gin.Engine, authMiddlewire handler.Middlewire, order handler_order_svc.OrderHandler) {
	engin.POST("/webhook", order.StripeWebHook)
	
	orderManagement := engin.Group("/order")
	{
		orderManagement.GET("/", order.Payment)

		orderManagement.Use(authMiddlewire.UserAuthMiddlewire)
		{
			orderManagement.GET("/all",order.GetAllOrders)
			orderManagement.POST("/", order.CreateOrder)
		}
	}
}
