# Start from the official Go image for building the binary
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary with the correct architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o users ./cmd/users

# Start a new stage from a minimal base image
FROM alpine:3.18

# Install any necessary dependencies (optional)
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/users /users

RUN chmod +x /users

EXPOSE 8081

# Set the binary as the entrypoint
ENTRYPOINT ["/users"]
