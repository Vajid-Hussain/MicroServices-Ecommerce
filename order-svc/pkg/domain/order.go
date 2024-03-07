package domain_order_svc

import "time"

type Order struct {
	ID            uint `gorm:"primary key"`
	UserID        uint
	PaymentMethod string
	OrderIDOnline string
}

type OrderProducts struct {
	ItemID        uint `gorm:"primarykey"`
	UserID        string
	OrderID       string
	Orderid       Order `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
	InventoryID   uint
	Quantity      uint
	Price         uint
	Discount      uint `gorm:"default:0"`
	PayableAmount uint
	OrderDate     time.Time
	EndDate       time.Time
	PaymentStatus string
	OrderStatus   string
}
