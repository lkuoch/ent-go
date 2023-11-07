package resolvers

import (
	ent "lkuoch/ent-todo/ent/generated"
	gql "lkuoch/ent-todo/ent/generated/gql"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}
