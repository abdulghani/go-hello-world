package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"hello_world/src/graph/generated"
	"hello_world/src/graph/model"
	"hello_world/src/resolvers"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	return resolvers.SingleUpload(ctx, file)
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return resolvers.Hello(ctx)
}

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	return "", nil
}

func (r *queryResolver) AccessToken(ctx context.Context) (string, error) {
	return resolvers.AccessToken(ctx)
}

func (r *queryResolver) IsTokenValid(ctx context.Context, token string) (bool, error) {
	return resolvers.IsTokenValid(ctx, token)
}

func (r *queryResolver) Countries(ctx context.Context) (string, error) {
	return resolvers.Countries(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
