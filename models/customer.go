package models

import "time"

type Customer struct {
	ID        uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Fullname  string `gorm:"size:100;not null" json:"fullname" bindding:"required"`
	Email     string `gorm:"size:100;unique" json:"email" bindding:"required,email"`
	Password  string `gorm:"size:100" json:"password" bindding:"required"`
	CreatedAt time.Time
}
