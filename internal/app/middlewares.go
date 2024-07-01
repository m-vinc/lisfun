package app

import (
	"lisfun/internal/app/context"

	"github.com/labstack/echo/v4"
)

func (app *App) InjectRequestContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestContext := context.DefaultRequestContext(app.Context())

			requestContext.SetViewContext(context.DefaultViewContext(requestContext))

			c.Set(string(context.RequestContextKey), requestContext)
			return next(c)
		}
	}
}
