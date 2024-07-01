package home

import (
	"lisfun/internal/app/context"
)

func Controller(context *context.AppContext) error {
	return (&homeController{AppContext: context}).Register()
}

type homeController struct {
	*context.AppContext
}

func (homeController *homeController) Register() error {
	homeController.GET("/", context.RequestContextWrap(homeController.Home))
	homeController.GET("/err", context.RequestContextWrap(homeController.Err))

	return nil
}
