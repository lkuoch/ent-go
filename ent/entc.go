//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	extension, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithSchemaPath("ent.graphql"),
		entgql.WithRelaySpec(true),
		entgql.WithWhereInputs(true),
	)

	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	options := []entc.Option{
		entc.Extensions(extension),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Target:  "./ent/generated",
		Package: "lkuoch/ent-todo/ent/generated",
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeatureUpsert,
		},
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("pulid").ParseFiles("./ent/schema/templates/pulid.tmpl")),
		},
	}, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
