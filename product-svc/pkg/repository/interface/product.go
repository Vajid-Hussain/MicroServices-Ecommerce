package interfaceRepositoryProductService

import (
	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
)

type ICategoryRepository interface {
	InsertCategory(*requestmodel_product_svc.Category) error
	GetAllCategory(int, int) (*[]responsemodel_product_svc.CategoryDetails, error)
	InsertBrand(*requestmodel_product_svc.Brand) error
	GetAllBrand(int, int) (*[]responsemodel_product_svc.BrandDetails, error)
}