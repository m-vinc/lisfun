package auth

import (
	"lisfun/internal/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (authController *authController) Initiate(echoContext echo.Context, requestContext *context.RequestContext) error {
	return echoContext.Redirect(http.StatusFound, "")
}
