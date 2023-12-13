package transaction_repository

import (
	transaction_model "digitalwallet-service/src/core/model/transaction"
	database "digitalwallet-service/src/data/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepository struct {
	db *pgxpool.Pool
}

func GetTransactionRepository() transactionRepository {
	connection := database.GetPgxPool()
	return transactionRepository{connection}
}

func (repository transactionRepository) Save(transaction transaction_model.Transaction) (*transaction_model.Transaction, error) {

	var err error
	if transaction.Id, err = uuid.NewRandom(); err != nil {
		return nil, err
	}

	_, err = database.Exec("insert into transactions (id, account_id, external_id, amount, type, status) values ($1, $2, $3, $4, $5, $6)",
		transaction.Id, transaction.AccountId, transaction.ExternalId, transaction.Amount, transaction.Type, transaction.Status)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repository transactionRepository) FindNewTransactionsByAccountId(accountId uuid.UUID) ([]transaction_model.Transaction, error) {
	rows, err := database.Query("select id, account_id, external_id, amount, type, status from transactions where status = 'NEW' and account_id = $1 order by created_at", accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []transaction_model.Transaction

	for rows.Next() {
		var tx transaction_model.Transaction
		if err := rows.Scan(&tx.Id, &tx.AccountId, &tx.ExternalId, &tx.Amount, &tx.Type, &tx.Status); err != nil {
			return nil, err
		}

		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func (repository transactionRepository) UpdateStatus(transactionId uuid.UUID, status transaction_model.TransactionStatus) error {
	_, err := database.Exec("update transactions set status = $2 where id = $1",
		transactionId, status)
	if err != nil {
		return err
	}

	return nil
}
