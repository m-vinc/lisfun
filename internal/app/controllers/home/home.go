package home

import (
	"lisfun/internal/app/context"
	apperrors "lisfun/internal/app/errors"
	"lisfun/internal/app/views/pages/home"
	"lisfun/internal/services/models"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (homeController *homeController) Home(echoContext echo.Context, requestContext *context.RequestContext) error {
	requestContext.ViewContext.SetTitle("Home")

	if requestContext.UserContext != nil {
		playbackState, err := homeController.SpotifyService.PlaybackState(echoContext.Request().Context(), &models.SpotifyPlaybackStateParams{
			User: requestContext.UserContext.User,
		})
		homeController.Logger.Info().Err(err).Any("player", playbackState).Msg("currently_playing")
	}

	return errors.WithStack(
		home.Home(
			requestContext,
		).Render(echoContext.Request().Context(), echoContext.Response().Writer),
	)
}

func (homeController *homeController) Err(_ echo.Context, requestContext *context.RequestContext) error {
	return errors.WithStack(apperrors.ErrNotFound)
}
