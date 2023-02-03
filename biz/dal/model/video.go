package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id" column:"author_id"`
	PlayUrl       string `json:"play_url" column:"play_url"`
	CoverUrl      string `json:"cover_url" column:"cover_url"`
	FavoriteCount int64  `json:"favorite_count" column:"favorite_count"`
	CommentCount  int64  `json:"comment_count" column:"comment_count"`
	Likes         string `json:"likes" column:"likes"`
	Comments      string `json:"comments" column:"comments"`
	Title         string `json:"title" column:"title"`
}

func (v *Video) TableName() string {
	return "video"
}

func FindVideoById(v *Video, id uint) error {
	return dal.DB.First(&v, id).Error
}

func FindVideosByUserId(v []*Video, id int64) error {
	return dal.DB.Where("author_id = ?", id).Find(&v).Error
}

func CreateVideo(v *Video) error {
	return dal.DB.Create(v).Error
}
