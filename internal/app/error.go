package app

import (
	"lisfun/internal/app/common"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (app *App) ErrorHandler() error {
	app.HTTPErrorHandler = func(err error, echoContext echo.Context) {
		requestID := echoContext.Response().Header().Get(echo.HeaderXRequestID)
		accept := echoContext.Request().Header.Get("Accept")

		httpError := &echo.HTTPError{}
		if !errors.As(err, &httpError) {
			httpError = common.ErrInternalServerError
		}

		app.logger.Error().
			Str("request_id", requestID).
			Stack().
			Err(err).
			Msg("")

		switch accept {
		case "application/json":
			_ = echoContext.JSON(httpError.Code, map[string]any{"code": httpError.Code, "message": httpError.Message})
		default:
			_ = echoContext.HTML(httpError.Code, "<p>error</p>")
		}
	}

	return nil
}
