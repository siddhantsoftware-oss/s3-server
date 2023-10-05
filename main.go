package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	InitEnv()
	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowCredentials: true,
	}))
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${status}] ${method} ${uri}\n",
		Output: router.Logger.Output(),
	}))
	Routes(router)
	router.Logger.Fatal(router.Start(":3000"))
}
