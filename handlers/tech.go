package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"net/http"
)

func GetGameState(w http.ResponseWriter, r *http.Request) {
	gameState := models.GameState{
		Rooms: make([]models.Room, 0, len(rooms)),
	}

	for _, room := range rooms {
		gameState.Rooms = append(gameState.Rooms, room)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}
