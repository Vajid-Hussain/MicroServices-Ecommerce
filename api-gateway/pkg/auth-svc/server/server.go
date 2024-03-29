package server_auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/clind"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	routes_auth "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/routes"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

func InitAuthClind(config config.Config, engin *gin.Engine) (*handler.Middlewire,error) {

	clind, err := clind.InitServiceClind(&config)
	if err != nil {
		log.Fatalln("Dial port is not correct cross check")
	}

	authHandler := handler.NewUserHandler(clind.Clind, engin)
	authAdminHandler := handler.NewAdminHandler(clind.Clind, engin)

	authMiddlewire := handler.NewAuthMiddlewire(clind.Clind)

	routes_auth.RouteAuthAdmin(engin, authAdminHandler, authMiddlewire)
	routes_auth.RouteAuthUser(engin, authHandler, authMiddlewire)

	return authMiddlewire, nil
}
