package tokens

import (
	"lisfun/internal/db"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type Service struct {
	db *db.Client

	validate *validator.Validate
	logger   zerolog.Logger
}

func New(db *db.Client, logger zerolog.Logger) (*Service, error) {
	svc := &Service{
		db:       db,
		validate: validator.New(),
		logger:   logger,
	}

	return svc, nil
}
