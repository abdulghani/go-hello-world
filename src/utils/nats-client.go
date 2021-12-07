package utils

import (
	"os"

	"github.com/nats-io/nats.go"
)

var client *nats.Conn

func GetNatsClient() *nats.Conn {
	if client != nil {
		return client
	}
	client, err := nats.Connect(os.Getenv("NATS_SERVER"))

	if err != nil {
		panic(err)
	}

	return client
}
