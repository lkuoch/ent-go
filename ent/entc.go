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
		entgql.WithSchemaPath("ent-schema.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
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
	}, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
