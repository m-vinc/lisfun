package spotify

import (
	"context"
	"lisfun/internal/services/models"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func (s *Service) Client(ctx context.Context, user *models.User) (*spotify.Client, error) {
	dbToken, err := s.tokensService.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	oauthToken := &oauth2.Token{AccessToken: dbToken.AccessToken, TokenType: "Bearer", RefreshToken: dbToken.RefreshToken, Expiry: dbToken.ExpireAt}

	auth := spotifyauth.New(spotifyauth.WithClientID(s.spotifyClientID))

	oauthToken, err = auth.RefreshToken(ctx, oauthToken)
	if err != nil {
		return nil, err
	}

	err = s.tokensService.RefreshToken(ctx, &models.TokensRefreshTokenParams{
		NewToken: oauthToken,
		Token:    dbToken,
	})
	if err != nil {
		return nil, err
	}

	spotifyClient := spotify.New(auth.Client(ctx, oauthToken))
	if err != nil {
		return nil, err
	}

	return spotifyClient, err
}
