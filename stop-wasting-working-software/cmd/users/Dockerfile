# syntax=docker/dockerfile:1

################################################################################
# Create a stage for building the application.
ARG GO_VERSION=1.22
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage bind mounts to go.sum and go.mod to avoid having to copy them into
# the container.
COPY go.mod go.sum ./
RUN go mod download -x

# Copy the entire application source code
COPY . .

# Build the application for the user service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /users ./cmd/users

################################################################################
# Create a new stage for running the application that contains the minimal
# runtime dependencies for the application.
FROM alpine:latest AS final

# Install runtime dependencies
RUN apk --update add \
  ca-certificates \
  tzdata \
  && \
  update-ca-certificates

# Create a non-privileged user that the app will run under.


# Copy the executable from the "build" stage.
COPY --from=build /users /

# Ensure the binary is executable
RUN chmod +x /users
# Expose the port that the application listens on.
EXPOSE 8081

# What the container should run when it is started.
ENTRYPOINT [ "/users" ]

