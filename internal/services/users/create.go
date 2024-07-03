package users

import (
	"context"
	"lisfun/internal/services/models"
)

func (s *Service) Create(ctx context.Context, params *models.UsersCreateParams) (*models.User, error) {
	err := s.validate.StructCtx(ctx, params)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
