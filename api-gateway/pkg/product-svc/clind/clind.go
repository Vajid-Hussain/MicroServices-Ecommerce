package clind_prodcut_svc

import (
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/config"
	"github.com/vajid-hussain/mobile-mart-microservice/pkg/product-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClind struct{
	Clind pb.ProductServiceClient
}

func InitializeClindPoductService(config config.Config) (*ProductClind,error){
	cc,err:= grpc.Dial(config.ProductPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		return nil, err
	}

	c:= ProductClind{Clind: pb.NewProductServiceClient(cc)}

	return &c,nil
}

