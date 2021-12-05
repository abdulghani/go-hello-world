package utils

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtKey struct {
	PublicKey  []byte
	PrivateKey []byte
}

var keyFile jwtKey

func getKeyFile() jwtKey {
	if keyFile.PrivateKey != nil && keyFile.PublicKey != nil {
		return keyFile
	}
	publicKeyPath := os.Getenv("JWT_PUBLIC_KEY_FILE")
	privateKeyPath := os.Getenv("JWT_PRIVATE_KEY_FILE")

	publicKeyFile, _ := ioutil.ReadFile(publicKeyPath)
	privateKeyFile, _ := ioutil.ReadFile(privateKeyPath)

	keyFile.PublicKey = publicKeyFile
	keyFile.PrivateKey = privateKeyFile

	return keyFile
}

func SignToken(payload interface{}) (string, error) {
	key := getKeyFile()
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Id:        "abdul.ghani@kaisa.id",
	})

	return token.SignedString(key.PrivateKey)
}

func VerifyToken(tokenstr string) (*jwt.Token, error) {
	key := getKeyFile()

	return jwt.ParseWithClaims(tokenstr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key.PublicKey, nil
	})
}
