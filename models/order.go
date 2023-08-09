package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Fullname  string `gorm:"size:100;not null" json:"fullname"`
	Address  string `gorm:"size:100;not null" json:"address"`
	Cart Cart
	CartID uint32 `gorm:"not null" json:"cart_id"`
	TotalAmount int    `gorm:"not null" json:"total_amount"`
}