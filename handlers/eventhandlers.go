package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"

	"github.com/gorilla/websocket"
)

func HandleCreateRoomEvent(client *models.Client, payload interface{}) {
	roomID := utils.GenerateRoomID()
	room := &models.Room{
		ID:      roomID,
		Clients: make(map[*models.Client]bool),
	}
	client.Room = room
	room.Clients[client] = true

	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: roomID,
	}
	sendEvent(client, response)
}

func sendEvent(client *models.Client, event models.GameEvent) {
	message, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("Error sending event:", err)
	}
}

// func broadcastEvent(room *models.Room, event models.GameEvent) {
// 	message, err := json.Marshal(event)
// 	if err != nil {
// 		log.Println("Error marshaling event:", err)
// 		return
// 	}

// 	for client := range room.Clients {
// 		err = client.Conn.WriteMessage(websocket.TextMessage, message)
// 		if err != nil {
// 			log.Println("Error broadcasting event:", err)
// 		}
// 	}
// }
