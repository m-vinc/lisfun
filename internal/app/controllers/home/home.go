package home

import (
	"lisfun/internal/app/context"
	apperrors "lisfun/internal/app/errors"
	"lisfun/internal/app/views/pages/home"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (homeController *homeController) Home(echoContext echo.Context, requestContext *context.RequestContext) error {
	requestContext.ViewContext.SetTitle("Home")

	return errors.WithStack(
		home.Home(
			requestContext,
		).Render(echoContext.Request().Context(), echoContext.Response().Writer),
	)
}

func (homeController *homeController) Err(_ echo.Context, requestContext *context.RequestContext) error {
	return errors.WithStack(apperrors.ErrNotFound)
}
