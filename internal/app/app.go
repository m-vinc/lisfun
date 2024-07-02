package app

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"lisfun/internal/app/models"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type App struct {
	*echo.Echo
	o sync.Once

	config *models.AppConfig
	logger zerolog.Logger
}

func New(config *models.AppConfig) (*App, error) {
	app := &App{
		Echo: echo.New(),

		config: config,
	}

	var err error

	app.o.Do(func() {
		err = app.Logger()
		if err != nil {
			return
		}

		err = app.ErrorHandler()
		if err != nil {
			return
		}

		err = app.Routes()
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Context() *appcontext.AppContext {
	return &appcontext.AppContext{
		Echo:   app.Echo,
		Config: app.config,
	}
}

func (app *App) Run(_ context.Context) error {
	return errors.WithStack(app.Start(app.config.Port))
}
