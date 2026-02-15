package service

import (
	"hearthstone-clone-backend/internal/models"
	"sync"
)

// RoomManager manages all game rooms with thread-safe operations
type RoomManager struct {
	rooms map[string]*models.Room
	mu    sync.RWMutex
}

// NewRoomManager creates a new RoomManager
func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*models.Room),
	}
}

// CreateRoom creates a new room and adds the client to it
func (rm *RoomManager) CreateRoom(client *models.Client) *models.Room {
	roomID := GenerateRandomID()
	room := models.NewRoom(roomID)
	room.AddClient(client)

	rm.mu.Lock()
	rm.rooms[roomID] = room
	rm.mu.Unlock()

	return room
}

// GetRoom retrieves a room by ID
func (rm *RoomManager) GetRoom(roomID string) *models.Room {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	return rm.rooms[roomID]
}

// RemoveRoom removes a room by ID
func (rm *RoomManager) RemoveRoom(roomID string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.rooms, roomID)
}

// RemoveClientFromRooms removes a client from all rooms they're in
func (rm *RoomManager) RemoveClientFromRooms(client *models.Client) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for _, room := range rm.rooms {
		room.RemoveClient(client)
	}
}

// RoomCount returns the number of active rooms
func (rm *RoomManager) RoomCount() int {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	return len(rm.rooms)
}
