// handlers/room.go
package handlers

import (
	"encoding/json"
	"net/http"

	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"

	"github.com/google/uuid"
)

var rooms = make(map[string]models.Room)

// CreateRoom creates a new game room and generates an invite link
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	roomID := uuid.New().String()
	inviteLink := "http://localhost:8080/join-game?roomId=" + roomID

	room := models.Room{
		ID:         roomID,
		InviteLink: inviteLink,
		Players:    []models.Player{}, // Initialize empty player list
	}

	rooms[roomID] = room

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

// JoinGame allows a user to join an existing game room
func JoinGame(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("roomId")
	room, exists := rooms[roomID]

	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Check if the room already has two players
	if len(room.Players) >= 2 {
		http.Error(w, "Room is full", http.StatusForbidden)
		return
	}

	// Add the new player (for simplicity, we will use a UUID as player ID)
	playerID := uuid.New().String()
	newPlayer := models.Player{
		ID:     playerID,
		Health: 3, // Set initial health for the player
	}
	room.Players = append(room.Players, newPlayer)
	rooms[roomID] = room // Update the room in the map

	msg := []byte("Player " + playerID + " has joined the room.")
	utils.HubInstance.Broadcast <- utils.Message{RoomID: roomID, Data: msg}

	// Check if both players are in the room
	if len(room.Players) == 2 {
		startMessage := []byte("Game started")
		utils.HubInstance.Broadcast <- utils.Message{RoomID: roomID, Data: startMessage}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
