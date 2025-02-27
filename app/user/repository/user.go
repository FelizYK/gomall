package repository

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique"`
	PasswordHashed string
}

func (u User) TableName() string {
	return "user"
}

func GetById(ctx context.Context, id uint32) (user *User, err error) {
	err = DB.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).First(&user).Error
	return
}

func GetByEmail(ctx context.Context, email string) (user *User, err error) {
	err = DB.WithContext(ctx).Model(&User{}).
		Where("email = ?", email).First(&user).Error
	return
}

func Create(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).
		Create(user).Error
}
