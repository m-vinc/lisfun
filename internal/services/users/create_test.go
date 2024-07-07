package users_test

import (
	"context"
	"lisfun/internal/services/models"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func FuzzCreateEmail(f *testing.F) {
	invalidEmailTestCase := []string{"toto@toto.fr", "", " ", "121", "!@#$%^&*(){}[]"}

	for _, ie := range invalidEmailTestCase {
		f.Add(ie)
	}

	f.Fuzz(func(t *testing.T, fuzzEmail string) {
		s := NewServiceMock(t)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		t.Log("testing", fuzzEmail)
		params := &models.UsersCreateParams{
			Username:       faker.Username(),
			Email:          fuzzEmail,
			ExternalUserID: faker.Username(),
		}

		_, err := s.Create(ctx, params)
		if err != nil {
			ves, ok := err.(validator.ValidationErrors)
			if !ok {
				t.Fatal("a validation error must be thrown while creating a user with invalid email", err)
			}

			for _, ve := range ves {
				if ve.Field() == "Email" {
					return
				}
			}
		}

		// t.Fatal("must throw a validation error while creating a user with invalid email", err)
	})
}

func TestCreateUser(t *testing.T) {
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

	_, err = s.Get(ctx, &models.UsersGetParams{
		Username: user.Username,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateWithDuplicateUsername(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	_, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	usernameDupParams := &models.UsersCreateParams{
		Username:       params.Username,
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	_, err = s.Create(ctx, usernameDupParams)
	if err != nil {
		if errors.Is(err, models.ErrUsernameAlreadyExist) {
			return
		}

		t.Fatal("ErrUsernameAlreadyExist must be thrown while creating a user with a taken username")
	}

	t.Fatal("no error was thrown while creating a user with a taken username")
}

func TestCreateWithDuplicateEmail(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	_, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	emailDupParams := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          params.Email,
		ExternalUserID: faker.Username(),
	}

	_, err = s.Create(ctx, emailDupParams)
	if err != nil {
		if errors.Is(err, models.ErrEmailAlreadyExist) {
			return
		}

		t.Fatal("ErrEmailAlreadyExist must be thrown while creating a user with a taken email")
	}

	t.Fatal("no error was thrown while creating a user with a taken email")
}

func TestCreateWithDuplicateExternalUserID(t *testing.T) {
	s := NewServiceMock(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	_, err := s.Create(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	externalUserIDdupParams := &models.UsersCreateParams{
		Username:       faker.Username(),
		Email:          faker.Email(),
		ExternalUserID: faker.Username(),
	}

	_, err = s.Create(ctx, externalUserIDdupParams)
	if err != nil {
		if errors.Is(err, models.ErrExternalUserIDAlreadyExist) {
			return
		}

		t.Fatal("ErrExternalUserIDAlreadyExist must be thrown while creating a user with a taken external user id")
	}

	t.Fatal("no error was thrown while creating a user with a taken external user id")
}
