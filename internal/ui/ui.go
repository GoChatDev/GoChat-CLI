package ui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GoChatDev/GoChat-CLI/internal/client"
)

// Run starts the CLI user interface for messaging
func Run(wsClient *client.WebSocketClient) {
	fmt.Println("Welcome to the CLI Chat!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		if scanner.Scan() {
			message := scanner.Text()
			if err := wsClient.SendMessage(message); err != nil {
				fmt.Printf("Failed to send message: %v\n", err)
			}
		}

		// Receive messages from the WebSocket server
		go func() {
			for {
				msg, err := wsClient.ReceiveMessage()
				if err != nil {
					fmt.Printf("Error receiving message: %v\n", err)
					break
				}
				fmt.Printf("Received: %s\n", msg)
			}
		}()
	}
}
