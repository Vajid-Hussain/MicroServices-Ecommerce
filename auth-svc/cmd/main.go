package main

import (
	"log"
	"net"

	config_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/config"
	di_auth_svc "github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/di"
	"github.com/vajid-hussain/mobile-mart-microservice-auth/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config_auth_svc.InitConfig()
	if err != nil {
		log.Fatalln("error at laoding env", err)
	}

	lis, err := net.Listen("tcp", config.General.Port)
	if err != nil {
		log.Fatalln("cannot make a tcp socket: ", err)
	}

	server := di_auth_svc.InitializeServer(*config)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}
