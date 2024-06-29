package middleware

import (
	"errors"
	"net/http"

	"github.com/kunalkashyap-1/go_prac_api/api"
	"github.com/kunalkashyap-1/go_prac_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid Username and token")

func Autherization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Autherization")
		var err error

		if username == "" || token == "" {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
		}

		next.ServeHTTP(w, r)
	})
}
