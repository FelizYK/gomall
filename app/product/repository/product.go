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

type ProductQuery struct {
	ctx context.Context
}

func NewProductQuery(ctx context.Context) ProductQuery {
	return ProductQuery{ctx: ctx}
}

func (q ProductQuery) GetProducts() (products []Product, err error) {
	err = db.WithContext(q.ctx).Model(&Product{}).
		Find(&products).Error
	return
}

func (q ProductQuery) GetProductById(id uint32) (product Product, err error) {
	err = db.WithContext(q.ctx).Model(&Product{}).
		Where("id = ?", id).First(&product).Error
	return
}

func (q ProductQuery) SearchProducts(query string) (products []Product, err error) {
	err = db.WithContext(q.ctx).Model(&Product{}).
		Where("name LIKE ? or description like ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return
}
