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

func FindUserById(u *User, id uint) error {
	return dal.DB.First(&u, id).Error
}

func FindUserByUsername(u *User, username string) error {
	return dal.DB.First(&u, "username = ?", username).Error
}

func RegisterUser(u *User) error {
	return dal.DB.Create(u).Error
}
