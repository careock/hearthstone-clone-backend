package models

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// CardConfig структура для хранения конфигурации карт
type CardConfig struct {
	Cards   []Card   `json:"cards"`
	Minions []Minion `json:"minions"`
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

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func (c *Client) SendMessage(message []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, message)
}
