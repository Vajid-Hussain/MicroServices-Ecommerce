package repository_poduct_svc

import (
	"errors"
	"fmt"

	requestmodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/requestmodel"
	responsemodel_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel"
	resCustomError_product_svc "github.com/vajid-hussain/mobile-mart-microservice-product/pkg/model/responsemodel/custom_error"
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

func (d *categoryRepository) GetInventory(offSet int, limit int) (*[]responsemodel_product_svc.InventoryShowcase, error) {
	var inventory []responsemodel_product_svc.InventoryShowcase

	query := "SELECT * FROM inventories ORDER BY inventories.id OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&inventory).Error
	if err != nil {
		return nil, errors.New("can't get inventory data from db")
	}

	return &inventory, nil
}

//cart

func (d *categoryRepository) IsInventoryExistInCart(inventoryID string, userID string) (int, error) {
	var inventoryCount int

	query := "SELECT count(*) FROM carts WHERE inventory_id=? AND user_id=? AND status='active' "
	result := d.DB.Raw(query, inventoryID, userID).Scan(&inventoryCount)
	if result.Error != nil {
		return 0, errors.New("face some issue while finding inventory is exist in cart")
	}
	return inventoryCount, nil
}

func (d *categoryRepository) InsertToCart(cart *requestmodel_product_svc.Cart) (*responsemodel_product_svc.Cart, error) {
	var insertCart *responsemodel_product_svc.Cart
	query := "INSERT INTO carts (user_id, inventory_id, quantity) VALUES (?, ?,  ?)   RETURNING *;"
	result := d.DB.Raw(query, cart.UserID, cart.InventoryID, cart.Quantity).Scan(&insertCart)

	if result.Error != nil {
		return nil, errors.New("face some issue while inventory insert to cart ")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError_product_svc.ErrNoRowAffected
	}
	return insertCart, nil
}

func (d *categoryRepository) GetCart(userID string) (*[]responsemodel_product_svc.CartInventory, error) {
	var cartView *[]responsemodel_product_svc.CartInventory
	query := "SELECT * FROM carts INNER JOIN inventories ON id=inventory_id  WHERE carts.user_id=? AND carts.status='active'"
	result := d.DB.Raw(query, userID).Scan(&cartView)
	if result.Error != nil {
		return nil, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user have no cart")
	}
	return cartView, nil
}

func (d *categoryRepository) GetCartCriteria(userID string) (uint, error) {

	var count uint
	query := "SELECT SUM(quantity) FROM carts WHERE user_id=? AND status='active'"
	result := d.DB.Raw(query, userID)
	result.Row().Scan(&count)
	if result.Error != nil {
		return 0, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError_product_svc.ErrNoRowAffected
	}
	return count, nil
}

func (d *categoryRepository) DeleteInventoryFromCart(inventoryID string, userID string) error {

	query := "UPDATE carts SET status='delete' WHERE inventory_id = ? AND user_id= ? AND status= 'active'"
	result := d.DB.Exec(query, inventoryID, userID)
	if result.Error != nil {
		return errors.New("face some issue while delete inventory in cart")
	}
	if result.RowsAffected == 0 {
		return resCustomError_product_svc.ErrNoRowAffected
	}
	return nil
}

func (d *categoryRepository) GetInventoryUnits(inventoryID string) (*uint, error) {

	var units uint
	query := "SELECT units FROM inventories WHERE id=?"
	result := d.DB.Raw(query, inventoryID).Scan(&units)
	if result.Error != nil {
		return nil, errors.New("face some issue while get inventory units")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError_product_svc.ErrNoRowAffected
	}
	return &units, nil
}

func (d *categoryRepository) UpdateInventoryUnits(inventoryID string, units uint) error {

	query := "UPDATE inventories SET units= ? WHERE id =?"
	result := d.DB.Exec(query, units, inventoryID)
	if result.Error != nil {
		return errors.New("face some issue while updating inventory unit")
	}
	if result.RowsAffected == 0 {

		return resCustomError_product_svc.ErrNoRowAffected
	}
	return nil
}
