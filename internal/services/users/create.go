package users

import (
	"context"
	"lisfun/internal/db"
	dbuser "lisfun/internal/db/user"
	"lisfun/internal/services/models"

	"github.com/pkg/errors"
)

func (s *Service) Create(ctx context.Context, params *models.UsersCreateParams) (*models.User, error) {
	err := s.validate.StructCtx(ctx, params)
	if err != nil {
		return nil, err
	}

	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	err = s.unique(tx.Queryable(ctx), params.Username, params.Email, params.ExternalUserID)
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	u, err := tx.User.Create().
		SetUsername(params.Username).
		SetEmail(params.Email).
		SetExternalUserID(params.ExternalUserID).
		Save(ctx)
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	return (&models.User{}).Wrap(u), tx.Release(ctx, nil)
}

func (s *Service) unique(q db.Queryable, username, email, externalUserID string) error {
	err := s.uniqueUsername(q, username)
	if err != nil {
		return err
	}

	err = s.uniqueEmail(q, email)
	if err != nil {
		return err
	}

	err = s.uniqueExternalUserID(q, externalUserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) uniqueEmail(q db.Queryable, email string) error {
	ok, err := q.UserClient().Query().
		Where(
			dbuser.Or(
				dbuser.EmailEQ(email),
			),
		).Exist(q.Context())
	if err != nil {
		return err
	}

	if ok {
		return errors.WithStack(models.ErrEmailAlreadyExist)
	}

	return nil
}

func (s *Service) uniqueUsername(q db.Queryable, username string) error {
	if username == "" {
		return nil
	}

	ok, err := q.UserClient().Query().
		Where(
			dbuser.Or(
				dbuser.UsernameEQ(username),
			),
		).Exist(q.Context())
	if err != nil {
		return err
	}

	if ok {
		return errors.WithStack(models.ErrUsernameAlreadyExist)
	}

	return nil
}

func (s *Service) uniqueExternalUserID(q db.Queryable, externalUserID string) error {
	ok, err := q.UserClient().Query().
		Where(
			dbuser.Or(
				dbuser.ExternalUserIDEQ(externalUserID),
			),
		).Exist(q.Context())
	if err != nil {
		return err
	}

	if ok {
		return errors.WithStack(models.ErrExternalUserIDAlreadyExist)
	}

	return nil
}
