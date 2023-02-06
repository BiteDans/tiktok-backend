package model

import (
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

func FindMessageBySenderandReceiverId(m []*Message, sender_id uint, receiver_id uint) error {
	return dal.DB.Where("to_user_id = ? AND from_user_id = ?", receiver_id, sender_id).Find(&m).Error
}

func CreateMessage(m *Message) error {
	return dal.DB.Create(m).Error
}
