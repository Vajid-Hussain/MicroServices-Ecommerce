package main

import (
	"log"

	"github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	db_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/db"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln("error at laoding env", err)
	}

	db, err := db_auth_svc.InitDB(&config.DB)
	if err!=nil{
		log.Fatal(err)
	}
	
}
