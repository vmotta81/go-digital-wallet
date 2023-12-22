package account_model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID
	Balance   int64
	CreatedAt time.Time
}

type LockedAccount struct {
	Id            uuid.UUID
	AccountId     uuid.UUID
	ProcessNumber string
}
