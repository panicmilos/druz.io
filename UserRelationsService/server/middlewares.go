package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/panicmilos/druz.io/UserRelationsService/errors"
	"github.com/panicmilos/druz.io/UserRelationsService/helpers"
	"github.com/panicmilos/druz.io/UserRelationsService/models"
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

var AuthenticateMiddlewate = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			errors.Handle(errors.NewErrUnauthorized("Missing Authorization Header"), w)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := helpers.VerifyJwtToken(tokenString)
		if err != nil {
			errors.Handle(errors.NewErrUnauthorized("Error verifying JWT token: "+err.Error()), w)
			return
		}

		r.Header.Set("name", claims.(jwt.MapClaims)["name"].(string))
		r.Header.Set("role", claims.(jwt.MapClaims)["role"].(string))

		next.ServeHTTP(w, r)
	})
}

var AuthorizeMiddlewate = func(next http.Handler, roles ...models.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role := r.Header.Get("role")
		contains := false
		for _, ar := range roles {
			if role == strconv.Itoa((int(ar))) {
				contains = true
				break
			}
		}

		if contains {
			next.ServeHTTP(w, r)
		} else {
			errors.Handle(errors.NewErrForbidden("You don't have permission for this action"), w)
		}
	})
}
