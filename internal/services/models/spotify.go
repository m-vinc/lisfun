package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SpotifyPlaybackStateParams struct {
	User *User
}

type SpotifyPlaybackStateResult struct{}

var ErrTokenNotFound = echo.NewHTTPError(http.StatusForbidden, "unable to get a valid user's token")
