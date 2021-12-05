package utils

import (
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func ShouldAuth(ctx context.Context) error {
	token, _ := ctx.Value("auth_token").(*jwt.Token)
	if token == nil || !token.Valid {
		return errors.New("user should be authenticated")
	}
	return nil
}
