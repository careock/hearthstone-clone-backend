// utils/websocket.go
package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	RoomID string
}

type Hub struct {
	Clients   map[*Client]bool
	Broadcast chan []byte
	Mutex     sync.Mutex
}

// Exported Hub variable
var HubInstance = Hub{
	Clients:   make(map[*Client]bool),
	Broadcast: make(chan []byte),
}

// RegisterClient registers a new client to the hub
func RegisterClient(client *Client) {
	HubInstance.Mutex.Lock()
	HubInstance.Clients[client] = true
	HubInstance.Mutex.Unlock()
}

// UnregisterClient unregisters a client from the hub
func UnregisterClient(client *Client) {
	HubInstance.Mutex.Lock()
	delete(HubInstance.Clients, client)
	HubInstance.Mutex.Unlock()
}

// StartHub starts the hub to listen for broadcast messages
func StartHub() {
	for {
		msg := <-HubInstance.Broadcast
		HubInstance.Mutex.Lock()
		for client := range HubInstance.Clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				UnregisterClient(client)
				client.Conn.Close()
			}
		}
		HubInstance.Mutex.Unlock()
	}
}
