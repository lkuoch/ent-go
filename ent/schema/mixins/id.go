package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	ulid "github.com/oklog/ulid/v2"
)

type IdMixin struct {
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			Immutable().
			Default(ulid.Make().String()),
	}
}
