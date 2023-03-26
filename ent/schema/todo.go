package schema

import (
	"lkuoch/ent-todo/ent/schema/enums"
	"lkuoch/ent-todo/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Mixins of the Todo.
func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IdMixin{},
		mixins.TimeMixin{},
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			MaxLen(26).
			NotEmpty(),

		field.
			Enum("priority").
			GoType(enums.TodoPriority("")).
			Default(string(enums.None)),

		field.
			Time("time_completed").
			Nillable().
			Optional(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todos"),
	}
}
