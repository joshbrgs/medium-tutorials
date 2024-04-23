package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServerOptions struct {
	Port string
}

type GrpcServerOptions struct{}

type MQOptions struct{}

func DefaultHTTPServerOptions() *HTTPServerOptions {
	return &HTTPServerOptions{
		Port: "8080",
	}
}

func DefaultGrpcServerOptions() *GrpcServerOptions {
	return &GrpcServerOptions{}
}

func DefaultMQOptions() *MQOptions {
	return &MQOptions{}
}

func WithHTTPServerPort(port string) *HTTPServerOptions {
	return &HTTPServerOptions{
		Port: port,
	}
}

func NewHTTPServer() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func NewGrpcServer() {
}

func NewMQConnection() {
}
