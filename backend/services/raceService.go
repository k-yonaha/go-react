package services

import (
	"backend/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateRaceSchedule(db *gorm.DB, schedule models.RaceSchedule) error {
	return db.Create(&schedule).Error
}

func RaceScheduleExists(db *gorm.DB, date time.Time) (bool, error) {

	var count int64
	// race_dateが一致するレース情報が存在するか確認
	err := db.Model(&models.RaceSchedule{}).Where("race_date = ?", date).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("レース情報の存在確認に失敗しました: %v", err)
	}
	return count > 0, nil
}

func GetRaceScheduleByCourseName(db *gorm.DB, courseName string) (*models.RaceSchedule, error) {
	var raceSchedule models.RaceSchedule
	err := db.Where("course_name = ? AND race_time > ?", courseName, time.Now()).
		Order("race_time ASC").First(&raceSchedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.RaceSchedule{}, nil // ここで空の構造体を返す
		}
		return nil, err
	}
	return &raceSchedule, nil
}
