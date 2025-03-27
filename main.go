package main

import (
	"net/http"

	"github.com/Broadcast-Server/server"
)

func main() {
	hub := server.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", server.HandleConnections(hub))

	http.ListenAndServe(":3000", nil)
}
