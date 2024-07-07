package tokens

import (
	"context"
	"lisfun/internal/db"
	dbtoken "lisfun/internal/db/token"
	dbuser "lisfun/internal/db/user"
	"lisfun/internal/services/models"
)

func (s *Service) Create(ctx context.Context, params *models.TokensCreateParams) (*models.Token, error) {
	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	dbToken, err := tx.Token.Query().Where(
		dbtoken.HasOwnerWith(dbuser.IDEQ(params.User.ID)),
	).ForUpdate().
		First(ctx)
	notfound := db.IsNotFound(err)

	if err != nil && !notfound {
		return nil, tx.Release(ctx, err)
	}

	if notfound {
		token, err := tx.Token.Create().
			SetOwner(params.User.User).
			SetAccessToken(params.AccessToken).
			SetRefreshToken(params.RefreshToken).
			SetExpireAt(params.ExpireAt).
			Save(ctx)
		if err != nil {
			return nil, tx.Release(ctx, err)
		}

		return (&models.Token{}).Wrap(token), tx.Release(ctx, nil)
	}

	dbToken, err = dbToken.Update().
		SetAccessToken(params.AccessToken).
		SetRefreshToken(params.RefreshToken).
		SetExpireAt(params.ExpireAt).
		Save(ctx)
	if err != nil {
		return nil, tx.Release(ctx, err)
	}

	return (&models.Token{}).Wrap(dbToken), tx.Release(ctx, nil)
}
