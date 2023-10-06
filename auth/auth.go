package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func setupGoogleConfig() {
	googleConfig = &oauth2.Config{
		ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  viper.GetString("API_URL") + "/auth/google/callback",
	}
}

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

var googleConfig *oauth2.Config

func loginWithGoogle(c echo.Context) error {
	state := getState(c)
	url := googleConfig.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if validateState(c, c.QueryParam("state")) {
		token, err := googleConfig.Exchange(c.Request().Context(), code)
		if err != nil {
			return c.String(500, "Something went wrong our system.")
		}
		accessTokenCookie := new(http.Cookie)
		accessTokenCookie.Name = "access-token"
		accessTokenCookie.Value = token.AccessToken
		accessTokenCookie.Expires = time.Now().Add(time.Hour * 24)
		accessTokenCookie.Path = "/"
		accessTokenCookie.Domain = viper.GetString("COOKIE_DOMAIN")
		c.SetCookie(accessTokenCookie)
		return c.Redirect(http.StatusPermanentRedirect, viper.GetString("CLIENT_URL"))

	}
	return c.Redirect(http.StatusPermanentRedirect, viper.GetString("CLIENT_URL")+"?err=NotFound")

}

func signOutUser(c echo.Context) error {
	nullCookie := new(http.Cookie)
	nullCookie.Name = "access-token"
	nullCookie.Value = ""
	nullCookie.Expires = time.Now()
	nullCookie.Path = "/"
	nullCookie.Domain = viper.GetString("COOKIE_DOMAIN")
	c.SetCookie(nullCookie)
	return c.Redirect(http.StatusPermanentRedirect, viper.GetString("CLIENT_URL"))
}

func AuthRouter(e *echo.Echo) {
	setupGoogleConfig()
	auth := e.Group("auth")
	auth.GET("/me", getUser)
	auth.GET("/login/google", loginWithGoogle)
	auth.GET("/google/callback", googleCallback)
	auth.GET("/sign-out", signOutUser)
}
