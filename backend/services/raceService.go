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
			return &models.RaceSchedule{}, nil
		}
		return nil, err
	}
	return &raceSchedule, nil
}

func GetRaceSchedulesByDate(db *gorm.DB) (map[string][]models.RaceSchedule, error) {
	var raceSchedules []models.RaceSchedule

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := db.Where("race_time >= ? AND race_time < ?", startOfDay, endOfDay).
		Order("race_time ASC").Find(&raceSchedules).Error
	if err != nil {
		return nil, fmt.Errorf("レース情報の取得に失敗しました: %v", err)
	}

	// course_nameでグループ化
	groupedByCourse := make(map[string][]models.RaceSchedule)
	for _, raceSchedule := range raceSchedules {
		groupedByCourse[raceSchedule.CourseName] = append(groupedByCourse[raceSchedule.CourseName], raceSchedule)
	}

	return groupedByCourse, nil
}
