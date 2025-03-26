package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*Client]struct{}
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]struct{}),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = struct{}{}
			h.mutex.Unlock()

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.conn.Close()
			}
			h.mutex.Unlock()
		case message := <-h.broadcast:
			h.mutex.Lock()
			for key, _ := range h.clients {
				key.conn.WriteMessage(websocket.TextMessage, message)
			}
			h.mutex.Unlock()
		}
	}
}
