package clind

import (
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/auth-svc/pb"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClind struct {
	Clind pb.AuthServiceClient
}

func InitServiceClind(url *config.Config) (*AuthClind, error) {
	cc, err := grpc.Dial(url.AuthPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := AuthClind{
		Clind: pb.NewAuthServiceClient(cc),
	}
	return &c, nil
}
