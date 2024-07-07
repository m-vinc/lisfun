package app

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"lisfun/internal/app/models"
	"lisfun/internal/db"
	"lisfun/internal/services/spotify"
	"lisfun/internal/services/tokens"
	"lisfun/internal/services/users"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type App struct {
	*echo.Echo
	o sync.Once

	logger zerolog.Logger
	config *models.AppConfig

	database *db.Client

	usersService   *users.Service
	tokensService  *tokens.Service
	spotifyService *spotify.Service
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

		err = app.Database()
		if err != nil {
			return
		}

		err = app.Services()
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

		Logger:         app.logger,
		UsersService:   app.usersService,
		TokensService:  app.tokensService,
		SpotifyService: app.spotifyService,
	}
}

func (app *App) Run(_ context.Context) error {
	return errors.WithStack(app.Start(app.config.Port))
}
