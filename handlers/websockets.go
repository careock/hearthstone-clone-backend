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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := &utils.Client{Conn: conn}
	utils.RegisterClient(client)

	// Listen for messages from the client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// Broadcast the message to all clients
		utils.HubInstance.Broadcast <- msg
	}

	utils.UnregisterClient(client)
}
