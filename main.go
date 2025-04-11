// main.go
package main

import (
	"encoding/json"
	"hearthstone-clone-backend/handlers"
	"hearthstone-clone-backend/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	client := &models.Client{Conn: conn}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var event models.GameEvent
		err = json.Unmarshal(message, &event)
		if err != nil {
			log.Println("Error unmarshaling event:", err)
			continue
		}

		switch event.Type {
		case "createRoom":
			handlers.HandleCreateRoomEvent(client, event.Payload)
		case "joinRoom":
			handlers.HandleJoinRoomEvent(client, event.Payload)
		default:
			log.Println("Unknown event type:", event.Type)
		}
	}
}
