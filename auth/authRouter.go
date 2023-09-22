package auth

import "github.com/labstack/echo/v4"

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

func getUser(c echo.Context) error {
	var user User
	return c.JSON(200, user)
}

func AuthRouter(e *echo.Echo) {
	auth := e.Group("auth")
	auth.GET("/me", getUser)
}
