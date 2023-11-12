package schema

import (
	"lkuoch/ent-todo/ent/schema/mixins"
	"lkuoch/ent-todo/ent/schema/types"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo entity
type Todo struct {
	ent.Schema
}

func (Todo) EntityName() string {
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
			String("body").
			Nillable(),

		field.
			Enum("item_priority").
			GoType(types.ItemPriority("")).
			Default(string(types.ItemPriority_None)),

		field.
			Enum("item_status").
			GoType(types.ItemStatus("")).
			Default(string(types.ItemStatus_InProgress)).
			Annotations(
				entgql.Type("ItemStatus"),
			),

		field.
			Time("time_completed").
			Nillable().
			Optional(),
	}
}

// Edges of the Todo
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		// A todo can belong to only one user
		edge.From("user", User.Type).
			Ref("todos").
			Unique().
			Required().
			Comment("Todo belongs to single User"),

		// A todo can have many tasks
		edge.To("tasks", Task.Type).
			Annotations(entgql.RelayConnection()).
			Comment("Todo has multiple Tasks"),
	}
}

// Annotations of the Todo
func (t Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: t.EntityName()},
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
		index.Fields("title", "item_priority", "item_status"),
	}
}
