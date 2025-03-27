package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := strings.ToLower(r.URL.Query().Get("username"))
		if username == "" {
			http.Error(w, "Missing 'username' query parameter", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		for client := range hub.clients {
			if username == strings.ToLower(client.username) {
				conn.WriteMessage(websocket.TextMessage, []byte("Username already exists"))
				conn.Close()
				return
			}
		}

		client := &Client{conn: conn, username: username}
		hub.register <- client

		defer func() {
			hub.unregister <- client
			client.conn.Close()
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			hub.broadcast <- message
		}
	}
}
