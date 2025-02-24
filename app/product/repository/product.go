package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Picture     string
	Price       float32

	Categories []Category `gorm:"many2many:product_category"`
}

func (u Product) TableName() string {
	return "product"
}
