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

type UserQuery struct {
	ctx context.Context
}

func NewUserQuery(ctx context.Context) UserQuery {
	return UserQuery{ctx: ctx}
}

func (q UserQuery) GetById(id uint32) (user *User, err error) {
	err = db.WithContext(q.ctx).Model(&User{}).
		Where("id = ?", id).First(&user).Error
	return
}

func (q UserQuery) GetByEmail(email string) (user *User, err error) {
	err = db.WithContext(q.ctx).Model(&User{}).
		Where("email = ?", email).First(&user).Error
	return
}

func (q UserQuery) Create(user *User) error {
	return db.WithContext(q.ctx).
		Create(user).Error
}
