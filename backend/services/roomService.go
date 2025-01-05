package services

import (
	"backend/models"

	"gorm.io/gorm"
)

// 全ての部屋を取得
func GetAllRooms(db *gorm.DB) ([]models.Room, error) {
	var rooms []models.Room
	if err := db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
