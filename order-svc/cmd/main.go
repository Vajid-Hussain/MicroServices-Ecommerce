package main

import (
	"log"
	"net"

	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
	di_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/di"
	"github.com/vajid-hussain/mobile-mart-microservice-order/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config_order_svc.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	serverFunc, err := di_order_svc.InitializeOrderServer(config)
	if err != nil {
		log.Fatal(err)
	}

	lis,err:= net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer:=grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, serverFunc)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

