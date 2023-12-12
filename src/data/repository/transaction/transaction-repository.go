package transaction_repository

import (
	"database/sql"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	database "digitalwallet-service/src/data/repository"

	"github.com/google/uuid"
)

type transactionRepository struct {
	db *sql.DB
}

func GetTransactionRepository() transactionRepository {
	connection := database.GetConnection()
	return transactionRepository{connection}
}

func (repository transactionRepository) Save(transaction transaction_model.Transaction) (*transaction_model.Transaction, error) {

	var err error
	if transaction.Id, err = uuid.NewRandom(); err != nil {
		return nil, err
	}

	_, err = database.ExecStatement("insert into transactions (id, account_id, external_id, amount, type, status) values ($1, $2, $3, $4, $5, $6)",
		transaction.Id, transaction.AccountId, transaction.ExternalId, transaction.Amount, transaction.Type, transaction.Status)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
