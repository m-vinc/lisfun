package app

import (
	"lisfun/internal/app/controllers/home"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) Routes() error {
	app.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	_ = home.Controller(app.Context())

	return nil
}
