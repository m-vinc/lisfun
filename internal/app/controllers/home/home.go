package home

import (
	"lisfun/internal/app/common"
	"lisfun/internal/app/views/pages/home"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (homeController *homeController) Home(echoContext echo.Context) error {
	return errors.WithStack(
		home.Home(
			common.DefaultViewContext(homeController.AppContext),
		).
			Render(echoContext.Request().Context(), echoContext.Response().Writer),
	)
}

func (homeController *homeController) Err(_ echo.Context) error {
	return errors.WithStack(common.ErrNotFound)
}
