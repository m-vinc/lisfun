package app

import (
	"lisfun/internal/app/context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func (app *App) Logger() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	loglevel, err := zerolog.ParseLevel(app.config.LogLevel)
	if err != nil {
		loglevel = zerolog.InfoLevel
	}

	logger := zerolog.New(os.Stdout).
		Level(loglevel).
		Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).
		With().
		Timestamp().
		Caller().
		Stack().
		Logger()

	app.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			requestContext := context.RequestContextFromEcho(c)

			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			l := logger.Info().
				Str("uri", v.URI).
				Int("status", v.Status).
				Str("request_id", requestID)

			if requestContext.UserContext != nil {
				l.Str("user_id", requestContext.UserContext.ID.String())
			}

			l.Msg("")

			return nil
		},
	}))

	app.logger = logger

	return nil
}
