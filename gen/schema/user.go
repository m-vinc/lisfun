package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("username").Nillable().Optional(),
		field.String("first_name").Nillable().Optional(),
		field.String("email").Nillable(),
		field.String("external_user_id").Nillable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Nillable(),
		field.Time("updated_at").
			UpdateDefault(time.Now).
			Nillable().
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", Token.Type),
	}
}
