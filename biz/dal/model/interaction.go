package model

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  int64  `json:"user_id" column:"user_id"`
	VideoId int64  `json:"video_id" column:"video_id"`
	Content string `json:"content" column:"content"`
}

type Like struct {
	gorm.Model
	UserId  int64 `json:"user_id" column:"user_id"`
	VideoId int64 `json:"video_id" column:"video_id"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (l *Like) TableName() string {
	return "like"
}

func FindCommentsByVideoId(id int64) ([]*Comment, error) {
	var _comments []*Comment
	err := dal.DB.Where("video_id = ?", id).Order("created_at desc").Find(&_comments).Error
	return _comments, err
}

func FindCommentsByUserId(id int64) ([]*Comment, error) {
	var _comments []*Comment
	err := dal.DB.Where("user_id = ?", id).Find(&_comments).Error
	return _comments, err
}

func FindCommentsByUserIdAndVideoId(user_id int64, video_id int64) ([]*Comment, error) {
	var _comments []*Comment
	err := dal.DB.Where("user_id = ? AND video_id = ?", user_id, video_id).Find(&_comments).Error
	return _comments, err
}

func FindCommentById(id int64) (*Comment, error) {
	var _comments *Comment
	err := dal.DB.Where("id = ?", id).First(&_comments).Error
	return _comments, err
}

func GetCommentCount(id int64) (int64, error) {
	var commentCount int64
	err := dal.DB.Model(Comment{}).Where("video_id = ?", id).Count(&commentCount).Error
	return commentCount, err
}

func GetLikeCount(id int64) (int64, error) {
	var likeCount int64
	err := dal.DB.Model(Like{}).Where("video_id = ?", id).Count(&likeCount).Error
	return likeCount, err
}

func CreateComment(c *Comment) error {
	return dal.DB.Create(c).Error
}

func FindLikedVideosIdByUserId(user_id int64) ([]*Like, error) {
	var likes []*Like
	err := dal.DB.Where("user_id = ?", user_id).Find(&likes).Error
	return likes, err
}

func LikeVideo(l *Like) error {
	return dal.DB.Create(l).Error
}

func IsVideoLiked(l *Like) error {
	return dal.DB.Where("user_id = ? AND video_id = ?", l.UserId, l.VideoId).First(&l).Error
}

func UnlikeVideo(l *Like) error {
	return dal.DB.Where("user_id = ? AND video_id = ?", l.UserId, l.VideoId).Delete(&l).Error
}

func DeleteComment(c *Comment) error {
	return dal.DB.Delete(&c).Error
}
