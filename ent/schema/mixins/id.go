package mixins

import (
	pulid "lkuoch/ent-todo/ent/schema/types/pulid"

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
	TableName() string
}

func NewIdMixin(meta IdMixinConfig) *IdMixin {
	return &IdMixin{
		tableName: meta.TableName(),
	}
}

func (i IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			Immutable().
			GoType(pulid.ID(i.tableName)).
			DefaultFunc(func() pulid.ID {
				return pulid.New(i.tableName)
			}),
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
		Annotation{Prefix: pulid.NewPrefix(i.tableName)},
	}
}
