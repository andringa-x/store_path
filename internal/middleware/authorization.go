package middleware

import (
	"errors"
	"net/http"

	"github.com/andringa-x/store_path/api"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string = r.Header.Get("Authorization")

		if token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		if token != "test" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
