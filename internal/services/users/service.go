package users

import (
	"lisfun/internal/db"
	"lisfun/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type Service struct {
	db *db.Client

	logger   zerolog.Logger
	validate *validator.Validate
}

func New(db *db.Client, logger zerolog.Logger) (*Service, error) {
	svc := &Service{
		db:       db,
		logger:   logger,
		validate: types.Validate,
	}

	return svc, nil
}
