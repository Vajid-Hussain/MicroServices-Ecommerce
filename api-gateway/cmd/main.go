package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	server_auth "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/server"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	engin := gin.Default()

	err = server_auth.InitGinServer(*config, engin)
	if err != nil {
		log.Fatal(err)
	}

	err = engin.Run(config.Port)
	if err != nil {
		fmt.Println("err:", err)
	}
}
