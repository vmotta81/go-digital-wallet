package account_controller

import (
	account_model "digitalwallet-service/src/core/model/account"
	transaction_model "digitalwallet-service/src/core/model/transaction"
	account_usecase "digitalwallet-service/src/core/usecase/account"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {

	accountId, err := account_usecase.Create()
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
	} else {
		account := account_model.Account{
			Id: *accountId,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}

func Cashin(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	var transaction transaction_model.Transaction

	if err := json.Unmarshal(body, &transaction); err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	accountId, err := uuid.Parse(params["account-id"])
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}
	transaction.AccountId = accountId

	savedTransaction, err := account_usecase.Cashin(transaction)
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedTransaction)
}

func Cashout(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	var transaction transaction_model.Transaction

	if err := json.Unmarshal(body, &transaction); err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	accountId, err := uuid.Parse(params["account-id"])
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}
	transaction.AccountId = accountId

	savedTransaction, err := account_usecase.Cashout(transaction)
	if err != nil {
		createErrorResponseByError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedTransaction)
}

func createErrorResponseByError(w http.ResponseWriter, err error, errorCode int) {
	createErrorResponse(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
}

func createErrorResponse(w http.ResponseWriter, message string, errorCode int) {

	var mapMessage = map[string]string{
		"message": message,
	}

	response, err := json.Marshal(mapMessage)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "%s"}`, message), errorCode)
	}

	http.Error(w, string(response), errorCode)
}
