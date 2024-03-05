package responsemodel_product_svc


type CategoryDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type BrandDetails struct {
	ID   string `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}

type ValidatonError struct {
    Error string
}

type InventoryRes struct {
    ID                 uint    `json:"id"`
    Productname        string  `json:"productname" validate:"required,min=3,max=100"`
    Description        string  `json:"description" validate:"required,min=5"`
    BrandID            uint    `json:"brandID" validate:"required"`
    CategoryID         uint    `json:"categoryID" validate:"required"`
    Mrp                uint    `json:"mrp" validate:"required,min=0"`
    Discount           uint    `form:"discount" validate:"required,min=0,max=99,number"`
    Saleprice          uint    `json:"saleprice" validate:"required,min=0"`
    CategoryDiscount   uint    `json:"categoryDiscount,omitempty"`
    NetDiscount        uint    `json:"netDiscount,omitempty"`
    FinalPrice         uint    `json:"priceApplyCategoryDiscount,omitempty"`
    Units              uint64  `json:"units" validate:"required,min=0"`
    Os                 string  `json:"os"`
    CellularTechnology string  `json:"cellularTechnology"`
    Ram                uint    `json:"ram" validate:"required,min=0"`
    Screensize         float64 `json:"screensize" validate:"required,min=0"`
    Batterycapacity    uint    `json:"batterycapacity" validate:"required,min=0"`
    Processor          string  `json:"processor" validate:"required"`
    ImageURL           string  `json:"imageURL" validate:"required"`
}

type InventoryShowcase struct {
    ID                              uint   `json:"productID"`
    Productname                     string `json:"productname"`
    Mrp                             int    `json:"mrp" `
    Discount                        uint   `form:"discount" `
    Saleprice                       int    `json:"saleprice" `
    CategoryDiscount                uint   `json:"categoryDiscount,omitempty"`
    Units                           uint   `json:"units"`
}

type Cart struct {
    UserID      string `json:"cartid userid" swaggerignore:"true"`
    InventoryID string `json:"inventoryid" validate:"required,number"`
    Quantity    uint   `json:"quantity" swaggerignore:"true"`
    Price      uint   `json:"price,omitempty" swaggerignore:"true"`
}
