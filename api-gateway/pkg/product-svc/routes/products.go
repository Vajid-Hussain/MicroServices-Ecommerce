package routes_prodcuts_svc

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	handler_product_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/handler"
)

func ProductRoutes(engin *gin.Engine, category *handler_product_svc.ProductHandler, authMiddlewire *handler.Middlewire) {
	engin.Use(authMiddlewire.AdminAuthMiddlewire)
	{
		categoryManagementent := engin.Group("/category")
		{
			categoryManagementent.POST("/", category.CreateCategory)
			categoryManagementent.GET("", category.GetAllCategory)
		}

		brandManagementent := engin.Group("/brand")
		{
			brandManagementent.POST("/", category.CreateBrand)
			brandManagementent.GET("", category.GetAllBrand)
		}

		productManagement:= engin.Group("product")
		{
			productManagement.POST("/",category.AddProduct)
		}
	}
}
