package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// /api/hello のリクエストを処理するハンドラー
func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, world from Go (Echo)!",
	})
}

// /api/greet/{name} のリクエストを処理するハンドラー
func GetGreeting(c echo.Context) error {
	name := c.Param("name") // URLパラメータを取得
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, " + name + "!",
	})
}
