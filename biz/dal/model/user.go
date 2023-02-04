package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"BiteDans.com/tiktok-backend/biz/model/douyin/core/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string  `json:"username" column:"username"`
	Password   string  `json:"password" column:"password"`
	Followings []*User `gorm:"many2many:follow_relations;joinForeignKey:user_id;JoinReferences:follow_id"`
	Followers  []*User `gorm:"many2many:follow_relations;joinForeignKey:follow_id;JoinReferences:user_id"`
}

func (u *User) TableName() string {
	return "users"
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

func GetFollowInfoByIDs(from_user_id uint, to_user_id uint, user_resp *user.User) error {
	from_user := User{}
	to_user := User{}

	var err error

	err = FindUserById(&from_user, from_user_id)
	if err != nil {
		return err
	}

	err = FindUserById(&to_user, to_user_id)
	if err != nil {
		return err
	}

	result := []User{}
	err = dal.DB.Model(&from_user).Where("follow_id", to_user_id).Association("Followings").Find(&result)
	if err != nil {
		return err
	}

	user_resp.ID = int64(to_user.ID)
	user_resp.Name = to_user.Username
	user_resp.FollowCount = getFollowCount(&to_user)
	user_resp.FollowerCount = getFollowerCount(&to_user)
	user_resp.IsFollow = !(len(result) == 0)

	return nil
}

func getFollowCount(user *User) int64 {
	return dal.DB.Model(user).Association("Followings").Count()
}

func getFollowerCount(user *User) int64 {
	return dal.DB.Model(user).Association("Followers").Count()
}

func UserFollowAction(uc *User, ut *User, t uint) error {
	if t == 1 {
		return dal.DB.Model(uc).Association("Followings").Append(ut)
	}
	return dal.DB.Model(uc).Association("Followings").Delete(ut)
}
