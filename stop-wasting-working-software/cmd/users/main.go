package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joshbrgs/mongorm/cmd/mongorm"
	"github.com/medium-tutorials/bad-inc/cmd/users/users/users"
	"github.com/medium-tutorials/bad-inc/pkgs/server"
)

func main() {
	e := server.NewServer(
		server.WithPort(8081),
	)
	//
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

	userService := users.NewUserService(db, rmq)

	go func() {
		e.Logger.Fatal(e.Start(""))
	}()

	e.Logger.Info("server started on http://" + e.Server.Addr)

	e.POST("/", userService.CreateUserHandler)
	e.GET("/:id", userService.GetUserByIdHandler)
	e.DELETE("/:id", userService.DeleteUserHandler)
	e.PUT("/:id", userService.UpdateUserHandler)
	e.POST("/login", userService.LoginHandler)

	// Block main goroutine until a signal is received
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Gracefully shut down the server when a signal is received
	log.Println("Shutting down the server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
