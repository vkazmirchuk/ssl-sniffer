package main

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// Run web server
	go InitWebServer()

	// Create sniffer
	sniffer, err := NewSniffer(os.Args[1])
	if err != nil {
		fmt.Println("[ERROR create sniffer]", err)
		return
	}

	// Waiting for a message from sniffer
	for message := range sniffer.Run() {

		// Get all ws connections
		for _, conn := range WSConnections.GetAll() {

			// Send message to ws connection
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Println(err)
				continue
			}
		}

		// Print message to stdout
		fmt.Println(message)
	}
}
