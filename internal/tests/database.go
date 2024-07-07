package tests

import (
	"context"
	"lisfun/internal/db"

	_ "github.com/mattn/go-sqlite3"
)

func Database(ctx context.Context) (*db.Client, error) {
	db, err := db.Open("sqlite3", "file:ent?mode=memory&_fk=1")
	if err != nil {
		return nil, err
	}

	err = db.Schema.Create(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
