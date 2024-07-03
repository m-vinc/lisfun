// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "expire_at", Type: field.TypeTime},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "external_user_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserTokensColumns holds the columns for the "user_tokens" table.
	UserTokensColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "token_id", Type: field.TypeUUID},
	}
	// UserTokensTable holds the schema information for the "user_tokens" table.
	UserTokensTable = &schema.Table{
		Name:       "user_tokens",
		Columns:    UserTokensColumns,
		PrimaryKey: []*schema.Column{UserTokensColumns[0], UserTokensColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_tokens_user_id",
				Columns:    []*schema.Column{UserTokensColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_tokens_token_id",
				Columns:    []*schema.Column{UserTokensColumns[1]},
				RefColumns: []*schema.Column{TokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TokensTable,
		UsersTable,
		UserTokensTable,
	}
)

func init() {
	UserTokensTable.ForeignKeys[0].RefTable = UsersTable
	UserTokensTable.ForeignKeys[1].RefTable = TokensTable
}