package models

type Category struct {
	ID   uint32 `gorm:"primaryKey;auto_increment" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
}
