package models

type Cart struct {
	ID          uint32   `gorm:"primary_key:auto_increment" json:"id"`
	CustomerID  uint32   `gorm:"unique;not null"`
	Customer    Customer `gorm:"foreignKey:CustomerID"`
	CartItems   []CartItem
	TotalAmount int `gorm:"not null"`
}
