package model

import (
	"time"

	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id" column:"author_id"`
	AuthorUsername      string  `json:"author_username" column:"author_username"`
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

func FindUserByVideo(videoId int64) (int64, error) {
	var v *Video
	err := dal.DB.Where("id = ?", videoId).First(&v).Error
	return v.AuthorId, err
}

func FindVideosByUserId(id int64) ([]*Video, error) {
	var _videos []*Video
	err := dal.DB.Where("author_id = ?", id).Find(&_videos).Error
	return _videos, err
}

func CreateVideo(v *Video) error {
	return dal.DB.Create(v).Error
}

func GetVideoCount() (int64, error) {
	var totalRows int64
	err := dal.DB.Model(Video{}).Count(&totalRows).Error
	return totalRows, err
}

func GetUserVideoCount(id int64) (int64, error) {
	var totalRows int64
	err := dal.DB.Model(Video{}).Where("author_id = ?", id).Count(&totalRows).Error
	return totalRows, err
}

func FindLatestVideos(time time.Time) ([]*Video, error) {
	var _videos []*Video
	err := dal.DB.Order("created_at desc").Limit(30).Where("created_at < ?", time).Find(&_videos).Error
	return _videos, err
}