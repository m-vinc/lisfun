package auth

import (
	"net/url"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/spotify"
)

type SpotifyProviderConfig struct {
	Key         string
	Secret      string
	RedirectURL string
}

func spotifyProvider(cfg *SpotifyProviderConfig) error {
	_, err := url.Parse(cfg.RedirectURL)
	if err != nil {
		return err
	}

	spotifyProvider := spotify.New(cfg.Key, cfg.Secret, cfg.RedirectURL)
	goth.UseProviders(spotifyProvider)

	return nil
}
