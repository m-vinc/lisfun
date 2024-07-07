// Code generated by ent, DO NOT EDIT.

package token

import (
	"lisfun/internal/db/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldID, id))
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldAccessToken, v))
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldRefreshToken, v))
}

// ExpireAt applies equality check predicate on the "expire_at" field. It's identical to ExpireAtEQ.
func ExpireAt(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldExpireAt, v))
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldAccessToken, v))
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldAccessToken, v))
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldAccessToken, vs...))
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldAccessToken, vs...))
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldAccessToken, v))
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldAccessToken, v))
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldAccessToken, v))
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldAccessToken, v))
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldAccessToken, v))
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldAccessToken, v))
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldAccessToken, v))
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldAccessToken, v))
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldAccessToken, v))
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldRefreshToken, v))
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldRefreshToken, v))
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldRefreshToken, vs...))
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldRefreshToken, vs...))
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldRefreshToken, v))
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldRefreshToken, v))
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldRefreshToken, v))
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldRefreshToken, v))
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldRefreshToken, v))
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldRefreshToken, v))
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldRefreshToken, v))
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldRefreshToken, v))
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldRefreshToken, v))
}

// ExpireAtEQ applies the EQ predicate on the "expire_at" field.
func ExpireAtEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldExpireAt, v))
}

// ExpireAtNEQ applies the NEQ predicate on the "expire_at" field.
func ExpireAtNEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldExpireAt, v))
}

// ExpireAtIn applies the In predicate on the "expire_at" field.
func ExpireAtIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldExpireAt, vs...))
}

// ExpireAtNotIn applies the NotIn predicate on the "expire_at" field.
func ExpireAtNotIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldExpireAt, vs...))
}

// ExpireAtGT applies the GT predicate on the "expire_at" field.
func ExpireAtGT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldExpireAt, v))
}

// ExpireAtGTE applies the GTE predicate on the "expire_at" field.
func ExpireAtGTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldExpireAt, v))
}

// ExpireAtLT applies the LT predicate on the "expire_at" field.
func ExpireAtLT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldExpireAt, v))
}

// ExpireAtLTE applies the LTE predicate on the "expire_at" field.
func ExpireAtLTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldExpireAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(sql.NotPredicates(p))
}
