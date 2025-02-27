package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserId    uint
	ProductId uint
	Quantity  int
}

func (u CartItem) TableName() string {
	return "CartItem"
}

func GetCartByUserId(ctx context.Context, userId uint) (items []CartItem, err error) {
	err = DB.WithContext(ctx).Model(&CartItem{}).
		Where("user_id = ?", userId).Find(&items).Error
	return
}

func AddCart(ctx context.Context, userId uint, productId uint, quantity int) (err error) {
	var find CartItem
	err = DB.WithContext(ctx).Model(&CartItem{}).
		Where("user_id = ? AND product_id = ?", userId, productId).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if find.Model.ID == 0 {
		// create
		err = DB.WithContext(ctx).Model(&CartItem{}).
			Create(&CartItem{
				UserId:    userId,
				ProductId: productId,
				Quantity:  quantity,
			}).Error
	} else {
		// update
		err = DB.WithContext(ctx).Model(&CartItem{}).
			Where("user_id = ? AND product_id = ?", userId, productId).
			Update("quantity", find.Quantity+quantity).Error
	}
	return
}

func EmptyCart(ctx context.Context, userId uint) (err error) {
	err = DB.WithContext(ctx).Model(&CartItem{}).
		Where("user_id = ?", userId).Delete(&CartItem{}).Error
	return
}
