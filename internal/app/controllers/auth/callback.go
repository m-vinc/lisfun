package auth

import (
	"context"
	appcontext "lisfun/internal/app/context"
	"lisfun/internal/services/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
)

func (authController *authController) Callback(echoContext echo.Context, requestContext *appcontext.RequestContext) error {
	request := echoContext.Request().WithContext(context.WithValue(echoContext.Request().Context(), "provider", echoContext.Param("provider")))

	user, err := gothic.CompleteUserAuth(echoContext.Response().Writer, request)
	if err != nil {
		// Maybe add a flash but for now just redirect to the home after clearing the store
		authController.Logger.Info().Err(err).Msg("an error occured while completing login")
		gothic.Logout(echoContext.Response(), request)
		return echoContext.Redirect(307, "/")
	}

	u, err := authController.UsersService.Onboard(request.Context(), &models.UsersOnboardParams{
		GothUser: &user,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = authController.TokensService.Create(request.Context(), &models.TokensCreateParams{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
		User:         u,
		ExpireAt:     user.ExpiresAt,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	authController.Logger.Info().Any("user", u).Msg("user onboarded")

	session, err := authController.store.Get(request, LisfunSessionKey)
	if err != nil {
		// maybe add a flash
		authController.clearSession(session, echoContext.Request(), echoContext.Response().Writer)
		return echoContext.Redirect(http.StatusTemporaryRedirect, "/")
	}

	session.Values = map[any]any{"user_context": u.ID.String()}

	err = session.Save(request, echoContext.Response().Writer)
	if err != nil {
		return errors.WithStack(err)
	}

	return echoContext.Redirect(http.StatusTemporaryRedirect, "/")
}
