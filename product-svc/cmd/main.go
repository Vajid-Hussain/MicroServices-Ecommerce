package main

import (
	"log"
	"net"

	config_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/config"
	di_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/di"
	"github.com/vajid-hussain/mobile-mart-microservice-product/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config_product_svc.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := di_product_svc.InitializeServer(config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal(err)
	}

	grpcService := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcService, server)

	err = grpcService.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
