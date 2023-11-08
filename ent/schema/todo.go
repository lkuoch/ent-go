package schema

import (
	"lkuoch/ent-todo/ent/schema/mixins"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

func (Todo) TableName() string {
	return "todo"
}

// Mixins of the Todo.
func (t Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NewIdMixin(t),
		mixins.TimeMixin{},
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			MaxLen(26).
			NotEmpty().Annotations(
			entgql.OrderField("TITLE"),
		),

		field.
			Enum("priority").
			NamedValues(
				"High", "HIGH",
				"Medium", "MEDIUM",
				"Low", "LOW",
				"None", "NONE",
			).
			Default("NONE"),

		field.
			Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),

		field.
			Time("time_completed").
			Nillable().
			Optional(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		// A todo can belong to only one user
		edge.From("user", User.Type).
			Ref("todos").
			Unique(),
	}
}

// Annotations of the Todo.
func (t Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: t.TableName()},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Indexes of the Todo
func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		// Normal indexes
		index.Fields("title", "priority", "status"),
	}
}
