package usecase_prodcut_svc

import (
	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
	interfaceRepositoryProductService "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/repository/interface"
	interfaceUseCase_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/usecase/interface"
	helper_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/utils"
)

type categoryUseCase struct {
	repo interfaceRepositoryProductService.ICategoryRepository
}

func NewCategoryUseCase(repository interfaceRepositoryProductService.ICategoryRepository) interfaceUseCase_product_svc.ICategoryUseCase {
	return &categoryUseCase{repo: repository}
}

func (r *categoryUseCase) NewCategory(categoryDetails *requestmodel_product_svc.Category) (*responsemodel_product_svc.CategoryDetails, error) {

	err := r.repo.InsertCategory(categoryDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
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

	err := r.repo.InsertBrand(brandDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
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