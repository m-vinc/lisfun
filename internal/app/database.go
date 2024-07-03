package app

import (
	"context"
	"lisfun/internal/db"
	"time"

	_ "github.com/lib/pq"
)

func (app *App) Database() error {
	psql, err := db.Open("postgres", app.config.DatabaseURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err = psql.ExecContext(ctx, "SELECT 1;")
	if err != nil {
		return err
	}

	app.database = psql

	return nil
}
