package users

import (
	"context"
	"lisfun/internal/services/models"
)

func (s *Service) CompleteRegistration(ctx context.Context, params *models.UsersCompleteRegistrationParams) (*models.User, error) {
	err := s.validate.StructCtx(ctx, params)
	if err != nil {
		return nil, err
	}

	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	err = s.uniqueUsername(tx.Queryable(ctx), params.Username)
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	user, err := s.Get(ctx, &models.UsersGetParams{ID: &params.User.ID})
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	updatedUser, err := user.Update().
		SetFirstName(params.Firstname).
		SetUsername(params.Username).
		Save(ctx)
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	return (&models.User{}).Wrap(updatedUser), tx.Release(ctx, nil)
}
