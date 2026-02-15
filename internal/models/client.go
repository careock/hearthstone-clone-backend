package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Client represents a connected WebSocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	mu   sync.Mutex
}

// SendMessage sends a message to the client in a thread-safe manner
func (c *Client) SendMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Conn.WriteMessage(websocket.TextMessage, message)
}
