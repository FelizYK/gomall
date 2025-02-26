package repository

import (
	"context"

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

func GetProductById(ctx context.Context, id uint) (product Product, err error) {
	err = DB.WithContext(ctx).Model(&Product{}).
		Where("id = ?", id).First(&product).Error
	return
}

func SearchProducts(ctx context.Context, query string) (products []Product, err error) {
	err = DB.WithContext(ctx).Model(&Product{}).
		Where("name LIKE ? or description like ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return
}
