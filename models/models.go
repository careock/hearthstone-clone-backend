// models/models.go
package models

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID   string `json:"id"`
	Deck []Card `json:"deck"`
	Hand []Card `json:"hand"`
}

type Room struct {
	ID      string
	Clients map[*Client]bool
}

func (r *Room) BroadcastGameState(gameState *GameState) {
	message, err := json.Marshal(gameState)
	if err != nil {
		log.Printf("Error marshaling game state: %v", err)
		return
	}
	r.BroadcastMessage(message)
}

func (r *Room) BroadcastMessage(message []byte) {
	for client := range r.Clients {
		err := client.SendMessage(message)
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			client.Conn.Close()
			delete(r.Clients, client)
		}
	}
}

type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	Conn *websocket.Conn
	Room *Room
}

func (c *Client) SendMessage(message []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, message)
}

type GameEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type GameState struct {
	RoomID        string    `json:"roomID"`
	Players       []*Player `json:"players"`
	CurrentPlayer string    `json:"currentPlayer"`
	TurnNumber    int       `json:"turnNumber"`
	// Add other relevant game state properties
}
