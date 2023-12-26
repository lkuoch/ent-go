FROM golang:1.21-alpine

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Pull in rest
COPY . .

# Get filewatching going
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8081
CMD ["air"]