package controllers

import (
	"backend/database"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 部屋の一覧を取得
func GetRooms(c echo.Context) error {
	rooms, err := services.GetAllRooms(database.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "部屋が取得できません"})
	}
	return c.JSON(http.StatusOK, rooms)
}
