package account_usecase

import (
	account_repository "digitalwallet-service/src/data/repository/account"

	"github.com/google/uuid"
)

func Create() (*uuid.UUID, error) {

	repository := account_repository.GetAccountRepository()

	accountId, err := repository.Create()
	if err != nil {
		return nil, err
	}

	return accountId, nil
}
