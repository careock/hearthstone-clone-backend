// models/models.go
package models

import "github.com/gorilla/websocket"

type Player struct {
	ID   string `json:"id"`
	Hand []Card `json:"hand"`
}

type Room struct {
	ID      string    `json:"id"`
	Players []*Player `json:"players"`
}

type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	Conn *websocket.Conn
	Room *Room
}

type GameEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
