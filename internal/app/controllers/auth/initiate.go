package auth

import (
	"context"
	appcontext "lisfun/internal/app/context"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (authController *authController) Initiate(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	request := echoContext.Request().WithContext(context.WithValue(echoContext.Request().Context(), "provider", echoContext.Param("provider")))

	if gothUser, err := gothic.CompleteUserAuth(echoContext.Response().Writer, request); err == nil {
		authController.Logger.Info().
			Any("goth_user", gothUser).
			Msg("already logged in, reloading cookies")
		return echoContext.NoContent(204)
	}

	// url, err := gothic.GetAuthURL(echoContext.Response().Writer, request)
	// if err != nil {
	// 	return err
	// }

	gothic.BeginAuthHandler(echoContext.Response().Writer, request)
	return nil
}
