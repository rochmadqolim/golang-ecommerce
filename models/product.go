package models

type Product struct {
	ID           uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	CategoryName string `gorm:"size:100;not null" json:"category_name"`
	Name         string `gorm:"size:100;not null" json:"name"`
	Price        int    `gorm:"not null" json:"price"`
	Stock        int    `gorm:"not null" json:"stock"`
}