package cmd

import (
	"log"

	"github.com/Broadcast-Server/client"
	"github.com/spf13/cobra"
)

var username string

func ConnectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Establish a connection to the server",
		Long:  `The connect command initiates a connection to the specified server using the provided parameters.`,
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" {
				log.Fatal("Username is required. Use --username or -u flag.")
			}

			conn, err := client.ConnectWebSocket(username)
			if err != nil {
				log.Fatal("Connection error:", err)
			}
			defer conn.Close()

			go client.ReadMessages(conn, username)
			client.WriteMessages(conn, username)

			log.Println("Disconnected from server")
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "Specify the username")

	return cmd
}
