package main

import (
	"log"

	"github.com/GoChatDev/GoChat-CLI/internal/client"
	"github.com/GoChatDev/GoChat-CLI/internal/ui"
)

func main() {
	// Authenticate the user via SSH and get user ID
	userID, err := client.AuthenticateUser()
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// Start the WebSocket client with the authenticated user ID
	wsClient := client.NewWebSocketClient(userID)
	if err := wsClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}

	// Launch the UI for user interaction
	ui.Run(wsClient)
}
