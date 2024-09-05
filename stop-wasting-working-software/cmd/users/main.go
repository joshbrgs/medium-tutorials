package users

import (
	"fmt"
	"os"

	"github.com/joshbrgs/mongorm/cmd/mongorm"

	"github.com/medium-tutorials/bad-inc/pkgs/server"
)

func main() {
	server := server.NewServer(
		server.WithPort(8000),
	)

	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", user, pass)

	client, err := mongorm.Connect(connectionString)
	if err != nil {
		panic(err)
	}

	db := client.Database("users")

	userService := NewUserService(db)

	go func() {
		server.Logger.Fatal(server.Start(""))
	}()

	server.Logger.Info("server started on http://" + server.Server.Addr)

	server.POST("/users", userService.createUserHandler)
	server.GET("/users/:id", userService.getUserByIdHandler)
	server.DELETE("/users/:id", userService.deleteUserHandler)
	server.PUT("/users/:id", userService.updateUserHandler)
	server.POST("/login", userService.loginHandler)
}
