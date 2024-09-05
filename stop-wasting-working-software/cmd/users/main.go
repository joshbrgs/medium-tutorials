package users

import (
	"fmt"
	"log"
	"os"

	"github.com/joshbrgs/mongorm/cmd/mongorm"

	"github.com/medium-tutorials/bad-inc/pkgs/server"
)

func main() {
	e := server.NewServer(
		server.WithPort(8000),
	)

	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", user, pass)

	client, err := mongorm.Connect(connectionString)
	if err != nil {
		panic(err)
	}

	// defer client.Disconnect()

	url := "amqp://localhost:5672"

	// Create a connection with options
	conn, err := server.NewRabbitMQConnection(url,
		server.WithExchange("user_exchange", "fanout", nil),
		server.WithQueue("user_queue", nil),
		// WithCredentials(...), // Set username and password if needed
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Close connection on exit

	db := client.Database("users")

	userService := NewUserService(db)

	go func() {
		e.Logger.Fatal(e.Start(""))
	}()

	e.Logger.Info("server started on http://" + e.Server.Addr)

	e.POST("/users", userService.createUserHandler)
	e.GET("/users/:id", userService.getUserByIdHandler)
	e.DELETE("/users/:id", userService.deleteUserHandler)
	e.PUT("/users/:id", userService.updateUserHandler)
	e.POST("/login", userService.loginHandler)
}
