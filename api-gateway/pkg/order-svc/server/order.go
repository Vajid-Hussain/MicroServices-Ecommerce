package server_order_clind

import (
	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	order_clind "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/clind"
	handler_order_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/handler"
	routes_order "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/routes"
)

func InitOrderClind(config config.Config, engin *gin.Engine, authMiddlewire *handler.Middlewire) error {
	clind, err := order_clind.InitOrderClind(config)
	if err != nil {
		return err
	}

	engin.LoadHTMLGlob("./template/*.html")

	orderHandler := handler_order_svc.NewOrderHandler(clind.Clind)
	routes_order.OrderRoutes(engin, *authMiddlewire, *orderHandler)

	return nil
}
