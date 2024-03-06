package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	server_auth "github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/server"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	server_order_clind "github.com/vajid-hussain/mobile-mart-microservice/pkg/order-svc/server"
	server_product_svc "github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/server"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	engin := gin.Default()

	authMiddlewire, err := server_auth.InitAuthClind(*config, engin)
	if err != nil {
		log.Fatal(err)
	}

	err = server_product_svc.InitProductClind(config, engin, authMiddlewire)
	if err != nil {
		log.Fatal(err)
	}

	err = server_order_clind.InitOrderClind(*config, engin, authMiddlewire)
	if err != nil {
		log.Fatal(err)
	}


	err = engin.Run(config.Port)
	if err != nil {
		fmt.Println("err:", err)
	}
}
