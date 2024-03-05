package usecase_prodcut_svc

import (
	config_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/config"
	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
	interfaceRepositoryProductService "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/repository/interface"
	interfaceUseCase_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/usecase/interface"
	helper_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/utils"
)

type categoryUseCase struct {
	repo interfaceRepositoryProductService.ICategoryRepository
	config *config_product_svc.Config
}

func NewCategoryUseCase(repository interfaceRepositoryProductService.ICategoryRepository,config *config_product_svc.Config) interfaceUseCase_product_svc.ICategoryUseCase {
	return &categoryUseCase{repo: repository,
	config: config,}
}

func (r *categoryUseCase) NewCategory(categoryDetails *requestmodel_product_svc.Category) (*responsemodel_product_svc.CategoryDetails, error) {

	category, err := r.repo.InsertCategory(categoryDetails)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryUseCase) GetAllCategory(page string, limit string) (*[]responsemodel_product_svc.CategoryDetails, error) {

	offSet, limits, err := helper_product_svc.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	categoryDetails, err := r.repo.GetAllCategory(offSet, limits)
	if err != nil {
		return nil, err
	}

	return categoryDetails, nil
}

//brand

func (r *categoryUseCase) CreateBrand(brandDetails *requestmodel_product_svc.Brand) (*responsemodel_product_svc.BrandDetails, error) {

	brand, err := r.repo.InsertBrand(brandDetails)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (r *categoryUseCase) GetAllBrand(page string, limit string) (*[]responsemodel_product_svc.BrandDetails, error) {

	offSet, limits, err := helper_product_svc.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	brandDetails, err := r.repo.GetAllBrand(offSet, limits)
	if err != nil {
		return nil, err
	}

	return brandDetails, nil
}

func (d *categoryUseCase) AddInventory(inventory *requestmodel_product_svc.InventoryReq) (*responsemodel_product_svc.InventoryRes, error) {

	sess := helper_product_svc.CreateSession(d.config)

	ImageURL, err := helper_product_svc.UploadImageToS3(inventory.Image, sess)
	if err != nil {
		return nil, err
	}

	inventory.ImageURL = ImageURL
	discountedPrice := helper_product_svc.FindDiscount(float64(inventory.Mrp), float64(inventory.Discount))
	inventory.Saleprice = discountedPrice

	product, err := d.repo.CreateProduct(inventory)
	if err != nil {
		return nil, err
	}

	return product, nil
}
