package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	ent "lkuoch/ent-todo/ent/generated"
	gql "lkuoch/ent-todo/ent/resolvers"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	return connectionString
}

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, getConnectionString())
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()
	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(gql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
