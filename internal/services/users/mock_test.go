package users_test

import (
	"context"
	"lisfun/internal/services/users"
	"lisfun/internal/tests"

	"github.com/rs/zerolog"
)

type ServiceMock struct {
	*users.Service
}

func NewServiceMock(tc tests.TestingContext) *ServiceMock {
	ctx := context.Background()
	db, err := tests.Database(ctx)
	if err != nil {
		tc.Fatal(err)
	}

	logger := zerolog.New(zerolog.NewTestWriter(tc))

	s, err := users.New(db, logger)
	if err != nil {
		tc.Fatal(err)
	}

	return &ServiceMock{Service: s}
}
