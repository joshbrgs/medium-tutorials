# syntax=docker/dockerfile:1

# Create a stage for building the application.
ARG GO_VERSION=1.22
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Install protoc and the Go plugins for protobuf
RUN apt-get update && apt-get install -y \
  protobuf-compiler \
  && rm -rf /var/lib/apt/lists/*

# Install the Go protobuf and gRPC plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

# Download dependencies as a separate step to take advantage of Docker's caching.
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the entire application source code
COPY . .

# Generate the protobuf code
RUN protoc --proto_path=./cmd/nemesis/api/proto --go_out=./cmd/nemesis/api/gen --go-grpc_out=./cmd/nemesis/api/gen ./cmd/nemesis/api/proto/*.proto

# Build the application for the nemesis service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /nemesis_service ./cmd/nemesis

# Create a new stage for running the application
FROM ubuntu:latest AS final

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
  ca-certificates \
  tzdata \
  && rm -rf /var/lib/apt/lists/*

# Create a non-privileged user
# Copy the executable from the build stage
COPY --from=build /nemesis_service /

# Ensure the binary is executable
RUN chmod +x /nemesis_service

# Expose the port that the application listens on
EXPOSE 50051

# Set the entry point
ENTRYPOINT [ "/nemesis_service" ]

