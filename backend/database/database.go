package database

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// データベース接続の初期化
func Init(user, password, dbname, host, port string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	// マイグレーション
	err = DB.AutoMigrate(&models.Room{})
	if err != nil {
		log.Fatalf("データベースのマイグレーションに失敗しました。: %v", err)
	}

	log.Println("データベース接続が確立されました")
}
