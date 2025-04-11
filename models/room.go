// models/room.go
package models

import "github.com/google/uuid"

type Room struct {
	ID             string    `json:"id"`
	Players        []*Player `json:"players"`
	CurrentTurn    int       `json:"currentTurn"`
	ActivePlayerID string    `json:"activePlayerID"`
	GamePhase      string    `json:"gamePhase"`
}

func NewRoom() *Room {
	return &Room{
		ID:          uuid.New().String(),
		Players:     make([]*Player, 0),
		CurrentTurn: 0,
		GamePhase:   "waiting",
	}
}
