package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string  `json:"username" column:"username"`
	Password   string  `json:"password" column:"password"`
	Avatar	   string  `json:"avatar" column:"avatar" gorm:"default:https://tiktok-bitedans.s3.us-west-1.amazonaws.com/%E5%A4%B4%E5%83%8F.jpg"`
	BackgroundImage	   string  `json:"backgroundImage" column:"background_image" gorm:"default:https://tiktok-bitedans.s3.us-west-1.amazonaws.com/%E8%83%8C%E6%99%AF.jpg"`
	Signature	   string  `json:"signature" column:"signature" gorm:"default:喜欢唱、跳、Rap、篮球。"`
	Followings []*User `gorm:"many2many:follow_relations;joinForeignKey:user_id;JoinReferences:follow_id"`
	Followers  []*User `gorm:"many2many:follow_relations;joinForeignKey:follow_id;JoinReferences:user_id"`
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

func GetFollowRelation(fromUserId uint, toUserId uint) (bool, error) {
	var err error

	from_user := User{}
	from_user.ID = fromUserId

	result := []User{}
	err = dal.DB.Model(&from_user).Where("follow_id", toUserId).Association("Followings").Find(&result)
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

func UserFollow(uc *User, ut *User) error {
	return dal.DB.Model(uc).Association("Followings").Append(ut)
}

func UserUnfollow(uc *User, ut *User) error {
	return dal.DB.Model(uc).Association("Followings").Delete(ut)
}

func GetFollowListByUser(uList *[]*User, user *User) error {
	return dal.DB.Model(user).Association("Followings").Find(&uList)
}

func GetFollowerListByUser(uList *[]*User, user *User) error {
	return dal.DB.Model(user).Association("Followers").Find(&uList)
}

func GetFriendListById(friendIds *[]uint, id int64) error {
	return dal.DB.Raw(
		`SELECT fr1.follow_id FROM follow_relations AS fr1
		 INNER JOIN follow_relations AS fr2
		 ON fr1.follow_id = fr2.user_id
		 AND fr1.user_id = fr2.follow_id
		 WHERE fr1.user_id = ?
		 ORDER BY 1`, id).Scan(friendIds).Error
}
