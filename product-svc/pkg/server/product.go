package server_product_svc

import (
	"context"
	"fmt"

	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	"github.com/vajid-hussain/mobile-mart-microservice-product/pkg/pb"
	interfaceUseCase_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/usecase/interface"
	helper_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/utils"
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

	validationErr := helper_product_svc.EasyValidataion(categoryDetails)
	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := u.categoryUseCase.NewCategory(&categoryDetails)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCatogoryResponse{
		Id:   result.ID,
		Name: result.Name,
	}, nil
}

func (u *ProductHandler) GetAllCategory(ctx context.Context, req *pb.GetAllCategoryRequest) (*pb.GetAllCategoryResponse, error) {
	result, err := u.categoryUseCase.GetAllCategory(req.Ofset, req.Limit)
	if err != nil {
		return nil, err
	}

	var endResult []*pb.CreateCatogoryResponse
	for _, val := range *result {
		endResult = append(endResult, &pb.CreateCatogoryResponse{Id: val.ID, Name: val.Name})
	}

	return &pb.GetAllCategoryResponse{
		Categories: endResult,
	}, nil
}

func (u *ProductHandler) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandResponse, error) {
	var brandDetails requestmodel_product_svc.Brand
	brandDetails.Name = req.Name

	validationErr := helper_product_svc.EasyValidataion(brandDetails)
	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := u.categoryUseCase.CreateBrand(&brandDetails)
	if err != nil {
		return nil, err
	}

	fmt.Println("---", result)

	return &pb.CreateBrandResponse{
		Id:   result.ID,
		Name: result.Name,
	}, nil
}

func (u *ProductHandler) GetAllBrand(ctx context.Context, req *pb.GetAllBrandRequest) (*pb.GetAllBrandResponse, error) {
	result, err := u.categoryUseCase.GetAllBrand(req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	var endResult []*pb.CreateBrandResponse
	for _, val := range *result {
		endResult = append(endResult, &pb.CreateBrandResponse{Id: val.ID, Name: val.Name})
	}

	return &pb.GetAllBrandResponse{
		Brands: endResult,
	}, nil
}
