package helpers

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Name string
	Role string
}

func GetJwtToken(claims Claims) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": claims.Name,
		"role": claims.Role,
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	return token.SignedString(signingKey)
}

func VerifyJwtToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
