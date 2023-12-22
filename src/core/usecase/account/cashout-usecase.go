package account_usecase

import (
	transaction_model "digitalwallet-service/src/core/model/transaction"
	transaction_repository "digitalwallet-service/src/data/repository/transaction"
)

func Cashout(transaction transaction_model.Transaction) (*transaction_model.Transaction, error) {

	transaction.Type = transaction_model.Debit
	transaction.Status = transaction_model.New

	repository := transaction_repository.GetTransactionRepository()

	savedTransaction, err := repository.Save(transaction)
	if err != nil {
		return nil, err
	}

	go processTransaction(*savedTransaction)

	return savedTransaction, nil
}
