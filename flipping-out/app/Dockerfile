FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 3001

ENTRYPOINT ["./app"]
