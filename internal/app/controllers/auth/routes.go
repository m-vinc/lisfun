package auth

import "lisfun/internal/app/context"

func Controller(context *context.AppContext) error {
	authController := (&authController{AppContext: context})

	err := spotifyProvider(&SpotifyProviderConfig{
		Key:         context.Config.SpotifyProvider.Key,
		Secret:      context.Config.SpotifyProvider.Key,
		RedirectURL: context.Config.SpotifyProvider.RedirectURL,
	})
	if err != nil {
		return err
	}

	return authController.Register()
}

type authController struct {
	*context.AppContext
}

func (authController *authController) Register() error {
	authController.Use(authController.Authenticate())

	routes := authController.Group("/auth")

	routes.GET("/:provider", context.RequestContextWrap(authController.Initiate))
	routes.GET("/:provider/callback", context.RequestContextWrap(authController.Callback))

	return nil
}
