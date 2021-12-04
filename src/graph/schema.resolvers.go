package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"hello_world/src/graph/generated"
	"hello_world/src/graph/model"
	"hello_world/src/resolvers"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	return resolvers.SingleUpload(ctx, file)
}

func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	return resolvers.SingleUploadWithPayload(ctx, req)
}

func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	return resolvers.MultipleUpload(ctx, files)
}

func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	return resolvers.MultipleUploadWithPayload(ctx, req)
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return resolvers.Hello(ctx)
}

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
