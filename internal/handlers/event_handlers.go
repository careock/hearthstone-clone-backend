package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/internal/models"
	"hearthstone-clone-backend/internal/service"
	"log"
)

// EventHandler handles WebSocket events
type EventHandler struct {
	roomManager *service.RoomManager
}

// NewEventHandler creates a new EventHandler
func NewEventHandler(roomManager *service.RoomManager) *EventHandler {
	return &EventHandler{
		roomManager: roomManager,
	}
}

// sendEvent sends a GameEvent to a client
func sendEvent(client *models.Client, event models.GameEvent) {
	message, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}
	client.SendMessage(message)
}

// HandleCreateRoom handles the createRoom event
func (h *EventHandler) HandleCreateRoom(client *models.Client, payload interface{}) {
	room := h.roomManager.CreateRoom(client)

	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: room.ID,
	}

	log.Printf("Room created: %s, client: %s", room.ID, client.ID)
	sendEvent(client, response)
}

// HandleJoinRoom handles the joinRoom event
func (h *EventHandler) HandleJoinRoom(client *models.Client, payload interface{}) {
	// Safe type assertion for roomID
	roomID, ok := payload.(string)
	if !ok {
		log.Println("Invalid payload type for roomID, expected string")
		response := models.GameEvent{
			Type:    "error",
			Payload: "invalid room ID format",
		}
		sendEvent(client, response)
		return
	}

	room := h.roomManager.GetRoom(roomID)
	if room == nil {
		log.Printf("Room not found: %s", roomID)
		response := models.GameEvent{
			Type:    "error",
			Payload: "room not found",
		}
		sendEvent(client, response)
		return
	}

	// Check if room already has 2 players
	if room.ClientCount() >= 2 {
		log.Printf("Room %s is full", roomID)
		response := models.GameEvent{
			Type:    "error",
			Payload: "room is full",
		}
		sendEvent(client, response)
		return
	}

	room.AddClient(client)
	log.Printf("Client %s joined room %s", client.ID, roomID)

	// Start game when we have 2 players
	if room.ClientCount() == 2 {
		gameState := service.StartGame(room)
		room.BroadcastGameState(gameState)
	}
}

// HandlePlayCard handles the playCard event
func (h *EventHandler) HandlePlayCard(client *models.Client, payload interface{}) {
	log.Println("HandlePlayCard - TODO: Implement card playing logic")
	log.Printf("Client: %s, Payload: %+v", client.ID, payload)

	// TODO: Implement card playing logic
	// 1. Get the room the client is in
	// 2. Get the game state
	// 3. Validate the card play (mana cost, card in hand, etc.)
	// 4. Remove card from hand, add to board
	// 5. Update game state and broadcast
}
