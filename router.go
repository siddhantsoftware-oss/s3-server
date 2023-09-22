package main

import (
	"s3-server/example"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	example.ExampleRouter(e)
}
