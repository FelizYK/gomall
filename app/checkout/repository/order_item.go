package repository

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32
	Quantity     int32
	Cost         float32
	OrderIdRefer uint `gorm:"index"`
}

func (u OrderItem) TableName() string {
	return "order_item"
}
