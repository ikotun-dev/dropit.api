package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("my-extremely-secret-auth-checker")

/*
func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
*/

func CreateJWT(sessionKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["session_key"] = sessionKey // Include session key as a claim

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
