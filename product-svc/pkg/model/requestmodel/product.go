package requestmodel_product_svc

type Category struct {
	Name string `json:"name" validate:"required,alpha"`
}

type Brand struct {
	Name string `json:"name" validate:"required,alpha"`
}

type InventoryReq struct {
	Productname        string  `form:"productname" validate:"required,min=3,max=100"`
	Description        string  `form:"description" validate:"required,min=5"`
	BrandID            uint    `form:"brandID" validate:"required,number"`
	CategoryID         uint    `form:"categoryID" validate:"required,number"`
	Mrp                uint    `form:"mrp" validate:"required,min=0,number"`
	Discount           uint    `form:"discount" validate:"required,min=0,max=99,number"`
	Saleprice          uint    `form:"saleprice" swaggerignore:"true"`
	Units              uint64  `form:"units" validate:"required,min=0,number"`
	Os                 string  `form:"os" validate:"required"`
	CellularTechnology string  `form:"cellularTechnology" validate:"required"`
	Ram                uint    `form:"ram" validate:"required,min=1"`
	Screensize         float64 `form:"screensize" validate:"required,min=2"`
	Batterycapacity    uint    `form:"batterycapacity" validate:"required,min=500"`
	Processor          string  `form:"processor" validate:"required" `
	Image              []byte  `form:"image" swaggerignore:"true"`
	ImageURL           string  `swaggerignore:"true"`
}

type Cart struct {
    UserID      string `json:"cartid userid" swaggerignore:"true"`
    InventoryID string `json:"inventoryid" validate:"required,number"`
    Quantity    uint   `json:"quantity" swaggerignore:"true"`
}