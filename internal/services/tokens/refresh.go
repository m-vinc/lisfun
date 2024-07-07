package tokens

import (
	"context"
	dbtoken "lisfun/internal/db/token"
	"lisfun/internal/services/models"
)

func (s *Service) RefreshToken(ctx context.Context, params *models.TokensRefreshTokenParams) error {
	if params.Token.AccessToken == params.NewToken.AccessToken {
		return nil
	}

	tx, ctx, err := s.db.Acquire(ctx)
	if err != nil {
		return err
	}

	dbToken, err := tx.Token.Query().
		Where(dbtoken.IDEQ(params.Token.ID)).
		ForUpdate().
		First(ctx)
	if err != nil {
		return tx.Release(ctx, err)
	}

	_, err = dbToken.Update().
		SetAccessToken(params.NewToken.AccessToken).
		SetRefreshToken(params.NewToken.RefreshToken).
		SetExpireAt(params.NewToken.Expiry).
		Save(ctx)
	if err != nil {
		return tx.Release(ctx, err)
	}

	return tx.Release(ctx, nil)
}
