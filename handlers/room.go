package handlers

import (
	"hearthstone-clone-backend/models"
	"net/http"

	"github.com/google/uuid"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	// Generate a unique room ID
	roomID := uuid.New().String()

	// Create a new room
	room := models.Room{
		ID:         roomID,
		Players:    []models.Player{},
		InviteLink: "http://localhost:8080/join-game?roomId=" + roomID,
	}

	///// Store the room in the database (implementation not shown)

	// Return the invite link to the user
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(room.InviteLink))
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	// Get the roomId from the request parameters
	roomID := r.URL.Query().Get("roomId")

	// Retrieve the room from the database (implementation not shown)
	room := models.Room{ID: roomID}

	// Add the player to the room
	player := models.Player{
		ID:     uuid.New().String(),
		Name:   "Player " + string(len(room.Players)+1),
		Health: 3,
		Mana:   0,
		Hand:   []models.Card{},
		Board:  []models.Card{},
	}
	room.Players = append(room.Players, player)

	///// Update the room in the database (implementation not shown)

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
