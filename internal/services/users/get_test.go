package users_test

import (
	"context"
	"lisfun/internal/services/models"
	"lisfun/internal/types"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/pkg/errors"
)

func TestGetByID(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	createdUser, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	user, err := s.Get(ctx, &models.UsersGetParams{ID: &createdUser.ID})
	if err != nil {
		t.Fatal("cannot find the user", err)
	}

	if user == nil {
		t.Fatal("no user returned and no error returned")
	}

	if createdUser.ID.String() != user.ID.String() {
		t.Fatal("getted user id is not the right one", createdUser, user)
	}
}

func TestGetByUsername(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	createdUser, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	user, err := s.Get(ctx, &models.UsersGetParams{Username: &params.Username})
	if err != nil {
		t.Fatal("cannot find the user", err)
	}

	if user == nil {
		t.Fatal("no user returned and no error returned")
	}

	if *user.Username != *createdUser.Username {
		t.Fatal("getted user by username is not the right one", user, createdUser)
	}
}

func TestGetByUnknownUsername(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	user, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	user, err = s.Get(ctx, &models.UsersGetParams{Username: types.StringAddr(faker.Username())})
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return
		}

		t.Fatal("ErrUserNotFound must be thrown", err)
	}

	t.Fatal("no error has been thrown even if we search for a unknown user", user, err)
}

func TestGetByUsernameSelected(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	user, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	user, err = s.Get(ctx, &models.UsersGetParams{
		Username: &params.Username,
		Select:   []string{"id"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if user.ID.String() == "" || (user.Username != nil ||
		user.Email != nil ||
		user.ExternalUserID != nil ||
		user.CreatedAt != nil ||
		user.UpdatedAt != nil) {
		t.Fatal("only the id field is selected but some non-null value was found")
	}
}
