package requestmodel_order_svc

// type Order struct{
// 	UserID string
// 	ID string
// 	OrderIDOnline string
// 	PaymentMethod  string
// }

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

type UserCart struct {
	UserID         string `json:"user_id" validate:"" gorm:"-"`
	TotalPrice     uint   `json:"total_price"`
	InventoryCount uint   `json:"product_count"`
	OrderID        string
	PaymentID      string
	IntentUniquePaymentID string
	PaymentMethod  string
	OrderStatus    string
	PaymentStatus  string
	Cart           []CartInventory `json:"cart" gorm:"-"`
}
