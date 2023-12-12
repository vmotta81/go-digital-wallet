package account_usecase

import (
	account_model "digitalwallet-service/src/core/model/account"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	account_repository "digitalwallet-service/src/data/repository/account"
	locked_account_repository "digitalwallet-service/src/data/repository/locked-account"
	transaction_repository "digitalwallet-service/src/data/repository/transaction"
	"math/rand"

	"github.com/google/uuid"
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
		ProcessNumber: string(rune(processNumber)),
	}

	repository := locked_account_repository.GetLockedAccountRepository()
	savedLockedAccount, err := repository.Save(lockedAccount)
	if err != nil {
		return
	}

	channel := make(chan transaction_model.Transaction)

	go getTransactions(lockedAccount.AccountId, channel)

	for {
		txMessage, open := <-channel
		if !open {
			break
		}

		processTransaction(txMessage)
	}

	repository.Remove(*savedLockedAccount)
}

func getTransactions(accountId uuid.UUID, channel chan transaction_model.Transaction) {

	repository := transaction_repository.GetTransactionRepository()

	for {
		transactions, err := repository.FindNewTransactionsByAccountId(accountId)
		if err != nil {
			break
		}
		if len(transactions) == 0 {
			break
		}
		for _, tx := range transactions {
			channel <- tx
		}
	}

	close(channel)
}

func processTransaction(transaction transaction_model.Transaction) {

	transactionRepository := transaction_repository.GetTransactionRepository()

	transaction.Status = transaction_model.Processed

	if err := transactionRepository.UpdateStatus(transaction.Id, transaction.Status); err != nil {
		return
	}

	accountRepository := account_repository.GetAccountRepository()

	if transaction.Type == transaction_model.Credit {
		if err := accountRepository.SumBalance(transaction.AccountId, transaction.Amount); err != nil {
			return
		}
	} else if transaction.Type == transaction_model.Debit {
		if err := accountRepository.SubBalance(transaction.AccountId, transaction.Amount); err != nil {
			return
		}
	}

}
