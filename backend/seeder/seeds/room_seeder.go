package seeds

import (
	"backend/models"
	"log"

	"gorm.io/gorm"
)

func RoomSeeder(db *gorm.DB) {
	var roomCount int64
	db.Model(&models.Room{}).Count(&roomCount)
	if roomCount > 0 {
		log.Println("roomテーブルが既に存在しています。")
		return
	}

	rooms := []models.Room{
		// 関東地区
		{Name: "桐生"},
		{Name: "戸田"},
		{Name: "江戸川"},
		{Name: "平和島"},
		{Name: "多摩川"},

		// 東海地区
		{Name: "浜名湖"},
		{Name: "蒲郡"},
		{Name: "常滑"},
		{Name: "津"},

		// 近畿地区
		{Name: "三国"},
		{Name: "びわこ"},
		{Name: "住之江"},
		{Name: "尼崎"},

		// 四国地区
		{Name: "鳴門"},
		{Name: "丸亀"},

		// 中国地区
		{Name: "児島"},
		{Name: "宮島"},
		{Name: "徳山"},
		{Name: "下関"},

		// 九州地区
		{Name: "若松"},
		{Name: "芦屋"},
		{Name: "福岡"},
		{Name: "唐津"},
		{Name: "大村"},
	}

	for _, room := range rooms {
		if err := db.Create(&room).Error; err != nil {
			log.Fatalf("roomテーブルのseederが失敗: %v", err)
		}
	}
	log.Println("roomテーブルのseederが完了しました。")
}
