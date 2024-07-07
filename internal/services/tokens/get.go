package tokens

import (
	"context"
	"lisfun/internal/db"
	dbtoken "lisfun/internal/db/token"
	dbuser "lisfun/internal/db/user"
	"lisfun/internal/services/models"
)

func (s *Service) Get(ctx context.Context, user *models.User) (*models.Token, error) {
	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	token, err := tx.Token.Query().Where(
		dbtoken.HasOwnerWith(dbuser.IDEQ(user.ID)),
	).Order(db.Desc(dbtoken.FieldID)).
		First(ctx)

	notfound := db.IsNotFound(err)
	if err != nil && !notfound {
		return nil, tx.Release(ctx, err)
	}

	if notfound {
		return nil, tx.Release(ctx, models.ErrTokenNotFound)
	}

	return (&models.Token{}).Wrap(token), tx.Release(ctx, nil)
}
