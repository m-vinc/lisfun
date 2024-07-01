package auth

import "lisfun/internal/app/context"

func Controller(context *context.AppContext) error {
	return (&authController{AppContext: context}).Register()
}

type authController struct {
	*context.AppContext
}

func (authController *authController) Register() error {
	authController.Use(authController.Authenticate())

	authController.GET("/:provider", context.RequestContextWrap(authController.Initiate))
	authController.GET("/:provider/callback", context.RequestContextWrap(authController.Callback))

	return nil
}
