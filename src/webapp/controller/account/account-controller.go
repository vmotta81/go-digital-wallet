package account_controller

import (
	"digitalwallet-service/src/core/model"
	account_usecase "digitalwallet-service/src/core/usecase/account"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	accountId, err := account_usecase.Create()
	if err != nil {
		createError(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
	} else {
		account := model.Account{
			Id: *accountId,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}

func createError(w http.ResponseWriter, message string, errorCode int) {

	var mapMessage = map[string]string{
		"message": message,
	}

	response, err := json.Marshal(mapMessage)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "%s"}`, message), errorCode)
	}

	http.Error(w, string(response), errorCode)
}
