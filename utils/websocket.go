// utils/websocket.go
package utils

import (
	"hearthstone-clone-backend/models"
	"sync"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client
type Client struct {
	Conn   *websocket.Conn
	Player models.Player
	RoomID string // Add RoomID to identify the room
}

// Message struct to hold the message and room ID
type Message struct {
	RoomID string
	Data   []byte
}

// Hub maintains the set of active rooms and broadcasts messages to clients in those rooms
type Hub struct {
	Rooms     map[string]map[*Client]bool // Map of room IDs to clients
	Broadcast chan Message                // Channel for broadcasting messages
	Mutex     sync.Mutex
}

// Exported Hub variable
var HubInstance = Hub{
	Rooms:     make(map[string]map[*Client]bool),
	Broadcast: make(chan Message),
}

// RegisterClient registers a new client to the hub
func RegisterClient(client *Client) {
	HubInstance.Mutex.Lock()
	if _, ok := HubInstance.Rooms[client.RoomID]; !ok {
		HubInstance.Rooms[client.RoomID] = make(map[*Client]bool)
	}
	HubInstance.Rooms[client.RoomID][client] = true
	HubInstance.Mutex.Unlock()
}

// UnregisterClient unregisters a client from the hub
func UnregisterClient(client *Client) {
	HubInstance.Mutex.Lock()
	if clients, ok := HubInstance.Rooms[client.RoomID]; ok {
		delete(clients, client)
		if len(clients) == 0 {
			delete(HubInstance.Rooms, client.RoomID) // Remove room if empty
		}
	}
	HubInstance.Mutex.Unlock()
}

// StartHub starts the hub to listen for broadcast messages
func StartHub() {
	for {
		// Wait for a message to be sent to the Broadcast channel
		msg := <-HubInstance.Broadcast

		// Extract the room ID and message data
		roomID := msg.RoomID
		data := msg.Data

		HubInstance.Mutex.Lock()
		// Check if there are clients in the specified room
		if clients, ok := HubInstance.Rooms[roomID]; ok {
			for client := range clients {
				// Attempt to send the message to each client in the room
				err := client.Conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					// If there's an error, unregister the client
					UnregisterClient(client)
					client.Conn.Close()
				}
			}
		}
		HubInstance.Mutex.Unlock()
	}
}
