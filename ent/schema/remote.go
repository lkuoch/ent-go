package schema

import (
	"lkuoch/ent-todo/ent/schema/types"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Remote holds the schema definition for the Remote entity.
type Remote struct {
	ent.Schema
}

func (Remote) EntityName() string {
	return "remote"
}

// Fields of the Remote
func (r Remote) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			Immutable().
			GoType(types.ID(r.EntityName())).
			Annotations(
				entproto.Field(1),
			),

		field.
			String("data").
			Annotations(
				entproto.Field(2),
			),
	}
}

// Annotations of the Remote
func (r Remote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
