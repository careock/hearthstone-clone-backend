// models/models.go
package models

import "github.com/gorilla/websocket"

type Player struct {
	ID   string `json:"id"`
	Deck []Card `json:"deck"`
	Hand []Card `json:"hand"`
}

type Room struct {
	ID      string
	Clients map[*Client]bool
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
