package utils

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// STUPID KEY TURNS OUT TO BE THE SAME
var auth_key []byte

func getAuthKey() []byte {
	if auth_key != nil {
		return auth_key
	}
	auth_key, err := ioutil.ReadFile(os.Getenv("JWT_KEY_FILE"))
	ShouldPanic(err)

	return auth_key
}

func SignToken(payload interface{}) string {
	key := getAuthKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, &jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(time.Hour * 72).Unix(),
		Id:        "user@email.com",
	})

	tokenstr, err := token.SignedString(key)
	ShouldPanic(err)

	return tokenstr
}

func VerifyToken(tokenstr string) (*jwt.Token, error) {
	key := getAuthKey()
	return jwt.ParseWithClaims(tokenstr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
}
