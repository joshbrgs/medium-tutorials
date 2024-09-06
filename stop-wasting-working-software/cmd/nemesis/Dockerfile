# Use the official Golang image
FROM golang:1.19-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o nemesis_service .

# Expose the gRPC port
EXPOSE 50051

# Command to run the application
CMD ["./nemesis_service"]
