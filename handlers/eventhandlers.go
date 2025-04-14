package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"
)

func sendEvent(client *models.Client, event models.GameEvent) {
	message, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}
	client.SendMessage(message)
}

func HandleCreateRoomEvent(client *models.Client, payload interface{}) {
	roomID := utils.GenerateRoomID()
	room := &models.Room{
		ID:      roomID,
		Clients: make(map[*models.Client]bool),
	}
	room.Clients[client] = true
	utils.Rooms[roomID] = room
	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: roomID,
	}
	sendEvent(client, response)

	log.Println(utils.Rooms)
	log.Println(room)
	log.Printf("сlient:")
	log.Println(client)
}

func HandleJoinRoomEvent(client *models.Client, payload interface{}) {
	roomID := payload.(string)
	room := utils.Rooms[roomID]
	if room == nil {
		log.Printf("Room not found: %s", roomID)
		return
	}

	room.Clients[client] = true
	log.Printf("rooms:")
	log.Println(utils.Rooms)
	log.Printf("room:")
	log.Println(room)
	startGame(room)
}

func startGame(room *models.Room) {
	gameState := &models.GameState{
		RoomID:        room.ID,
		CurrentPlayer: room.Clients[client],
	}

	room.BroadcastGameState(gameState)
	log.Printf("gameState:")
	log.Println(gameState)
}
