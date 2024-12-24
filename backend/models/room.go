package models

import (
	"gorm.io/gorm"
)

// Room モデル
type Room struct {
	ID   uint
	Name string
}

// 全ての部屋を取得
func GetAllRooms(db *gorm.DB) ([]Room, error) {
	var rooms []Room
	if err := db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

// 部屋を作成
func CreateRoom(db *gorm.DB, room *Room) error {
	return db.Create(room).Error
}
