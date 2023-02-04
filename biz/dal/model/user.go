package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string  `json:"username" column:"username"`
	Password   string  `json:"password" column:"password"`
	Followings []*User `gorm:"many2many:follow_relation;joinForeignKey:user_Id;JoinReferences:following_id"`
	Followers  []*User `gorm:"many2many:follow_relation;joinForeignKey:following_id;JoinReferences:user_Id"`
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

func CreateUser(u *User) error {
	return dal.DB.Create(u).Error
}

func UserFollowAction(uc *User, ut *User, t uint) error {
	if t == 1 {
		return dal.DB.Model(uc).Association("Followings").Append(ut)
	}
	return dal.DB.Model(uc).Association("Followings").Delete(ut)
}
