package main

import (
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
	database.Init()

	db := database.DB

	// Seeder を実行
	Run(db)
}
