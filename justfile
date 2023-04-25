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

run:
  go run service/ent-todo.go
