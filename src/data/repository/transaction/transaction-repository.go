package transaction_repository

import (
	"database/sql"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	database "digitalwallet-service/src/data/repository"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	var externalId *string = nil

	if transaction.ExternalId != "" {
		*externalId = transaction.ExternalId
	}

	_, err = database.Exec("insert into transactions (id, account_id, external_id, amount, type, status, reason) values ($1, $2, $3, $4, $5, $6, $7)",
		transaction.Id, transaction.AccountId, externalId, transaction.Amount, transaction.Type, transaction.Status, transaction.Reason)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repository transactionRepository) FindNewTransactionsByAccountId(accountId uuid.UUID, startDate *time.Time) ([]transaction_model.Transaction, error) {

	var (
		rows pgx.Rows
		err  error
	)

	if startDate == nil {
		rows, err = database.Query("select id, account_id, external_id, amount, type, status, reason, created_at from transactions where status = 'NEW' and account_id = $1 order by created_at", accountId)
	} else {
		rows, err = database.Query("select id, account_id, external_id, amount, type, status, reason, created_at from transactions where status = 'NEW' and account_id = $1 and created_at > $2 order by created_at", accountId, startDate)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []transaction_model.Transaction

	var externalIdSql sql.NullString
	var reasonSql sql.NullString

	for rows.Next() {
		var tx transaction_model.Transaction
		if err := rows.Scan(&tx.Id, &tx.AccountId, &externalIdSql, &tx.Amount, &tx.Type, &tx.Status, &reasonSql, &tx.CreatedAt); err != nil {
			return nil, err
		}

		if externalIdSql.Valid {
			tx.ExternalId = externalIdSql.String
		}
		if reasonSql.Valid {
			tx.Reason = reasonSql.String
		}

		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func (repository transactionRepository) UpdateStatus(transactionId uuid.UUID, status transaction_model.TransactionStatus, reason string) error {
	_, err := database.Exec("update transactions set status = $2, reason = $3 where id = $1",
		transactionId, status, reason)
	if err != nil {
		return err
	}

	return nil
}
