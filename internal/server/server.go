package server

import (
	"encoding/json"
	"hearthstone-clone-backend/internal/handlers"
	"hearthstone-clone-backend/internal/models"
	"hearthstone-clone-backend/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins in development
		// In production, you should restrict this
		return true
	},
}

// Server represents the WebSocket server
type Server struct {
	roomManager  *service.RoomManager
	eventHandler *handlers.EventHandler
}

// NewServer creates a new Server instance
func NewServer() *Server {
	roomManager := service.NewRoomManager()
	return &Server{
		roomManager:  roomManager,
		eventHandler: handlers.NewEventHandler(roomManager),
	}
}

// Start starts the WebSocket server on the specified address
func (s *Server) Start(addr string) error {
	http.HandleFunc("/ws", s.handleWebSocket)
	http.HandleFunc("/health", s.handleHealth)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleHealth handles health check requests
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// handleWebSocket handles WebSocket connections
func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	clientID := service.GenerateRandomID()
	client := &models.Client{
		ID:   clientID,
		Conn: conn,
	}

	log.Printf("Client connected: %s", clientID)
	defer s.cleanupClient(client)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		var event models.GameEvent
		if err := json.Unmarshal(message, &event); err != nil {
			log.Printf("Error unmarshaling event: %v", err)
			continue
		}

		s.handleEvent(client, event)
	}
}

// handleEvent routes events to appropriate handlers
func (s *Server) handleEvent(client *models.Client, event models.GameEvent) {
	switch event.Type {
	case "createRoom":
		s.eventHandler.HandleCreateRoom(client, event.Payload)
	case "joinRoom":
		s.eventHandler.HandleJoinRoom(client, event.Payload)
	case "playCard":
		s.eventHandler.HandlePlayCard(client, event.Payload)
	default:
		log.Printf("Unknown event type: %s", event.Type)
	}
}

// cleanupClient removes a client from all rooms and closes the connection
func (s *Server) cleanupClient(client *models.Client) {
	log.Printf("Client disconnected: %s", client.ID)
	s.roomManager.RemoveClientFromRooms(client)
	client.Conn.Close()
}
