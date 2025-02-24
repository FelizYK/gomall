package repository

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string

	Products []Product `gorm:"many2many:product_category"`
}

func (u Category) TableName() string {
	return "category"
}
