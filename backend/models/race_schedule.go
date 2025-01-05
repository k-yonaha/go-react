package models

import (
	"time"
)

type Participant struct {
	CourseNumber string  `json:"course_number"` // 艇番
	PlayerNumber string  `json:"player_number"` // 選手登番
	Name         string  `json:"name"`          // 名前
	Age          int     `json:"age"`           // 年齢
	Region       string  `json:"region"`        // 支部
	Grade        string  `json:"grade"`         // 体級種別
	OtherInfo    *string `json:"other_info"`
}

type RaceSchedule struct {
	ID           uint
	CourseName   string    // ボートレース場
	RaceDate     time.Time `gorm:"type:date"` // 日付
	RaceProgram  string    `gorm:"type:text"` // 番組名
	RaceNumber   int       // 1R,2R
	RaceType     string    `gorm:"type:text"` // 予選、特選
	RaceDay      string    `gorm:"type:text"` // ◯日目
	RaceTime     time.Time `gorm:"type:time"` // 締切時刻
	Participants []byte    `gorm:"type:json"` // 選手情報
	Result       string    `gorm:"type:text"` // 結果
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
