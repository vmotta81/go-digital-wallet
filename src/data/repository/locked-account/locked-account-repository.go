package locked_account_repository

import (
	"database/sql"
	account_model "digitalwallet-service/src/core/model/account"
	database "digitalwallet-service/src/data/repository"
)

type lockedAccountRepository struct {
	db *sql.DB
}

func GetLockedAccountRepository() lockedAccountRepository {
	connection := database.GetConnection()
	return lockedAccountRepository{connection}
}

func (repository lockedAccountRepository) Save(lockedAccount account_model.LockedAccount) (*account_model.LockedAccount, error) {
	_, err := database.ExecStatement("insert into locked_accounts (account_id, process_number) values ($1, $2)",
		lockedAccount.AccountId, lockedAccount.ProcessNumber)
	if err != nil {
		return nil, err
	}

	return &lockedAccount, nil
}

func (repository lockedAccountRepository) Remove(lockedAccount account_model.LockedAccount) error {
	_, err := database.ExecStatement("delete from locked_accounts where account_id = $1",
		lockedAccount.AccountId)
	if err != nil {
		return err
	}

	return nil
}
