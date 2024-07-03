package auth

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"lisfun/internal/app/views/pages/home"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
)

func (authController *authController) Callback(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	request := echoContext.Request().WithContext(context.WithValue(echoContext.Request().Context(), "provider", echoContext.Param("provider")))

	user, err := gothic.CompleteUserAuth(echoContext.Response().Writer, request)

	log.Printf("%+v %+v %T", user, err, err)
	if err != nil {
		// Maybe add a flash but for now just redirect to the home after clearing the store
		gothic.Logout(echoContext.Response(), request)
		return echoContext.Redirect(307, "/")
	}

	return errors.WithStack(
		home.Home(
			requestContext,
		).
			Render(echoContext.Request().Context(), echoContext.Response().Writer),
	)
}
