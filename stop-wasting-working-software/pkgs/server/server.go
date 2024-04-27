package server

import (
	"net"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New Echo Server with Options Pattern
type ServerOptions func(*echo.Echo)

func WithPort(port int) ServerOptions {
	return func(e *echo.Echo) {
		l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			e.Logger.Fatal(err)
		}

		e.Listener = l
	}
}

func WithTimeout(timeout time.Duration) ServerOptions {
	return func(e *echo.Echo) {
		e.Server.ReadTimeout = timeout
		e.Server.WriteTimeout = timeout
	}
}

func NewServer(opts ...ServerOptions) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Default server options

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Listener = l
	e.Server.ReadTimeout = 30 * time.Second
	e.Server.WriteTimeout = 30 * time.Second

	for _, opt := range opts {
		opt(e)
	}

	return e
}
