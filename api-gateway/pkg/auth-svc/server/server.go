package server_auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/clind"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	routes_auth "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/routes"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

func InitGinServer(config config.Config, engin *gin.Engine) error {

	clind, err := clind.InitServiceClind(&config)
	if err != nil {
		log.Fatalln("Dial port is not not correct cross check")
	}

	authHandler := handler.NewUserHandler(clind.Clind, engin)

	routes_auth.RouteAuth(engin, authHandler)

	return nil
}
