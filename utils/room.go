package utils

import (
	"hearthstone-clone-backend/models"
)

var Rooms = make(map[string]*models.Room)

func FindRoomByID(roomID string) *models.Room {
	return Rooms[roomID]
}
