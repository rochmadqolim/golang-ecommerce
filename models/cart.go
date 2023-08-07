package models

import (
	"time"
)

type Cart struct {
	ID       uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Customers       Customer   `gorm:"foreignkey:CustomerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CustomerID uint32 `gorm:"not null"`
	CartItems   []CartItem
	Price int `gorm:"not null" json:"price"`
	UpdatedAt  time.Time
	

	
}