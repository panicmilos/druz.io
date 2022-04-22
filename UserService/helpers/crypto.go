package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GetRandomToken(length int32) string {
	rb := make([]byte, length)
	rand.Read(rb)

	return base64.URLEncoding.EncodeToString(rb)
}

func GetSaltedAndHashedPassword(password string, salt string) string {
	saltedPassword := password + salt
	saltedPasswordBytes := []byte(saltedPassword)
	hashedPasswordBytes := sha256.Sum256(saltedPasswordBytes)

	return base64.URLEncoding.EncodeToString(hashedPasswordBytes[:])
}
