package home

import (
	"lisfun/internal/app/common"
)

func Controller(context *common.AppContext) error {
	return (&homeController{AppContext: context}).Register()
}

type homeController struct {
	*common.AppContext
}

func (homeController *homeController) Register() error {
	homeController.GET("/", homeController.Home)
	homeController.GET("/err", homeController.Err)

	return nil
}
