package users

import (
	"context"
	"lisfun/internal/services/models"

	"github.com/pkg/errors"
)

func (s *Service) Onboard(ctx context.Context, params *models.UsersOnboardParams) (*models.User, error) {
	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.Get(ctx, &models.UsersGetParams{
		ExternalUserID: &params.GothUser.UserID,
	})

	notfound := errors.Is(err, models.ErrUserNotFound)
	if err != nil && !notfound {
		return nil, tx.Release(ctx, err)
	}

	if user == nil {
		user, err = s.Create(ctx, &models.UsersCreateParams{
			Username:       params.GothUser.NickName,
			Email:          params.GothUser.Email,
			ExternalUserID: params.GothUser.UserID,
		})
		if err != nil {
			return nil, tx.Release(ctx, err)
		}
	}

	return user, tx.Release(ctx, nil)
}
