package models

type Customer struct {
	ID       uint32 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"size:100;index" json:"fullname" binding:"required"`
	Email    string `gorm:"size:100;unique;not null" json:"email" binding:"required,email"`
	Password string `gorm:"size:100;not null" json:"password" binding:"required"`
}
