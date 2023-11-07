package main

import (
	"context"
	"log"
	"net/http"

	ent "lkuoch/ent-todo/ent/generated"
	"lkuoch/ent-todo/ent/generated/migrate"
	gql "lkuoch/ent-todo/ent/resolvers"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "test-db.db?_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Setup resolver name service
	gql.ResolverNameService{}.Init()

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(gql.NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
