package app

import (
	"lisfun/internal/app/context"
	apperrors "lisfun/internal/app/errors"
	pageserrors "lisfun/internal/app/views/pages/errors"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (app *App) ErrorHandler() error {
	app.HTTPErrorHandler = func(err error, echoContext echo.Context) {
		requestContext := context.RequestContextFromEcho(echoContext)

		requestID := echoContext.Response().Header().Get(echo.HeaderXRequestID)
		accept := echoContext.Request().Header.Get("Accept")

		httpError := &echo.HTTPError{}
		if !errors.As(err, &httpError) {
			httpError = apperrors.ErrInternalServerError
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
			echoContext.Response().WriteHeader(httpError.Code)

			_ = pageserrors.Error(
				requestContext,
				httpError,
			).Render(
				echoContext.Request().Context(),
				echoContext.Response().Writer,
			)
		}
	}

	return nil
}
