# This is a justfile. See https://github.com/casey/just

up:
  docker-compose build && docker-compose up

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
    --dev-url "sqlite://file?mode=memory&_fk=1"

migrate-apply:
  atlas migrate apply \
    --dir "file://ent/migrate/migrations" \
    --url "sqlite://test-db.db?_fk=1"

migrate-status:
  atlas migrate status \
    --dir "file://ent/migrate/migrations" \
    --url "sqlite://test-db.db?_fk=1"

migrate-diff:
  atlas migrate diff migration_name \
    --dir "file://ent/migrate/migrations" \
    --to "ent://ent/schema" \
    --dev-url "sqlite://test-db?mode=memory&_fk=1"

migrate-run:
  just migrate-create && just migrate-apply && just migrate-status

reset-all:
  rm -rf ./test-db.db && just migrate-create && just migrate-apply && just migrate-status && just regenerate

run:
  go run service/server.go
