package tests

import "github.com/rs/zerolog"

type TestingContext interface {
	zerolog.TestingLog

	Fatal(args ...any)
}
