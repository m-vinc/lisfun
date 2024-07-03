package main

import (
	"lisfun/internal/app"
	"lisfun/internal/app/models"

	"github.com/coder/serpent"
	"github.com/pkg/errors"
)

func (rootCmd *RootCmd) run() *serpent.Command {
	var (
		port     string
		loglevel string
		env      string

		databaseURL string

		spotifyKey         string
		spotifySecret      string
		spotifyRedirectURL string
	)

	cmd := &serpent.Command{
		Use:   "run",
		Short: "Run lisfun",
		Handler: func(inv *serpent.Invocation) error {
			app, err := app.New(&models.AppConfig{
				Env:      env,
				Port:     port,
				LogLevel: loglevel,
				SpotifyProvider: &models.SpotifyProviderAppConfig{
					Key:         spotifyKey,
					Secret:      spotifySecret,
					RedirectURL: spotifyRedirectURL,
				},
				DatabaseURL: databaseURL,
			})
			if err != nil {
				return errors.WithStack(err)
			}

			if err := app.Run(inv.Context()); err != nil {
				return errors.WithStack(err)
			}

			return nil
		},
	}

	cmd.Options = serpent.OptionSet{
		{
			Flag:        "env",
			Env:         "LISFUN_ENV",
			Description: "Current environment lisfun is running: [dev production].",
			Value:       serpent.StringOf(&env),
			Default:     "dev",
		}, {
			Flag:        "port",
			Env:         "LISFUN_PORT",
			Description: "Specify a lisfun port number.",
			Value:       serpent.StringOf(&port),
		}, {
			Flag:        "log_level",
			Env:         "LISFUN_LOG_LEVEL",
			Description: "Log level of the global logger.",
			Value:       serpent.StringOf(&loglevel),
			Default:     "info",
		}, {
			Flag:        "spotify_key",
			Env:         "LISFUN_SPOTIFY_KEY",
			Description: "Spotify oauth app key.",
			Value:       serpent.StringOf(&spotifyKey),
			Required:    true,
		}, {
			Flag:        "spotify_secret",
			Env:         "LISFUN_SPOTIFY_SECRET",
			Description: "Spotify oauth app secret.",
			Value:       serpent.StringOf(&spotifySecret),
			Required:    true,
		}, {
			Flag:        "spotify_redirect_url",
			Env:         "LISFUN_SPOTIFY_REDIRECT_URL",
			Description: "Spotify oauth redirect url.",
			Value:       serpent.StringOf(&spotifyRedirectURL),
			Required:    true,
		}, {
			Flag:        "database_url",
			Env:         "LISFUN_DATABASE_URL",
			Description: "Main database url used by lisfun.",
			Value:       serpent.StringOf(&databaseURL),
			Required:    true,
		},
	}

	return cmd
}
