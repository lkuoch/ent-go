package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"lkuoch/ent-todo/ent/generated"
	"lkuoch/ent-todo/ent/generated/gql"
	"lkuoch/ent-todo/ent/schema/types"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id types.ID) (generated.Noder, error) {
	return r.client.Noder(ctx, id, generated.WithNodeType(generated.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []types.ID) ([]generated.Noder, error) {
	return r.client.Noders(ctx, ids, generated.WithNodeType(generated.IDToType))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, after *entgql.Cursor[types.ID], first *int, before *entgql.Cursor[types.ID], last *int, orderBy []*generated.TaskOrder, where *generated.TaskWhereInput) (*generated.TaskConnection, error) {
	return r.client.Task.Query().
		Paginate(ctx, after, first, before, last,
			generated.WithTaskOrder(orderBy),
		)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[types.ID], first *int, before *entgql.Cursor[types.ID], last *int, orderBy []*generated.TodoOrder, where *generated.TodoWhereInput) (*generated.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			generated.WithTodoOrder(orderBy),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[types.ID], first *int, before *entgql.Cursor[types.ID], last *int, orderBy []*generated.UserOrder, where *generated.UserWhereInput) (*generated.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			generated.WithUserOrder(orderBy),
		)
}

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
