package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"backend/config"
	"backend/database"
	"backend/routes"
)

func main() {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	e := echo.New()

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Reactの開発サーバーを許可
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	database.Init(config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	routes.InitRoutes(e)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
