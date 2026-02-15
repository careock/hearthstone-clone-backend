// models/models.go
package models

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ManaCost    int    `json:"manaCost"`
	Type        string `json:"type"`
}

type Minion struct {
	Card
	Attack int `json:"attack"`
	Health int `json:"health"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func (c *Client) SendMessage(message []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, message)
}

type GameEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type GameState struct {
	ID            string
	RoomID        string   `json:"roomID"`
	CurrentPlayer string   `json:"currentPlayer"`
	TurnNumber    int      `json:"turnNumber"`
	Player1Hand   []Card   `json:"player1Hand"`
	Player2Hand   []Card   `json:"player2Hand"`
	Player1Deck   []Card   `json:"player1Deck"`
	Player2Deck   []Card   `json:"player2Deck"`
	Player1Board  []Minion `json:"player1Board"`
	Player2Board  []Minion `json:"player2Board"`
}
