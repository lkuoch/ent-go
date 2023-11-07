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

func NewIdMixin(tableName string) *IdMixin {
	return &IdMixin{tableName: tableName}
}

func (i IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			Immutable().
			GoType(pulid.ID("")).
			DefaultFunc(func() pulid.ID {
				return pulid.New(i.tableName)
			}),
	}
}

type Annotation struct {
	Prefix string
}

func (a Annotation) Name() string {
	return "PULID"
}

func (i IdMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: i.tableName},
	}
}
