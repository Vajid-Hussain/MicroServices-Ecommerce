package repository_poduct_svc

import (
	"errors"
	"fmt"

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

func (d *categoryRepository) InsertCategory(categoryDetails *requestmodel_product_svc.Category) (*responsemodel_product_svc.CategoryDetails, error) {
	var category responsemodel_product_svc.CategoryDetails
	query := "INSERT INTO categories(name) VALUES(?) RETURNING *"
	err := d.DB.Raw(query, categoryDetails.Name).Scan(&category).Error
	if err != nil {
		return nil, errors.New("canot make a new Category")
	}
	return &category, nil
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
func (d *categoryRepository) InsertBrand(name *requestmodel_product_svc.Brand) (*responsemodel_product_svc.BrandDetails, error) {
	fmt.Println("--", name)
	var brands responsemodel_product_svc.BrandDetails
	query := "INSERT INTO brands(name) VALUES(?) RETURNING *"
	err := d.DB.Raw(query, name.Name).Scan(&brands).Error
	if err != nil {
		return nil, errors.New("can't make a new brand alrady exit, come with new one")
	}
	return &brands, nil
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

func (d *categoryRepository) CreateProduct(inventory *requestmodel_product_svc.InventoryReq) (*responsemodel_product_svc.InventoryRes, error) {

	var insertedData responsemodel_product_svc.InventoryRes

	query := `INSERT INTO inventories (productname, description, brand_id, category_id, mrp, saleprice, units,os, cellular_technology, ram, screensize, Batterycapacity, processor, image_url, discount) VALUES (?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;`
	result := d.DB.Raw(query,
		inventory.Productname, inventory.Description, inventory.BrandID, inventory.CategoryID,
		inventory.Mrp, inventory.Saleprice, inventory.Units,
		inventory.Os, inventory.CellularTechnology, inventory.Ram,
		inventory.Screensize, inventory.Batterycapacity, inventory.Processor, inventory.ImageURL, inventory.Discount,
	).Scan(&insertedData)

	if result.Error != nil {
		return nil, errors.New("inventory is not added into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("inventory is not added in database , face some error")
	}
	return &insertedData, nil
}