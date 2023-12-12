package account_repository

import (
	"database/sql"
	database "digitalwallet-service/src/data/repository"
	"fmt"

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

	statement, err := repository.database.Prepare("insert into accounts (id, balance) values ($1, 0)")
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(accountId)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("account was not created")
	}

	return &accountId, nil
}
