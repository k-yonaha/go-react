package database

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// データベース接続の初期化
func Init(user, password, dbname, host, port string) {
	log.Println("データベース接続が確立されました:", user)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
	}

	// マイグレーション
	err = DB.AutoMigrate(&models.Room{}, &models.RaceSchedule{})
	if err != nil {
		log.Fatalf("データベースのマイグレーションに失敗しました。: %v", err)
	}

	log.Println("データベース接続が確立されました")
}
