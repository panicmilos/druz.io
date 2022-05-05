package server

import (
	"log"
	"net/http"

	"github.com/panicmilos/druz.io/UserRelationsService/services"
	"github.com/sarulabs/di"
)

var DiMiddleware = func(next http.Handler) http.Handler {
	return di.HTTPMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		defer next.ServeHTTP(w, r)
	}), services.Provider, func(msg string) {
		log.Fatal(msg)
	})

}
