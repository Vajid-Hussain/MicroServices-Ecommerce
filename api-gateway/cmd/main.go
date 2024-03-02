package main

import (
	"log"

	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/clind"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/handler"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/server"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	clind, err := clind.InitServiceClind(config)
	if err != nil {
		log.Fatalln("Dial port is not not correct cross check")
	}

	engin, err := server.InitGinServer(*config)
	if err != nil {
		log.Fatal(err)
	}

	handler.NewUserHandler(clind.Clind,engin)

}
