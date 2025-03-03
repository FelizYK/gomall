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

type OrderQuery struct {
	ctx context.Context
}

func NewOrderQuery(ctx context.Context) OrderQuery {
	return OrderQuery{ctx: ctx}
}

func (q OrderQuery) GetOrders(userId uint32) (orders []*Order, err error) {
	err = db.WithContext(q.ctx).Model(&Order{}).
		Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	return
}

func (q OrderQuery) AddOrder(order *Order) (err error) {
	err = db.WithContext(q.ctx).Model(&Order{}).
		Create(order).Error
	return
}
