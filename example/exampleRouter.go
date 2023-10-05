package example

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorld(c echo.Context) error {
	token, err := c.Cookie("access-token")
	if err != nil {
		return c.String(200, "Hello, World!")
	}
	res, _ := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + token.Value)

	user, _ := io.ReadAll(res.Body)
	return c.JSON(http.StatusFound, string(user))
}

func ExampleRouter(e *echo.Echo) {
	hello := e.Group("/hello")
	hello.GET("/world", helloWorld)
}
