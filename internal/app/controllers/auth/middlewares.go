package auth

import (
	"lisfun/internal/app/context"
	"log"

	"github.com/labstack/echo/v4"
)

// The goal of this middleware is to authenticate the request using anything we want and set a new
func (authController *authController) Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			requestContext := context.RequestContextFromEcho(echoContext)

			// Authenticate the request and call requestContext.SetUserContext

			log.Println(requestContext)
			return next(echoContext)
		}
	}
}
