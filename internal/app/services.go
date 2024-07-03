package app

import "lisfun/internal/services/users"

func (app *App) Services() error {
	var err error

	app.usersService, err = users.New(app.database, app.logger.With().Str("service", "usersService").Logger())
	if err != nil {
		return err
	}

	return nil
}
