package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

const serverAddress = "ws://localhost:3000/ws"

func ConnectWebSocket(username string) (*websocket.Conn, error) {
	serverURL := fmt.Sprintf("%s?username=%s", serverAddress, username)
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to server as:", username)
	return conn, nil
}

func ReadMessages(conn *websocket.Conn, username string) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		msg := string(message)

		if len(msg) > len(username)+2 && msg[:len(username)+2] == username+": " {
			continue
		}

		fmt.Println("\n[Server]:", msg)
	}
}

func WriteMessages(conn *websocket.Conn, username string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n[You] (Type message): ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if text == "/exit" {
			log.Println("Closing connection...")
			break
		}
		formattedMsg := fmt.Sprintf("%s: %s", username, text)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(formattedMsg)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
