FROM golang:1.20.3-bullseye

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum .
RUN go mod download

# Pull in rest
COPY . .

# Codegen
RUN go generate .

# Run app
CMD ["go", "run", "service/ent-todo.go"]