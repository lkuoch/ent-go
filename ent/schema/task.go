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

// Task entity
type Task struct {
	ent.Schema
}

func (Task) EntityName() string {
	return "task"
}

// Mixins of the Task
func (t Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.NewIdMixin(t),
	}
}

// Fields of the Task
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			NotEmpty().
			Annotations(
				entgql.OrderField("TITLE"),
			),

		field.
			Enum("item_status").
			GoType(types.ItemStatus("")).
			Default(string(types.ItemStatus_InProgress)).
			Annotations(
				entgql.Type("ItemStatus"),
			),
	}
}

// Edges of the Task
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// A Task belongs to a Todo
		edge.From("todo", Todo.Type).
			Ref("tasks").
			Unique().
			Required().
			Comment("Task belongs to single Todo"),
	}
}

// Annotations of the Task
func (t Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: t.EntityName()},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

// Indexes of the Task
func (Task) Indexes() []ent.Index {
	return []ent.Index{
		// Normal indexes
		index.Fields("title", "item_status"),
	}
}
