package main

import (
	"backend/config"
	"backend/database"
	"backend/seeder/seeds"
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	log.Println("Seeding rooms...")
	seeds.RoomSeeder(db) // room_seeder.go を呼び出し
}

func main() {
	// 設定を読み込む
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	database.Init(config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	db := database.DB

	// Seeder を実行
	Run(db)
}
