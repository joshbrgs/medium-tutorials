# syntax=docker/dockerfile:1

# Create a stage for building the application.
FROM golang:1.22 AS builder

WORKDIR /src
# Install required tools: buf for managing proto files and dependencies
RUN apt-get update && apt-get install -y curl unzip git

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install buf CLI
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v1.40.1/buf-Linux-x86_64 -o /usr/local/bin/buf && chmod +x /usr/local/bin/buf

# Download dependencies as a separate step to take advantage of Docker's caching.
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the entire application source code
COPY . .

WORKDIR /src/cmd/nemesis/api
# Fetch proto dependencies using buf
RUN buf dep update

# Generate Go code from proto files using buf
RUN buf generate

WORKDIR /src 

# Build the application for the nemesis service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /nemesis_service ./cmd/nemesis

# Create a new stage for running the application
FROM ubuntu:latest AS final

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
  ca-certificates \
  tzdata \
  && rm -rf /var/lib/apt/lists/*

# Create a non-privileged user
# Copy the executable from the build stage
COPY --from=builder /nemesis_service /

# Ensure the binary is executable
RUN chmod +x /nemesis_service

# Expose the port that the application listens on
EXPOSE 50051

# Set the entry point
ENTRYPOINT [ "/nemesis_service" ]

