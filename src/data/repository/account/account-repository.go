package account_repository

import (
	account_model "digitalwallet-service/src/core/model/account"
	database "digitalwallet-service/src/data/repository"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	database *pgxpool.Pool
}

func GetAccountRepository() accountRepository {
	connection := database.GetPgxPool()
	return accountRepository{connection}
}

func (repository accountRepository) Create() (*uuid.UUID, error) {
	accountId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = database.Exec("insert into accounts (id, balance) values ($1, 0)", accountId)
	if err != nil {
		return nil, err
	}

	return &accountId, nil
}

func (repository accountRepository) FindById(accountId uuid.UUID) (*account_model.Account, error) {
	rows, err := database.Query("select id, balance, created_at from accounts where id = $1", accountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var account account_model.Account
		if err := rows.Scan(&account.Id, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		return &account, nil
	}

	return nil, fmt.Errorf("Account not found")
}

func (repository accountRepository) SumBalance(accountId uuid.UUID, amount int64) error {
	_, err := database.Exec("update accounts set balance = (balance + $2) where id = $1",
		accountId, amount)
	if err != nil {
		return err
	}

	return nil
}

func (repository accountRepository) SubBalance(accountId uuid.UUID, amount int64) error {
	_, err := database.Exec("update accounts set balance = (balance - $2) where id = $1",
		accountId, amount)
	if err != nil {
		return err
	}

	return nil
}
