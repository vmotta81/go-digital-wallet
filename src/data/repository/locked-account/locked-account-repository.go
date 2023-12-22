package locked_account_repository

import (
	account_model "digitalwallet-service/src/core/model/account"
	database "digitalwallet-service/src/data/repository"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type lockedAccountRepository struct {
	db *pgxpool.Pool
}

func GetLockedAccountRepository() lockedAccountRepository {
	connection := database.GetPgxPool()
	return lockedAccountRepository{connection}
}

func (repository lockedAccountRepository) Save(lockedAccount account_model.LockedAccount) (*account_model.LockedAccount, error) {
	_, err := database.Exec("insert into locked_accounts (account_id, process_number) values ($1, $2)",
		lockedAccount.AccountId, lockedAccount.ProcessNumber)
	if err != nil {
		return nil, err
	}

	return &lockedAccount, nil
}

func (repository lockedAccountRepository) Remove(lockedAccount account_model.LockedAccount) error {
	_, err := database.Exec("delete from locked_accounts where account_id = $1",
		lockedAccount.AccountId)
	if err != nil {
		return err
	}

	fmt.Println("Finish process: ", time.Now())

	return nil
}
