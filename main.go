package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"
)

func handleWebSocketMessage(client *models.Client, message []byte) {
	var event models.GameEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		log.Printf("Error unmarshaling event: %v", err)
		return
	}

	switch event.Type {
	case "createRoom":
		handleCreateRoomEvent(client, event.Payload)
	case "joinRoom":
		handleJoinRoomEvent(client, event.Payload)
	default:
		log.Printf("Unknown event type: %s", event.Type)
	}
}

func handleCreateRoomEvent(client *models.Client, payload interface{}) {
	roomID := utils.GenerateRoomID()
	room := &models.Room{
		ID:      roomID,
		Players: []*models.Player{},
	}
	utils.Rooms[roomID] = room

	// Send the room ID back to the client
	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: roomID,
	}
	client.SendMessage(response)
}

func handleJoinRoomEvent(client *utils.Client, payload interface{}) {
	roomID := payload.(string)
	room := utils.Rooms[roomID]
	if room == nil {
		log.Printf("Room not found: %s", roomID)
		return
	}

	player := &models.Player{
		ID:   client.ID,
		Deck: utils.CreateDeck(),
	}
	room.Players = append(room.Players, player)
	client.Player = player
	client.Room = room

	// Deal the initial hand
	dealInitialHand(player)

	// Broadcast the updated game state to all players in the room
	broadcastGameState(room)
}

func dealInitialHand(player *models.Player) {
	handSize := 3 // Adjust the initial hand size as needed
	for i := 0; i < handSize; i++ {
		if len(player.Deck) > 0 {
			card := player.Deck[0]
			player.Deck = player.Deck[1:]
			player.Hand = append(player.Hand, card)
		}
	}
}

func broadcastGameState(room *models.Room) {
	gameState := &models.GameState{
		Players: room.Players,
	}

	stateJSON, err := json.Marshal(gameState)
	if err != nil {
		log.Printf("Error marshaling game state: %v", err)
		return
	}

	for _, player := range room.Players {
		client := utils.GetClientByID(player.ID)
		if client != nil {
			client.SendMessage(stateJSON)
		}
	}
}
