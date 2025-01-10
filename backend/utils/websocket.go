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

var rooms = make(map[string][]*websocket.Conn)

func HandleWebSocket(c echo.Context) error {
	roomId := c.Param("roomId")
	log.Printf("接続された部屋ID: %s\n", roomId)

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocketのアップグレードに失敗しました:", err)
		return c.String(http.StatusInternalServerError, "WebSocket接続のアップグレードに失敗しました")
	}
	defer conn.Close()

	// 部屋ごとに接続を管理
	rooms[roomId] = append(rooms[roomId], conn)

	log.Println("新しい接続が確立されました。")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("メッセージの読み取りに失敗しました:", err)
			break
		}

		// メッセージを同じ部屋にいる全てのクライアントに送信
		for _, c := range rooms[roomId] {
			err := c.WriteMessage(messageType, msg)
			if err != nil {
				log.Println("メッセージの送信に失敗しました:", err)
			}
		}
	}

	// 接続が切断された後にリストから削除
	for i, c := range rooms[roomId] {
		if c == conn {
			rooms[roomId] = append(rooms[roomId][:i], rooms[roomId][i+1:]...)
			break
		}
	}

	return nil
}
