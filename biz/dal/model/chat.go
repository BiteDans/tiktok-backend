package model

import (
	"time"

	"BiteDans.com/tiktok-backend/biz/dal"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ToUserId   int64  `json:"to_user_id" column:"to_user_id"`
	FromUserId int64  `json:"from_user_id" column:"from_user_id"`
	Content    string `json:"content" column:"content"`
}

func (m *Message) TableName() string {
	return "message"
}

func FindMessageById(m *Message, id uint) error {
	return dal.DB.First(&m, id).Error
}

func FindMessageBySenderandReceiverIdBeforeTime(m []*Message, senderId uint, receiverId uint, time time.Time) ([]*Message, error) {
	var message []*Message
	err := dal.DB.Where("((to_user_id = ? AND from_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)) AND created_at > ?", receiverId, senderId, receiverId, senderId, time).Order("created_at asc").Find(&message).Error
	return message, err
}

func CreateMessage(m *Message) error {
	return dal.DB.Create(m).Error
}

func FindLatestMessage(fromUserId int64, toUserId int64) (*Message, error) {
	var message *Message
	err := dal.DB.Where("(to_user_id = ? AND from_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", toUserId, fromUserId, toUserId, fromUserId).Order("created_at desc").First(&message).Error
	return message, err
}