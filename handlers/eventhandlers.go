package handlers

import (
	"fmt"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"net/http"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	room := utils.CreateRoom()
	// Return the room ID to the client
	fmt.Fprintf(w, room.ID)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("roomID")
	playerID := r.URL.Query().Get("playerID")
	// Find the room by ID
	room := utils.FindRoomByID(roomID)
	if room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	// Create a new player
	player := &models.Player{ID: playerID}
	// Join the room
	utils.JoinRoom(room, player)
	// Draw the initial hand for the player
	deck := utils.CreateDeck()
	utils.DrawInitialHand(player, deck)
	// Return a success response to the client
	fmt.Fprintf(w, "Joined room successfully")
}
