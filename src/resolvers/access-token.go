package resolvers

import (
	"context"
	"hello_world/src/utils"
)

func AccessToken(ctx context.Context) (string, error) {
	token := utils.SignToken("hello world")

	return token, nil
}
