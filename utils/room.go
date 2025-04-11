package utils

import (
	"hearthstone-clone-backend/models"

	"github.com/google/uuid"
)

func GenerateRoomID() string {
	// Generate a unique room ID (you can use a library like github.com/google/uuid)
	return uuid.New().String()
}

var Rooms = make(map[string]*models.Room)

func FindRoomByID(roomID string) *models.Room {
	return Rooms[roomID]
}
