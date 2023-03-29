# This is a justfile. See https://github.com/casey/just

list:
  just --list

build:
  go build service/ent-todo.go

generate:
  go generate .

run:
  go run service/ent-todo.go
