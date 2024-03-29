package interfaceUseCase_product_svc

import (
	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
)

type ICategoryUseCase interface {
	NewCategory(*requestmodel_product_svc.Category) (*responsemodel_product_svc.CategoryDetails, error)
	GetAllCategory(string, string) (*[]responsemodel_product_svc.CategoryDetails, error)
	CreateBrand(*requestmodel_product_svc.Brand) (*responsemodel_product_svc.BrandDetails, error)
	GetAllBrand(string, string) (*[]responsemodel_product_svc.BrandDetails, error)
	AddInventory( *requestmodel_product_svc.InventoryReq) (*responsemodel_product_svc.InventoryRes, error)
	GetAllInventory( string, string) (*[]responsemodel_product_svc.InventoryShowcase, error) 
	CreateCart( *requestmodel_product_svc.Cart) (*responsemodel_product_svc.Cart, error)
	ShowCart( string) (*responsemodel_product_svc.UserCart, error)
	GetCartForOrder( string) (*responsemodel_product_svc.UserCart, error)
}