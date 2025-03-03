package repository

import (
	"context"

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

func GetCategoryByName(ctx context.Context, name string) (category Category, err error) {
	err = db.WithContext(ctx).Model(&Category{}).
		Where("name = ?", name).Preload("Products").First(&category).Error
	return
}
