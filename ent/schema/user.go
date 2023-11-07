package schema

import (
	"lkuoch/ent-todo/ent/schema/mixins"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Name
func (User) Name() string {
	return "users"
}

// Mixins of the User.
func (u User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NewIdMixin(u.Name()),
		mixins.TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("username").
			Unique().
			MaxLen(18).
			NotEmpty(),

		field.
			String("display_name").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// Normal indexes
		index.Fields("display_name"),

		// Unique indexes
		index.Fields("username").Unique(),
	}
}
