package repository

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email     string
	Firstname string
	Lastname  string
	Street    string
	City      string
	Province  string
	Country   string
}

type CreditCard struct {
	CardNum         string
	ExpirationMonth int32
	ExpirationYear  int32
	Cvv             int32
}

type Order struct {
	gorm.Model
	UserId     uint32
	Consignee  Consignee  `gorm:"embedded"`
	CreditCard CreditCard `gorm:"embedded"`
	TotalCost  float32
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:ID"`
}

func (u Order) TableName() string {
	return "order"
}

func GetOrders(ctx context.Context, userId uint32) (orders []*Order, err error) {
	err = DB.WithContext(ctx).Model(&Order{}).
		Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	return
}

func AddOrder(ctx context.Context, order *Order) (err error) {
	err = DB.WithContext(ctx).Model(&Order{}).
		Create(order).Error
	return
}
