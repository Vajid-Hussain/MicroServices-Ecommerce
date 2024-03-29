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

func (u *ProductHandler) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	var product requestmodel_product_svc.InventoryReq

	product.Batterycapacity = uint(req.BatteryCapacity)
	product.BrandID = uint(req.BrandId)
	product.CategoryID = uint(req.CategoryId)
	product.CellularTechnology = req.CellularTechnology
	product.Description = req.Description
	product.Discount = uint(req.Discount)
	product.Image = req.Image
	product.Mrp = uint(req.Mrp)
	product.Os = req.Os
	product.Processor = req.Processor
	product.Productname = req.Productname
	product.Ram = uint(req.Ram)
	product.Screensize = float64(req.Screensize)
	product.Units = uint64(req.Units)
	product.Saleprice = uint(req.Saleprice)

	validationErr := helper_product_svc.EasyValidataion(product)
	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := u.categoryUseCase.AddInventory(&product)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponse{
		Id:                 uint32(result.ID),
		Productname:        result.Productname,
		Description:        product.Description,
		BrandId:            req.BrandId,
		CategoryId:         req.CategoryId,
		Mrp:                uint32(result.Mrp),
		Discount:           uint32(result.Discount),
		Saleprice:          uint32(result.Saleprice),
		CategoryDiscount:   uint32(result.CategoryDiscount),
		NetDiscount:        uint32(result.NetDiscount),
		FinalPrice:         uint32(result.FinalPrice),
		Units:              result.Units,
		Os:                 result.Os,
		CellularTechnology: result.CellularTechnology,
		Ram:                uint32(result.Ram),
		Screensize:         result.Screensize,
		BatteryCapacity:    uint32(result.Batterycapacity),
		Processor:          product.Processor,
		ImageUrl:           result.ImageURL,
	}, nil
}

func (h *ProductHandler) GetAllProduct(ctx context.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductResponseSlice, error) {
	result, err := h.categoryUseCase.GetAllInventory(req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.GetAllProductResponse
	for _, val := range *result {
		finalResult = append(finalResult, &pb.GetAllProductResponse{ProductId: uint32(val.ID), Productname: val.Productname, Mrp: int32(val.Mrp), Discount: uint32(val.Discount), Saleprice: int32(val.Saleprice), CategoryDiscount: uint32(val.CategoryDiscount), Units: uint32(val.Units)})
	}

	return &pb.GetAllProductResponseSlice{
		Products: finalResult,
	}, nil
}

func (h *ProductHandler) CreateCart(ctx context.Context, req *pb.CreateCartRequst) (*pb.CreateCartResponse, error) {
	var cart = &requestmodel_product_svc.Cart{}
	cart.InventoryID = req.InventoryId
	cart.Quantity = uint(req.Quantity)
	cart.UserID = req.UserId

	validationErr := helper_product_svc.EasyValidataion(cart)
	var errorString string
	if len(validationErr) > 0 {
		for _, val := range validationErr {
			errorString += fmt.Errorf(val.Error).Error()
		}
	}
	if len(validationErr) > 0 {
		return nil, fmt.Errorf(errorString)
	}

	result, err := h.categoryUseCase.CreateCart(cart)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCartResponse{
		UserId:      result.UserID,
		InventoryId: result.InventoryID,
		Quantity:    uint32(result.Quantity),
	}, nil
}

func (h *ProductHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	result, err := h.categoryUseCase.ShowCart(req.UserID)
	if err != nil {
		return nil, err
	}

	var allCart []*pb.CartInventory

	for _, val := range result.Cart {
		cart := &pb.CartInventory{
			Productname: val.Productname,
			InventoryId: val.InventoryID,
			CategoryId:  val.CategoryID,
			Quantity:    uint32(val.Quantity),
			Price:       uint32(val.Price),
			Discount:    uint32(val.Discount),
			Saleprice:   uint32(val.Saleprice),
			FinalPrice:  uint32(val.FinalPrice),
			Title:       val.Title,
			Units:       val.Units,
		}
		allCart = append(allCart, cart)
	}

	return &pb.GetCartResponse{UserId: result.UserID,
		TotalPrice: uint32(result.TotalPrice),
		InventoryCount: uint32(result.InventoryCount),
		Cart: allCart,}, nil
}

func (h *ProductHandler) FetchCartForOrder(ctx context.Context, req *pb.GetCartForOrderRequest) (*pb.GetCartForOrderResponse, error) {
	result, err := h.categoryUseCase.GetCartForOrder(req.UserID)
	if err != nil {
		return nil, err
	}

	var allCart []*pb.CartInventory

	for _, val := range result.Cart {
		cart := &pb.CartInventory{
			Productname: val.Productname,
			InventoryId: val.InventoryID,
			CategoryId:  val.CategoryID,
			Quantity:    uint32(val.Quantity),
			Price:       uint32(val.Price),
			Discount:    uint32(val.Discount),
			Saleprice:   uint32(val.Saleprice),
			FinalPrice:  uint32(val.FinalPrice),
			Title:       val.Title,
			Units:       val.Units,
		}
		allCart = append(allCart, cart)
	}

	return &pb.GetCartForOrderResponse{UserId: result.UserID,
		TotalPrice: uint32(result.TotalPrice),
		InventoryCount: uint32(result.InventoryCount),
		Cart: allCart,}, nil
}
