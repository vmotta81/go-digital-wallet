package transaction_model

import "github.com/google/uuid"

type TransactionType string

const (
	Debit    TransactionType = "DEBIT"
	Credit   TransactionType = "CREDIT"
	Transfer TransactionType = "TRANSFER"
)

type TransactionStatus string

const (
	New       TransactionStatus = "NEW"
	Processed TransactionStatus = "PROCESSED"
)

type Transaction struct {
	Id         uuid.UUID
	AccountId  uuid.UUID
	ExternalId string
	Amount     int64
	Type       TransactionType
	Status     TransactionStatus
}
