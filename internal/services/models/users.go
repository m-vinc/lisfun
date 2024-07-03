package models

import "lisfun/internal/db"

type User struct {
	*db.User
}

// Users = service, Create = action, params = is a "dto" like struct
type UsersCreateParams struct {
	Username       string `json:"username" yaml:"username" validate:""`
	Email          string `json:"email" yaml:"email" validate:""`
	ExternalUserID int    `json:"external_user_id" yaml:"external_user_id"`
}
