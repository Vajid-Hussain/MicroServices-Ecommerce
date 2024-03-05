package domain_products_svc

type Category struct {
	ID     uint `gorm:"unique key"`
	Name   string
	Status string `gorm:"default:active"`
}

type Brand struct {
	ID     uint   `gorm:"unique key"`
	Name   string `gorm:"unique"`
	Status string `gorm:"default:active"`
}

type Inventories struct {
	ID                 uint `gorm:"primary key"`
	Productname        string
	Description        string
	BrandID            uint     `gorm:"not null"`
	Brand              Brand    `gorm:"forgienkey:BrandID;association_foreignkey:ID"`
	CategoryID         uint     `gorm:"not null"`
	Category           Category `gorm:"forgienKey:CategoryID;association_foreignkey:ID"`
	Mrp                int
	Discount           uint
	Saleprice          int
	Units              int64
	Os                 string
	CellularTechnology string
	Ram                int
	Screensize         float64
	Batterycapacity    int
	Processor          string
	ImageURL           string
	Status             string `gorm:"default:active"`
}

type Cart struct {
	UserID      uint
	InventoryID uint
	Product     Inventories `gorm:"foreignkey:InventoryID;association_foreignkey:ID"`
	Quantity    uint        `gorm:"default:1"`
	Status      string      `gorm:"default:active"`
}
