package todo

//go:generate go run -mod=mod ./ent/entc.go
//go:generate go run entgo.io/contrib/entproto/cmd/entproto -path ./ent/schema
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate go generate ./ent/proto/...
