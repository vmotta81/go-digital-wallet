package account_repository

import (
	"database/sql"
	database "digitalwallet-service/src/data/repository"

	"github.com/google/uuid"
)

type accountRepository struct {
	database *sql.DB
}

func GetAccountRepository() accountRepository {
	connection := database.GetConnection()
	return accountRepository{connection}
}

func (repository accountRepository) Create() (*uuid.UUID, error) {
	accountId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = database.ExecStatement("insert into accounts (id, balance) values ($1, 0)", accountId)
	if err != nil {
		return nil, err
	}

	return &accountId, nil
}
