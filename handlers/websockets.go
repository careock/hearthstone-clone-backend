// handlers/websocket.go
package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Get room ID from query parameters
	roomID := r.URL.Query().Get("roomId")
	if roomID == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	// Upgrade the connection to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Create a new client and register it to the room
	client := &utils.Client{Conn: conn, RoomID: roomID}
	utils.RegisterClient(client)

	// Listen for messages from the client
	for {
		// Read messages from the WebSocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break // Exit the loop if there's an error
		}

		// Broadcast the message to clients in the same room
		utils.HubInstance.Broadcast <- utils.Message{RoomID: roomID, Data: msg}

	}

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			utils.UnregisterClient(client)
			break
		}

		var message models.Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		switch message.Type {
		case models.MessageTypeJoinRoom:
			// Handle join room event
			// ...

		case models.MessageTypeLeaveRoom:
			// Handle leave room event
			// ...

		case models.MessageTypePlayCard:
			// Handle play card event
			// ...

		case models.MessageTypeAttack:
			// Handle attack event
			// ...

		case models.MessageTypeEndTurn:
			// Handle end turn event
			// ...

		// Add more cases for other message types

		default:
			log.Println("Unknown message type:", message.Type)
		}
	}

	// Unregister the client when done
	utils.UnregisterClient(client)
}
