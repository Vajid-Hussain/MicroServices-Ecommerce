package server_product_svc

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	clind_prodcut_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/clind"
	handler_product_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/handler"
	routes_prodcuts_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/routes"
)

func InitProductClind(config *config.Config, engin *gin.Engine, authMiddlewire *handler.Middlewire) error {

	clind, _ := clind_prodcut_svc.InitializeClindPoductService(*config)
	categoryHandler := handler_product_svc.NewProductHandler(clind.Clind)

	routes_prodcuts_svc.ProductRoutes(engin.Group("/"), categoryHandler, authMiddlewire)
	routes_prodcuts_svc.ProductUserRoutes(engin.Group("/"), categoryHandler, authMiddlewire)

	return nil
}
