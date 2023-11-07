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
	tableHash string
	tableName string
}

func NewIdMixin(tableName string) *IdMixin {
	return &IdMixin{
		tableName: tableName,
		tableHash: pulid.ExtractPrefixHash(pulid.New(tableName)),
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
	return a.Prefix
}

func (i IdMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: i.tableHash},
	}
}
