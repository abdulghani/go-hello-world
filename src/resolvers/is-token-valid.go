package resolvers

import (
	"context"
	"hello_world/src/utils"
)

func IsTokenValid(ctx context.Context, token string) (bool, error) {
	res, err := utils.VerifyToken(token)

	return res.Valid, err
}
