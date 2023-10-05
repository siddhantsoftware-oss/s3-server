package main

import (
	"tecna/auth"
	"tecna/example"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	example.ExampleRouter(e)
	auth.AuthRouter(e)
}
