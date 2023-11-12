# Simple TODO server written with ent-go

## Context:

- `ent/schema` go models are transformed into graphql types + go resolvers: `ent-schema.graphql`, `ent/resolvers/ent-schema.go`
- `typedef.graphql` queries/mutations are transformed into go resolver: `ent/resolvers/typedef.go`

## Development flow:

0. `just migrate-run`
1. `just regenerate`
   - Needed when changing schema files
2. `just run`
   - Spins up graphql server

## Arch
```mermaid
erDiagram
    USER ||--o{ TODO : has
    USER {
        ID id
        datetime created_at
        datetime updated_at

        string username
        string display_name

        string password_hash
        string password_salt
    }

    TODO {
        ID id
        datetime created_at
        datetime updated_at

        string title
        string body
        datetime time_completed
        string priority FK
        string status FK
    }
    TODO ||--o{ TASK : has
    TODO ||--|| PRIORITY : has
    TODO ||--|| STATUS : has

    TASK {
        ID id
        string title
        string status FK
    }
    TASK ||--|| STATUS : has

    PRIORITY {
        string HIGH
        string MEDIUM
        string LOW
        string NONE
    }

    STATUS {
        string IN_PROGRESS
        string COMPLETED
    }
```

## Tasks
- [x] Setup migrations
- [x] Setup Relay Node Interface
- [x] Setup relations
- [x] Setup more tables + operations
- [ ] Setup permissions
- [ ] Setup multi-tenancy