# Use a specific version of the golang image based on Debian
FROM golang:1.21-alpine

WORKDIR /app

# Install glibc on Alpine
RUN apk --no-cache add ca-certificates wget \
	&& wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
	&& wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-2.31-r0.apk \
	&& apk add --no-cache --allow-untrusted --force-overwrite glibc-2.31-r0.apk


# Install dependencies
RUN apk update \
	&& apk add --no-cache just protobuf build-base curl protobuf protobuf-dev

# Download and install protoc
ARG PROTOC_VERSION=3.19.1
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local \
	&& rm protoc-${PROTOC_VERSION}-linux-x86_64.zip

# Install Go tools and plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
	&& go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest \
	&& go install github.com/cosmtrek/air@latest

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

EXPOSE 8081
CMD ["air"]
