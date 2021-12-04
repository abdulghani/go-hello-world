package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		print("Not loading .env file. file doesn't exists.")
	}
}
