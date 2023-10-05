package auth

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func createNewRandomState(c echo.Context) string {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(100000000))
	cookie := new(http.Cookie)
	cookie.Domain = "/"
	cookie.Value = randInt.String()
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Name = "oauth-state"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return randInt.String()
}

func getState(c echo.Context) string {
	var state string
	cookie, err := c.Cookie("oauth-state")
	if err != nil {
		state = createNewRandomState(c)
	} else {
		state = cookie.Name
	}
	return state
}

func validateState(c echo.Context, givenState string) bool {
	state := getState(c)
	return state == givenState
}
