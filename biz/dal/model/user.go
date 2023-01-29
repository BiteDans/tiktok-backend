package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" column:"username"`
	Password string `json:"password" column:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func RegisterUser(u *User) error {
	return dal.DB.Create(u).Error
}
