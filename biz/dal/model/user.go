package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"BiteDans.com/tiktok-backend/biz/model/douyin/extra/follow"
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

func GetFollowInfoByUsers(from_user *User, to_user *User, user_resp *follow.User) error {
	var err error
	user_resp.ID = int64(to_user.ID)
	user_resp.Name = to_user.Username
	user_resp.FollowCount = GetFollowCount(to_user)
	user_resp.FollowerCount = GetFollowerCount(to_user)
	if user_resp.IsFollow, err = GetFollowRelation(from_user.ID, to_user.ID); err != nil {
		return err
	}
	return nil
}

func GetFollowRelation(from_user_id uint, to_user_id uint) (bool, error) {
	var err error

	from_user := User{}
	from_user.ID = from_user_id

	result := []User{}
	err = dal.DB.Model(&from_user).Where("follow_id", to_user_id).Association("Followings").Find(&result)
	if err != nil {
		return false, err
	}

	return !(len(result) == 0), nil
}

func GetFollowCount(user *User) int64 {
	return dal.DB.Model(user).Association("Followings").Count()
}

func GetFollowerCount(user *User) int64 {
	return dal.DB.Model(user).Association("Followers").Count()
}

func UserFollowAction(uc *User, ut *User, t uint) error {
	if t == 1 {
		return dal.DB.Model(uc).Association("Followings").Append(ut)
	}
	return dal.DB.Model(uc).Association("Followings").Delete(ut)
}

func GetFollowListByUser(uList *[]*User, user *User) error {
	return dal.DB.Model(user).Association("Followings").Find(&uList)
}

func GetFollowerListByUser(uList *[]*User, user *User) error {
	return dal.DB.Model(user).Association("Followers").Find(&uList)
}
