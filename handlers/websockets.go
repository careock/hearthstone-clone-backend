// handlers/websocket.go
package handlers

import (
	"hearthstone-clone-backend/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
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

	// Unregister the client when done
	utils.UnregisterClient(client)
}
