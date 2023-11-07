# Simple TODO server written with ent-go

## Context:

- `ent/schema` go models are transformed into graphql types + go resolvers: `ent-schema.graphql`, `ent/resolvers/ent-schema.go`
- `typedef.graphql` queries/mutations are transformed into go resolver: `ent/resolvers/typedef.go`

## Development flow:

1. `just generate`
   - Needed when changing schema files
2. `just run`
   - Spins up graphql server


## Tasks
- [x] Setup migrations
- [x] Setup Relay Node Interface
- [ ] Setup relations
- [ ] Setup more tables + operations
- [ ] Setup permissions
- [ ] Setup multi-tenancy