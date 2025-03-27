package cmd

import (
	"log"
	"net/http"

	"github.com/Broadcast-Server/server"
	"github.com/spf13/cobra"
)

func StartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "This command will start the server",
		Run: func(cmd *cobra.Command, args []string) {
			hub := server.NewHub()
			go hub.Run()
			http.HandleFunc("/ws", server.HandleConnections(hub))
			log.Println("Server started at http://localhost:3000/")
			http.ListenAndServe(":3000", nil)
		},
	}
	return cmd
}
