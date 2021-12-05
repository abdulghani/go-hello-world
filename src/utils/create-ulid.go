package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func CreateUlid() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	value := ulid.MustNew(ulid.Timestamp(t), entropy)

	return value.String()
}
