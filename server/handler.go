package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Broadcast-Server/common"
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

		hub.mutex.Lock()
		for client := range hub.clients {
			if username == strings.ToLower(client.username) {
				hub.mutex.Unlock()
				websocketError, _ := common.NewWebsocketError("Username already taken").ToJson()
				conn.WriteMessage(websocket.TextMessage, websocketError)
				conn.Close()
				return
			}
		}
		hub.mutex.Unlock()

		client := &Client{conn: conn, username: username}
		hub.register <- client

		defer func() {
			hub.unregister <- client
			client.conn.Close()
			log.Printf("Client %s disconnected\n", username)
		}()

		log.Printf("Client %s connected\n", username)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error from %s: %v\n", username, err)
				break
			}
			formatMessage := fmt.Sprintf("%s: %s", username, message)
			hub.broadcast <- []byte(formatMessage)
		}
	}
}
