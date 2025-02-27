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

func GetProducts(ctx context.Context) (products []Product, err error) {
	err = DB.WithContext(ctx).Model(&Product{}).
		Find(&products).Error
	return
}

func GetProductById(ctx context.Context, id uint32) (product Product, err error) {
	err = DB.WithContext(ctx).Model(&Product{}).
		Where("id = ?", id).First(&product).Error
	return
}

func SearchProducts(ctx context.Context, query string) (products []Product, err error) {
	err = DB.WithContext(ctx).Model(&Product{}).
		Where("name LIKE ? or description like ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return
}
