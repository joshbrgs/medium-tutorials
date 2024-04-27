package server

import (
	"time"

	"google.golang.org/grpc"
)

// GRPCServerOption represents a function that modifies gRPC server options
type GRPCServerOption func(*grpc.Server)

// WithGRPCAddress sets the address option for the gRPC server
func WithGRPCAddress(address string) GRPCServerOption {
	return func(s *grpc.Server) {
		s.Addr = address
	}
}

// WithGRPCMaxConcurrentStreams sets the MaxConcurrentStreams option for the gRPC server
func WithGRPCMaxConcurrentStreams(maxConcurrentStreams uint32) GRPCServerOption {
	return func(s *grpc.Server) {
		s.MaxConcurrentStreams = maxConcurrentStreams
	}
}

// WithGRPCIdleTimeout sets the IdleTimeout option for the gRPC server
func WithGRPCIdleTimeout(idleTimeout time.Duration) GRPCServerOption {
	return func(s *grpc.Server) {
		s.IdleTimeout = idleTimeout
	}
}

// NewGRPCServer initializes a new instance of gRPC server with the provided options
func NewGRPCServer(opts ...GRPCServerOption) *grpc.Server {
	// Create a new gRPC server instance
	s := grpc.NewServer()

	// Apply provided options
	for _, opt := range opts {
		opt(s)
	}

	return s
}
