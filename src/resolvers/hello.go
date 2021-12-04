package resolvers

import "context"

func Hello(ctx context.Context) (string, error) {
	return "hello world", nil
}
