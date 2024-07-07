package app

import (
	"lisfun/internal/services/spotify"
	"lisfun/internal/services/tokens"
	"lisfun/internal/services/users"
)

func (app *App) Services() error {
	var err error

	app.usersService, err = users.New(app.database, app.logger.With().Str("service", "usersService").Logger())
	if err != nil {
		return err
	}

	app.tokensService, err = tokens.New(app.database, app.logger.With().Str("service", "tokensService").Logger())
	if err != nil {
		return err
	}

	app.spotifyService, err = spotify.New(app.tokensService, app.database, app.config.SpotifyProvider, app.logger.With().Str("service", "spotifyService").Logger())
	if err != nil {
		return err
	}

	return nil
}
