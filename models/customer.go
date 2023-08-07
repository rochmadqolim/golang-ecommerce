package models

import "time"

type Customer struct {
	ID        uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Fullname  string `gorm:"size:100;not null" json:"fullname"`
	Email     string `gorm:"size:100;unique;not null" json:"email"`
	Password  string `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
