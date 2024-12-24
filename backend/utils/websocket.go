package utils

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return err
	}
	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("メッセージの読み取りに失敗しました:", err)
			break
		}

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("メッセージの書き込みに失敗しました:", err)
			break
		}
	}
	return nil
}
