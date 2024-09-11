package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joshbrgs/mongorm/cmd/mongorm"
	pb "github.com/medium-tutorials/bad-inc/cmd/nemesis/api/gen"
	"github.com/medium-tutorials/bad-inc/cmd/nemesis/service"
	"github.com/medium-tutorials/bad-inc/pkgs/server"

	"google.golang.org/grpc"
)

func main() {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	connectionString := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", user, pass)

	client, ctx, err := mongorm.Connect(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to mongodb with uri: %s", connectionString)
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB")
	}()

	ruser := os.Getenv("RABBITMQ_DEFAULT_USER")
	rpass := os.Getenv("RABBITMQ_DEFAULT_PASS")
	url := fmt.Sprintf("amqp://%s:%s@rabbitmq:5672", ruser, rpass)

	// Create a connection with options
	rmq, err := server.NewRabbitMQ(server.WithRabbitMQURL(url))
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ connection: %v", err)
	}
	defer rmq.Close()

	db := client.Database("users")

	grpcServer := grpc.NewServer()
	pb.RegisterNemesisServiceServer(grpcServer, service.NewNemesisServiceServer(db, rmq))

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Nemesis Service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
