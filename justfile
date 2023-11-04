# This is a justfile. See https://github.com/casey/just

up:
  docker-compose build && docker-compose up

list:
  just --list

build:
  go build service/ent-todo.go

generate:
  go generate .

regenerate:
  rm -rf ./ent-schema.graphql ./ent/generated && just generate

migration-create:
  atlas migrate diff migration_name \
    --dir "file://ent/migrate/migrations" \
    --to "ent://ent/schema" \
    --dev-url "sqlite://file?mode=memory&_fk=1"

migration-apply:
  atlas migrate apply \
    --dir "file://ent/migrate/migrations" \
    --url "sqlite://test-db.db?_fk=1"

migration-status:
  atlas migrate status \
    --dir "file://ent/migrate/migrations" \
    --url "sqlite://test-db.db?_fk=1"

migration-diff:
  atlas migrate diff migration_name \
    --dir "file://ent/migrate/migrations" \
    --to "ent://ent/schema" \
    --dev-url "sqlite://test-db?mode=memory&_fk=1"

migration-run:
  just migration-create && just migration-apply && just migration-status

run:
  go run service/server.go
