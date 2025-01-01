package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/database"
	"backend/routes"
)

func main() {

	e := echo.New()

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Reactの開発サーバーを許可
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	database.Init()

	routes.InitRoutes(e)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
