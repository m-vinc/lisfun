package auth

import (
	"lisfun/internal/app/context"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

func Controller(context *context.AppContext) error {
	authController := (&authController{AppContext: context})

	err := spotifyProvider(&SpotifyProviderConfig{
		Key:         context.Config.SpotifyProvider.Key,
		Secret:      context.Config.SpotifyProvider.Secret,
		RedirectURL: context.Config.SpotifyProvider.RedirectURL,
	})
	if err != nil {
		return err
	}

	authController.store = authController.newStore()
	gothic.Store = authController.store

	return authController.Register()
}

type authController struct {
	*context.AppContext

	store sessions.Store
}

func (authController *authController) Register() error {
	authController.Use(authController.Authenticate())

	routes := authController.Group("/auth")

	routes.GET("/:provider", context.RequestContextWrap(authController.Initiate))
	routes.GET("/:provider/callback", context.RequestContextWrap(authController.Callback))

	return nil
}

func (authController *authController) newStore() sessions.Store {
	store := sessions.NewCookieStore([]byte(authController.AppContext.Config.SecretKey))
	return store
}
