package readers

import (
	"log"
	"net"
	"time"

	"github.com/medium-tutorials/bad-inc/pkgs/server"
	"google.golang.org/grpc"
)

func main() {
	// Define gRPC server options
	opts := []grpc.ServerOption{
		server.WithGRPCMaxConcurrentStreams(100),    // Limit to 100 concurrent streams
		server.WithGRPCIdleTimeout(5 * time.Minute), // Set idle timeout to 5 minutes
	}

	// Create the gRPC server with the defined options
	grpcServer := server.NewGRPCServer(opts...)

	// Listen on a TCP address
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Register your gRPC services here
	// Example: pb.RegisterYourServiceServer(grpcServer, &yourServiceImplementation{})

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
