package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID
	Balance   int64
	createdAt time.Time
}
