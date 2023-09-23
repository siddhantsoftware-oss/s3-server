package main

import (
	"tecna/example"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	example.ExampleRouter(e)
}
