package home

import (
	"lisfun/internal/app/models"
)

func Controller(context *models.AppContext) error {
	return (&homeController{AppContext: context}).Register()
}

type homeController struct {
	*models.AppContext
}

func (homeController *homeController) Register() error {
	homeController.GET("/", homeController.Home)
	homeController.GET("/err", homeController.Err)

	return nil
}
