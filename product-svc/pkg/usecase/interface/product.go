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
}