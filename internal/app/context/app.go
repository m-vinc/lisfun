package context

import (
	"lisfun/internal/app/models"
	"lisfun/internal/services/users"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type AppContext struct {
	*echo.Echo

	Logger zerolog.Logger
	Config *models.AppConfig

	UsersService *users.Service
}
