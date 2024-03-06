package interfaceRepositoryProductService

import (
	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
)

type ICategoryRepository interface {
	InsertCategory(*requestmodel_product_svc.Category)  (*responsemodel_product_svc.CategoryDetails, error)
	GetAllCategory(int, int) (*[]responsemodel_product_svc.CategoryDetails, error)
	InsertBrand(*requestmodel_product_svc.Brand) (* responsemodel_product_svc.BrandDetails,error)
	GetAllBrand(int, int) (*[]responsemodel_product_svc.BrandDetails, error)
	CreateProduct( *requestmodel_product_svc.InventoryReq) (*responsemodel_product_svc.InventoryRes, error) 
	GetInventory( int, int) (*[]responsemodel_product_svc.InventoryShowcase, error)
	IsInventoryExistInCart( string,  string) (int, error) 
	InsertToCart( *requestmodel_product_svc.Cart) (*responsemodel_product_svc.Cart, error)
	GetCartCriteria( string) (uint, error)
	GetCart( string) (*[]responsemodel_product_svc.CartInventory, error)
	DeleteInventoryFromCart( string, string) error 
}