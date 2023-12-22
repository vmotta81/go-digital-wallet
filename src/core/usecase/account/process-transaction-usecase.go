package account_usecase

import (
	account_model "digitalwallet-service/src/core/model/account"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	account_repository "digitalwallet-service/src/data/repository/account"
	locked_account_repository "digitalwallet-service/src/data/repository/locked-account"
	transaction_repository "digitalwallet-service/src/data/repository/transaction"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func processTransaction(transaction transaction_model.Transaction) {
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
	defer repository.Remove(*savedLockedAccount)

	channel := make(chan transaction_model.Transaction)

	go getTransactions(lockedAccount.AccountId, channel, nil)

	for {
		txMessage, open := <-channel
		if !open {
			break
		}

		processTransactionMessage(txMessage)
	}
}

func getTransactions(accountId uuid.UUID, channel chan transaction_model.Transaction, startDate *time.Time) {

	repository := transaction_repository.GetTransactionRepository()

	for {
		transactions, err := repository.FindNewTransactionsByAccountId(accountId, startDate)
		if err == nil {
			for _, tx := range transactions {
				channel <- tx
				startDate = &tx.CreatedAt
			}
		}
		if len(transactions) == 0 {
			break
		}
	}

	close(channel)
}

func processTransactionMessage(transaction transaction_model.Transaction) {
	var status *transaction_model.TransactionStatus = new(transaction_model.TransactionStatus)
	*status = transaction_model.Processed

	var reason *string = new(string)

	defer updateTransactionsStatus(transaction.Id, status, reason)

	accountRepository := account_repository.GetAccountRepository()

	if transaction.Type == transaction_model.Credit {
		if err := accountRepository.SumBalance(transaction.AccountId, transaction.Amount); err != nil {
			*reason = fmt.Sprintf("SumBalance Error: %s", err)
			*status = transaction_model.Failed
			return
		}
	} else if transaction.Type == transaction_model.Debit {
		account, err := accountRepository.FindById(transaction.AccountId)
		if err != nil {
			*reason = fmt.Sprintf("Account::FindById Error: %s", err)
			*status = transaction_model.Failed
			return
		}

		if (account.Balance - transaction.Amount) >= 0 {
			if err := accountRepository.SubBalance(transaction.AccountId, transaction.Amount); err != nil {
				*reason = fmt.Sprintf("SubBalance Error: %s", err)
				*status = transaction_model.Failed
				return
			}
		} else {
			*reason = fmt.Sprintf("No Funds Error - balance: %d", account.Balance)
			*status = transaction_model.Failed
		}
	}

}

func updateTransactionsStatus(transactionId uuid.UUID, status *transaction_model.TransactionStatus, reason *string) {
	transactionRepository := transaction_repository.GetTransactionRepository()
	if err := transactionRepository.UpdateStatus(transactionId, *status, *reason); err != nil {
		return
	}
}
