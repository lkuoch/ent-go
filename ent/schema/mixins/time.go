package mixins

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}
