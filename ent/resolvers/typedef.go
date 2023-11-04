package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"lkuoch/ent-todo/ent/generated"
	"lkuoch/ent-todo/ent/generated/gql"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input generated.CreateTodoInput) (*generated.Todo, error) {
	return r.client.Todo.Create().SetInput(input).Save(ctx)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input generated.CreateUserInput) (*generated.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
