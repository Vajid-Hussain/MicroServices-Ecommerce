package routes_prodcuts_svc

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	handler_product_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/handler"
)

func ProductUserRoutes(engin *gin.RouterGroup, category *handler_product_svc.ProductHandler, authMiddlewire *handler.Middlewire) {
	
	engin.GET("", category.GetAllProduct)

	engin.Use(authMiddlewire.UserAuthMiddlewire)
	{
		cartManagement := engin.Group("/user")
		{
			cartManagement.POST("/cart", category.CrateCart)
			cartManagement.GET("/cart", category.GetCart)
		}
	}
}
