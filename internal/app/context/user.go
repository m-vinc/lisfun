package context

import (
	"encoding/gob"
	"lisfun/internal/services/models"

	"github.com/google/uuid"
)

func init() {
	gob.Register(UserContext{})
}

type UserContext struct {
	ID uuid.UUID

	User *models.User
}
