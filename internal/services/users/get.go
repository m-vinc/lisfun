package users

import (
	"context"
	"lisfun/internal/db"
	dbuser "lisfun/internal/db/user"
	"lisfun/internal/services/models"

	"github.com/pkg/errors"
)

func (s *Service) Get(ctx context.Context, params *models.UsersGetParams) (*models.User, error) {
	err := s.validate.StructCtx(ctx, params)
	if err != nil {
		return nil, err
	}

	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	q := tx.User.Query()

	if params.ID != nil {
		q.Where(dbuser.IDEQ(*params.ID))
	}

	if params.Username != nil {
		q.Where(dbuser.UsernameEQ(*params.Username))
	}

	if params.ExternalUserID != nil {
		q.Where(dbuser.ExternalUserID(*params.ExternalUserID))
	}

	if params.Select != nil {
		q.Select(params.Select...)
	}

	u, err := q.First(ctx)
	notfound := db.IsNotFound(err)
	if err != nil && !notfound {
		return nil, tx.Release(ctx, err)
	}

	if notfound {
		return nil, tx.Release(ctx, errors.WithStack(models.ErrUserNotFound))
	}

	return (&models.User{}).Wrap(u), tx.Release(ctx, nil)
}
