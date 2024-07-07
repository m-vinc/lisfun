package spotify

import (
	"lisfun/internal/app/models"
	"lisfun/internal/db"
	"lisfun/internal/services/tokens"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type Service struct {
	db              *db.Client
	tokensService   *tokens.Service
	spotifyClientID string

	validate *validator.Validate
	logger   zerolog.Logger
}

func New(tokensService *tokens.Service, db *db.Client, cfg *models.SpotifyProviderAppConfig, logger zerolog.Logger) (*Service, error) {
	svc := &Service{
		db:              db,
		spotifyClientID: cfg.Key,
		tokensService:   tokensService,
		validate:        validator.New(),
		logger:          logger,
	}

	return svc, nil
}
