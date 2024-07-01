package app

import (
	"embed"
	"lisfun/internal/app/controllers/home"
	apperrors "lisfun/internal/app/errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

//go:embed assets/*
var assetsFS embed.FS

func (app *App) Routes() error {
	app.Use(app.InjectRequestContext())

	app.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	_ = home.Controller(app.Context())

	app.Static()
	app.CatchAll()

	return nil
}

func (app *App) Static() {
	app.StaticFS("/assets/*", echo.MustSubFS(assetsFS, "assets"))
}

func (app *App) CatchAll() {
	app.RouteNotFound("*", func(_ echo.Context) error {
		return errors.WithStack(apperrors.ErrNotFound)
	})
}
