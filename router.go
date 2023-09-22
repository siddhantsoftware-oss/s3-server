package main

import (
	"server/example"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	example.ExampleRouter(e)
}
