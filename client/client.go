package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

const serverAddress = "ws://localhost:3000/ws"

func connectWebSocket(username string) (*websocket.Conn, error) {
	serverURL := fmt.Sprintf("%s?username=%s", serverAddress, username)
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to server as:", username)
	return conn, nil
}

func readMessages(conn *websocket.Conn, username string) {
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

func writeMessages(conn *websocket.Conn, username string) {
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

// func main() {
// 	username := "testuser"
// 	conn, err := connectWebSocket(username)
// 	if err != nil {
// 		log.Fatal("Connection error:", err)
// 	}
// 	defer conn.Close()
//
// 	go readMessages(conn, username)
// 	writeMessages(conn, username)
//
// 	log.Println("Disconnected from server")
// }
