# This is a justfile. See https://github.com/casey/just

up:
  podman compose up --build -d

down:
  podman compose down

list:
  just --list

build:
  go build service/server.go

generate:
  go generate .

regenerate:
  rm -rf ./ent.graphql ./ent/generated ./ent/proto && just generate

migrate-create:
  atlas migrate diff migration_name \
    --dir "file://ent/migrate/migrations" \
    --to "ent://ent/schema" \
    --dev-url "postgres://user:password@localhost:5432/ent?search_path=public&sslmode=disable"

migrate-apply:
  atlas migrate apply \
    --dir "file://ent/migrate/migrations" \
    --url "postgres://user:password@localhost:5432/ent?search_path=public&sslmode=disable"

migrate-status:
  atlas migrate status \
    --dir "file://ent/migrate/migrations" \
    --url "postgres://user:password@localhost:5432/ent?search_path=public&sslmode=disable"

migrate-diff:
  atlas migrate diff migration_name \
    --dir "file://ent/migrate/migrations" \
    --to "ent://ent/schema" \
    --dev-url "postgres://user:password@localhost:5432/ent?search_path=public&sslmode=disable"

migrate-run:
  just migrate-create && just migrate-apply && just migrate-status

reset-all:
  rm -rf ./test-db.db && just migrate-create && just migrate-apply && just migrate-status && just regenerate

run:
  go run service/server.go
