package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 部屋のメッセージを取得
func GetMessages(c echo.Context) error {
	roomId := c.Param("roomId")
	messages, err := models.GetMessages(database.DB, roomId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "メッセージが取得できません"})
	}
	return c.JSON(http.StatusOK, messages)
}
