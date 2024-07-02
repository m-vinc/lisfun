package auth

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (authController *authController) Initiate(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	request := echoContext.Request().WithContext(context.WithValue(echoContext.Request().Context(), "provider", echoContext.Param("provider")))

	if gothUser, err := gothic.CompleteUserAuth(echoContext.Response().Writer, request); err == nil {
		log.Println(gothUser, err)
		return echoContext.NoContent(204)
	}

	// url, err := gothic.GetAuthURL(echoContext.Response().Writer, request)
	// if err != nil {
	// 	return err
	// }

	// Needed to full reload the client using HTMX
	gothic.BeginAuthHandler(echoContext.Response().Writer, request)
	return nil
}
