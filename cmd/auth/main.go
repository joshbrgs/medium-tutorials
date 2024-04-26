package auth

import "github.com/medium-tutorials/bad-inc/pkgs/server"

func main() {
	server := server.NewServer(
		server.WithPort(8090),
	)

	go func() {
		server.Logger.Fatal(server.Start(""))
	}()

	server.Logger.Info("server started on http://" + server.Server.Addr)
}
