package server

import (
	"UserService/services"
	"log"
	"net/http"

	"github.com/sarulabs/di"
)

var DiMiddleware = func(next http.Handler) http.Handler {
	return di.HTTPMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)
	}), services.Provider, func(msg string) {
		log.Fatal(msg)
	})

}
