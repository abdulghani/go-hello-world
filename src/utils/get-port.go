package utils

import (
	"os"
)

const defaultPort = "3000"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return port
}
