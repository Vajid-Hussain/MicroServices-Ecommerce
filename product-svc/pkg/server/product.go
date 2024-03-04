package server_product_svc

import (
	"context"

	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	"github.com/vajid-hussain/mobile-mart-microservice-product/pkg/pb"
	interfaceUseCase_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/usecase/interface"
)

type ProductHandler struct {
	categoryUseCase interfaceUseCase_product_svc.ICategoryUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductHandler(useCase interfaceUseCase_product_svc.ICategoryUseCase) *ProductHandler {
	return &ProductHandler{categoryUseCase: useCase}
}

func (u *ProductHandler) CreateCatogory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCatogoryResponse, error) {
	var categoryDetails requestmodel_product_svc.Category
	categoryDetails.Name = req.Name

	result, err := u.categoryUseCase.NewCategory(&categoryDetails)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCatogoryResponse{
		Id:   result.ID,
		Name: result.Name,
	}, nil
}
