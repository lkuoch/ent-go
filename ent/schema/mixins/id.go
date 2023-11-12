package mixins

import (
	"lkuoch/ent-todo/ent/schema/types"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type IdMixin struct {
	mixin.Schema
	tableName string
}

type IdMixinConfig interface {
	EntityName() string
}

func NewIdMixin(meta IdMixinConfig) *IdMixin {
	return &IdMixin{
		tableName: meta.EntityName(),
	}
}

func (i IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			Immutable().
			GoType(types.ID(i.tableName)).
			DefaultFunc(func() types.ID {
				return types.New(i.tableName)
			}).
			Annotations(
				entgql.OrderField("ID"),
			),
	}
}

type Annotation struct {
	Prefix string
}

func (a Annotation) Name() string {
	return "__PULID__"
}

func (i IdMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: types.NewPrefix(i.tableName)},
	}
}
