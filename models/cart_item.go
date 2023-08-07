package models

type CartItem struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CartID      uint32 `gorm:"not null"`
	ProductName uint32 `gorm:"not null" json:"product_name"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	SubTotal    int    `gorm:"not null" json:"subtotal"`
}