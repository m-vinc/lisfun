package models

import (
	"lisfun/internal/db"
	"time"

	"golang.org/x/oauth2"
)

type Token struct {
	*db.Token
}

func (t *Token) Wrap(dbtoken *db.Token) *Token {
	t.Token = dbtoken

	return t
}

type Tokens []*Token

func (ts Tokens) Wrap(dbtokens []*db.Token) Tokens {
	ts = make([]*Token, len(dbtokens))

	for i, dbtoken := range dbtokens {
		ts[i] = (&Token{}).Wrap(dbtoken)
	}

	return ts
}

type TokensCreateParams struct {
	User *User

	AccessToken  string
	RefreshToken string
	ExpireAt     time.Time
}

type TokensRefreshTokenParams struct {
	NewToken *oauth2.Token
	Token    *Token
}
