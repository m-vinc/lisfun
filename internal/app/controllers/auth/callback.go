package auth

import (
	"lisfun/internal/app/context"

	"github.com/labstack/echo/v4"
)

func (authController *authController) Callback(echoContext echo.Context, requestContext *context.RequestContext) error {
	return echoContext.NoContent(204)

	//return errors.WithStack(
	//	home.Home(
	//		common.DefaultViewContext(homeController.AppContext),
	//	).
	//		Render(echoContext.Request().Context(), echoContext.Response().Writer),
	//)
}
