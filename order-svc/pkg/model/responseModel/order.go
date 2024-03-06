package responsemodel_order_svc

type UserCart struct {
    UserID         string `json:"user_id"`
    TotalPrice     uint   `json:"total_price"`
    InventoryCount uint   `json:"product_count"`
    OrderID        string `gorm:"column:id"`
    PaymentID      string `gorm:"column:order_id_online"`
    PaymentMethod  string  `gorm:"column:payment_method"`
    Cart           []CartInventory `json:"cart" gorm:"-"`
}

type CartInventory struct {
	Productname      string `json:"productname" validate:"required,min=3,max=100"`
	InventoryID      string `json:"productid" validate:"required,number"`
	CategoryID       string `json:"categoryID"`
	Quantity         uint   `json:"quantity"`
	Price            uint   `json:"mrp" gorm:"column:mrp"`
	Discount         uint   `json:"productDiscount"`
	Saleprice        uint   `json:"saleprice" validate:"required,min=0,number"`
	CategoryDiscount uint   `json:"categoryDiscount,omitempty"`
	FinalPrice       uint   `json:"payedAmount"`
	Title            string `json:"categoryDiscountTitle,omitempty"`
	Units            uint64 `json:"available units" validate:"required,min=0,number"`
}
