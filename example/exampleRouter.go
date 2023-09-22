package example

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func ExampleRouter(e *echo.Echo) {
	hello := e.Group("/hello")
	hello.GET("/world", helloWorld)
}
