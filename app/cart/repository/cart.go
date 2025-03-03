package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserId    uint32
	ProductId uint32
	Quantity  int32
}

func (u CartItem) TableName() string {
	return "CartItem"
}

type CartQuery struct {
	ctx context.Context
}

func NewCartQuery(ctx context.Context) CartQuery {
	return CartQuery{ctx: ctx}
}

func (q CartQuery) GetCartByUserId(userId uint32) (items []CartItem, err error) {
	err = db.WithContext(q.ctx).Model(&CartItem{}).
		Where("user_id = ?", userId).Find(&items).Error
	return
}

func (q CartQuery) AddCart(userId uint32, productId uint32, quantity int32) (err error) {
	var find CartItem
	err = db.WithContext(q.ctx).Model(&CartItem{}).
		Where("user_id = ? AND product_id = ?", userId, productId).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if find.Model.ID == 0 {
		// create
		err = db.WithContext(q.ctx).Model(&CartItem{}).
			Create(&CartItem{
				UserId:    userId,
				ProductId: productId,
				Quantity:  quantity,
			}).Error
	} else {
		// update
		err = db.WithContext(q.ctx).Model(&CartItem{}).
			Where("user_id = ? AND product_id = ?", userId, productId).
			Update("quantity", find.Quantity+quantity).Error
	}
	return
}

func (q CartQuery) EmptyCart(userId uint32) (err error) {
	err = db.WithContext(q.ctx).Model(&CartItem{}).
		Where("user_id = ?", userId).Delete(&CartItem{}).Error
	return
}
