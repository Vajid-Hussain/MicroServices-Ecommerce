package repository_poduct_svc

import (
	"errors"

	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
	interfaceRepositoryProductService "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/repository/interface"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaceRepositoryProductService.ICategoryRepository {
	return &categoryRepository{DB: db}
}

func (d *categoryRepository) InsertCategory(categoryDetails *requestmodel_product_svc.Category) error {
	query := "INSERT INTO categories(name) VALUES(?)"
	err := d.DB.Exec(query, categoryDetails.Name).Error
	if err != nil {
		return errors.New("canot make a new Category")
	}
	return nil
}

func (d *categoryRepository) GetAllCategory(offSet int, limit int) (*[]responsemodel_product_svc.CategoryDetails, error) {
	var categories []responsemodel_product_svc.CategoryDetails

	query := "SELECT * FROM categories WHERE status!='delete' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&categories).Error
	if err != nil {
		return nil, errors.New("can't fetch categories from database")
	}

	return &categories, nil
}

// Brand
func (d *categoryRepository) InsertBrand(name *requestmodel_product_svc.Brand) error {
	query := "INSERT INTO brands(name) VALUES(?)"
	err := d.DB.Exec(query, name.Name).Error
	if err != nil {
		return errors.New("can't make a new brand")
	}
	return nil
}

func (d *categoryRepository) GetAllBrand(offSet int, limit int) (*[]responsemodel_product_svc.BrandDetails, error) {
	var Brands []responsemodel_product_svc.BrandDetails

	query := "SELECT * FROM brands WHERE status!='delete' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&Brands).Error
	if err != nil {
		return nil, errors.New("can't fetch brand from database")
	}

	return &Brands, nil
}
