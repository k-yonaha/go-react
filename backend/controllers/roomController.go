package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 部屋の一覧を取得
func GetRooms(c echo.Context) error {
	rooms, err := models.GetAllRooms(database.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "部屋が取得できません"})
	}
	return c.JSON(http.StatusOK, rooms)
}

// 新しい部屋を作成
func CreateRoom(c echo.Context) error {
	name := c.FormValue("name")
	room := models.Room{Name: name}

	if err := models.CreateRoom(database.DB, &room); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "部屋が作成できません"})
	}
	return c.JSON(http.StatusCreated, room)
}
