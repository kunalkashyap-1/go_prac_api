package api

import (
	"encoding/json"
	"net/http"
)

// coin balance prams
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponseBalance struct {
	Code    int   //status code
	Balance int64 //Account balance
}

type Error struct {
	Code    int    //Error code
	Message string //Error message
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error occurred", http.StatusInternalServerError)
	}
)
