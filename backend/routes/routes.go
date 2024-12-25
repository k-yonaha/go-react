package routes

import (
	"backend/controllers"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// 部屋一覧を取得する
	e.GET("/api/rooms", controllers.GetRooms)
	e.GET("/api/rooms/:roomId/messages", controllers.GetMessages)
	e.POST("/api/rooms", controllers.CreateRoom)
	e.GET("/ws", utils.HandleWebSocket)
}
