# This is a justfile. See https://github.com/casey/just

list:
  just --list

build:
  go build service/ent-todo.go

generate:
  go generate .

regenerate:
  rm -rf ./schema.graphql ./ent/generated ./ent/resolvers/*.resolvers.go && just generate

run:
  go run service/ent-todo.go
