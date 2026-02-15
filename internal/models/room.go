package models

import (
	"encoding/json"
	"log"
	"sync"
)

// Room represents a game room with connected clients
type Room struct {
	ID        string
	Clients   map[*Client]bool
	GameState *GameState
	mu        sync.RWMutex
}

// NewRoom creates a new room with the given ID
func NewRoom(id string) *Room {
	return &Room{
		ID:      id,
		Clients: make(map[*Client]bool),
	}
}

// AddClient adds a client to the room in a thread-safe manner
func (r *Room) AddClient(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Clients[client] = true
}

// RemoveClient removes a client from the room in a thread-safe manner
func (r *Room) RemoveClient(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Clients, client)
}

// ClientCount returns the number of clients in the room
func (r *Room) ClientCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.Clients)
}

// SetGameState sets the game state for the room
func (r *Room) SetGameState(state *GameState) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.GameState = state
}

// GetGameState returns the current game state
func (r *Room) GetGameState() *GameState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.GameState
}

// BroadcastGameState broadcasts the game state to all clients
func (r *Room) BroadcastGameState(gameState *GameState) {
	message, err := json.Marshal(gameState)
	if err != nil {
		log.Printf("Error marshaling game state: %v", err)
		return
	}
	r.BroadcastMessage(message)
}

// BroadcastMessage sends a message to all clients in the room
func (r *Room) BroadcastMessage(message []byte) {
	r.mu.RLock()
	clients := make([]*Client, 0, len(r.Clients))
	for client := range r.Clients {
		clients = append(clients, client)
	}
	r.mu.RUnlock()

	for _, client := range clients {
		err := client.SendMessage(message)
		if err != nil {
			log.Printf("Error sending message to client %s: %v", client.ID, err)
			client.Conn.Close()
			r.RemoveClient(client)
		}
	}
}
