package resolvers

import (
	"context"
	"hello_world/src/utils"

	"github.com/dgrijalva/jwt-go"
)

func Hello(ctx context.Context) (string, error) {
	token, _ := utils.SignToken("hello world")

	raw, _ := ctx.Value("auth_token").(*jwt.Token)
	utils.Inspect("CONTEXT[\"auth_token\"]", raw.Header, raw.Claims, raw.Valid)

	return token, nil
}
