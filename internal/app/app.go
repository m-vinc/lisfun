package app

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type App struct {
	*echo.Echo
	o sync.Once

	config *Config
	logger zerolog.Logger
}

type Config struct {
	Env  string
	Port string

	LogLevel string
}

func New(config *Config) (*App, error) {
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
		Echo: app.Echo,
	}
}

func (app *App) Run(_ context.Context) error {
	return errors.WithStack(app.Start(app.config.Port))
}
