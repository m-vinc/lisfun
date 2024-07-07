package auth

import (
	"lisfun/internal/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (authController *authController) Logout(echoContext echo.Context, requestContext *context.RequestContext) error {
	session, err := authController.store.Get(echoContext.Request(), LisfunSessionKey)
	if err != nil {
		return err
	}

	if !session.IsNew {
		err = authController.clearSession(session, echoContext.Request(), echoContext.Response().Writer)
		if err != nil {
			return err
		}
	}

	return echoContext.Redirect(http.StatusTemporaryRedirect, "/")
}
