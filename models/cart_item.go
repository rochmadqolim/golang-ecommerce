package models

type CartItem struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CartID      uint32 `gorm:"not null" json:"cart_id"`
	ProductName string `gorm:"not null" json:"product_name"`
	Price       int    `gorm:"not null" json:"price"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	SubTotal    int    `gorm:"not null"`
}