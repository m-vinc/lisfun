package home

import (
	"lisfun/internal/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (homeController *homeController) Home(echoContext echo.Context) error {
	return errors.WithStack(echoContext.HTML(http.StatusOK, "<p>toto</p>"))
}

func (homeController *homeController) Err(_ echo.Context) error {
	return errors.WithStack(models.ErrNotFound)
}
