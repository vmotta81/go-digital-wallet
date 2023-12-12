package account_usecase

import (
	account_model "digitalwallet-service/src/core/model/account"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	locked_account_repository "digitalwallet-service/src/data/repository/locked-account"
	transaction_repository "digitalwallet-service/src/data/repository/transaction"
	"math/rand"
)

func Cashin(transaction transaction_model.Transaction) (*transaction_model.Transaction, error) {

	transaction.Type = transaction_model.Credit
	transaction.Status = transaction_model.New

	repository := transaction_repository.GetTransactionRepository()

	savedTransaction, err := repository.Save(transaction)
	if err != nil {
		return nil, err
	}

	processCashin(*savedTransaction)

	return savedTransaction, nil
}

func processCashin(transaction transaction_model.Transaction) {
	processNumber := rand.Intn(999999999)
	lockedAccount := account_model.LockedAccount{
		Id:            transaction.AccountId,
		AccountId:     transaction.AccountId,
		ProcessNumber: string(processNumber),
	}

	repository := locked_account_repository.GetLockedAccountRepository()
	savedLockedAccount, err := repository.Save(lockedAccount)
	if err != nil {
		return
	}

	repository.Remove(*savedLockedAccount)
}
