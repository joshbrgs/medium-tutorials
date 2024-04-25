package nemesis

import "github.com/medium-tutorials/bad-inc/pkgs/server"

func main() {
	options := server.DefaultHTTPServerOptions()
	server := server.NewHTTPServer(options)

	if err := server.Start(":" + options.Port); err != nil {
		server.Logger.Panic(err)
	}

	server.Logger.Info("server started on http://localhost:" + options.Port)
}
