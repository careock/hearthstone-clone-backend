package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"
)

func HandleCreateRoomEvent(client *models.Client, payload interface{}) {
	roomID := utils.GenerateRoomID()
	room := &models.Room{
		ID:      roomID,
		Clients: make(map[*models.Client]bool),
	}
	log.Println(room)
	utils.Rooms[roomID] = room
	client.Room = room
	room.Clients[client] = true

	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: roomID,
	}
	sendEvent(client, response)
	log.Println(utils.Rooms)
}

func HandleJoinRoomEvent(client *models.Client, payload interface{}) {
	log.Println(utils.Rooms)
	roomID := payload.(string)
	room := utils.Rooms[roomID]
	if room == nil {
		log.Printf("Room not found: %s", roomID)
		return
	}

	gameState := &models.GameState{
		RoomID: room.ID,
	}
	log.Println(gameState)
	room.BroadcastGameState(gameState)
}

func sendEvent(client *models.Client, event models.GameEvent) {
	message, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}

	client.SendMessage(message)
}

// func broadcastEvent(room *models.Room, event models.GameEvent) {
// 	message, err := json.Marshal(event)
// 	log.Println(room)
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
