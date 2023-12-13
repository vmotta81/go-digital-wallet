package account_repository

import (
	database "digitalwallet-service/src/data/repository"

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
