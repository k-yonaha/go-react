package models

import (
	"gorm.io/gorm"
)

// Message モデル
type Message struct {
	ID     uint `gorm:"primary_key"`
	RoomID uint
	Body   string
}

func GetMessages(db *gorm.DB, roomId string) ([]Message, error) {
	var messages []Message
	if err := db.Where("room_id = ?", roomId).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
