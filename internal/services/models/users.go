package models

import (
	"lisfun/internal/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
)

// User = Model
type User struct {
	*db.User
}

func (u *User) HasCompleteRegistration() bool {
	return u.FirstName != nil && u.Username != nil
}

func (u *User) Wrap(dbuser *db.User) *User {
	u.User = dbuser

	return u
}

type Users []*User

func (us Users) Wrap(dbusers []*db.User) Users {
	us = make([]*User, len(dbusers))

	for i, dbuser := range dbusers {
		us[i] = (&User{}).Wrap(dbuser)
	}

	return us
}

// Users = service, Create = action, params = is a "dto" like struct
type UsersCreateParams struct {
	Username       string `json:"username" yaml:"username" validate:""`
	Email          string `json:"email" yaml:"email" validate:"email"`
	ExternalUserID string `json:"external_user_id" yaml:"external_user_id"`
}

var (
	ErrUserNotFound               = echo.NewHTTPError(http.StatusNotFound, "user not found")
	ErrUsernameAlreadyExist       = echo.NewHTTPError(http.StatusConflict, "username already exist")
	ErrEmailAlreadyExist          = echo.NewHTTPError(http.StatusConflict, "username already exist")
	ErrExternalUserIDAlreadyExist = echo.NewHTTPError(http.StatusConflict, "username already exist")
)

type UsersGetParams struct {
	ID             *uuid.UUID `json:"id" yaml:"id" validate:"omitempty,uuid"`
	Username       *string    `json:"username" yaml:"username" validate:"omitempty"`
	ExternalUserID *string    `json:"external_user_id" yaml:"external_user_id" validate:"omitempty"`

	Select []string `json:"select" yaml:"select" validate:""`
}

type UsersOnboardParams struct {
	GothUser *goth.User
}

type UsersCompleteRegistrationParams struct {
	User *User `json:"-" yaml:"-" form:"-" validate:"required"`

	Username  string `json:"username" yaml:"username" validate:"required"`
	Firstname string `json:"firstname" yaml:"firstname" validate:"required"`
}
