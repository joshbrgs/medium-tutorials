package server

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// GRPCServerOption represents a function that modifies gRPC server options
type GRPCServerOption func(*grpc.ServerOption)

// WithGRPCMaxConcurrentStreams sets the MaxConcurrentStreams option for the gRPC server
func WithGRPCMaxConcurrentStreams(maxConcurrentStreams uint32) grpc.ServerOption {
	return grpc.MaxConcurrentStreams(maxConcurrentStreams)
}

// WithGRPCIdleTimeout sets the IdleTimeout option for the gRPC server
func WithGRPCIdleTimeout(idleTimeout time.Duration) grpc.ServerOption {
	return grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: idleTimeout,
	})
}

// NewGRPCServer initializes a new instance of gRPC server with the provided options
func NewGRPCServer(opts ...grpc.ServerOption) *grpc.Server {
	// Create a new gRPC server instance with the provided options
	return grpc.NewServer(opts...)
}

