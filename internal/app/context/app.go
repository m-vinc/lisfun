package context

import (
	"lisfun/internal/app/models"

	"github.com/labstack/echo/v4"
)

type AppContext struct {
	*echo.Echo

	Config *models.AppConfig
}
