package service

import (
	"hearthstone-clone-backend/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestNewRoomManager(t *testing.T) {
	rm := NewRoomManager()
	if rm == nil {
		t.Fatal("NewRoomManager() returned nil")
	}
	if rm.rooms == nil {
		t.Error("RoomManager rooms map is nil")
	}
	if rm.RoomCount() != 0 {
		t.Errorf("New RoomManager should have 0 rooms, got %d", rm.RoomCount())
	}
}

func TestCreateRoom(t *testing.T) {
	rm := NewRoomManager()
	client := &models.Client{ID: "test-client"}

	room := rm.CreateRoom(client)

	if room == nil {
		t.Fatal("CreateRoom() returned nil")
	}
	if room.ID == "" {
		t.Error("Room ID should not be empty")
	}
	if room.ClientCount() != 1 {
		t.Errorf("Room should have 1 client, got %d", room.ClientCount())
	}
	if rm.RoomCount() != 1 {
		t.Errorf("RoomManager should have 1 room, got %d", rm.RoomCount())
	}
}

func TestGetRoom(t *testing.T) {
	rm := NewRoomManager()
	client := &models.Client{ID: "test-client"}
	room := rm.CreateRoom(client)

	// Test getting existing room
	foundRoom := rm.GetRoom(room.ID)
	if foundRoom == nil {
		t.Error("GetRoom() returned nil for existing room")
	}
	if foundRoom.ID != room.ID {
		t.Errorf("GetRoom() returned wrong room, expected %s, got %s", room.ID, foundRoom.ID)
	}

	// Test getting non-existent room
	notFound := rm.GetRoom("non-existent-id")
	if notFound != nil {
		t.Error("GetRoom() should return nil for non-existent room")
	}
}

func TestRemoveRoom(t *testing.T) {
	rm := NewRoomManager()
	client := &models.Client{ID: "test-client"}
	room := rm.CreateRoom(client)

	rm.RemoveRoom(room.ID)

	if rm.RoomCount() != 0 {
		t.Errorf("RoomManager should have 0 rooms after removal, got %d", rm.RoomCount())
	}
	if rm.GetRoom(room.ID) != nil {
		t.Error("Room should not exist after removal")
	}
}

func createTestWebSocketConnection(t *testing.T) (*models.Client, *httptest.Server) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("WebSocket upgrade error: %v", err)
		}
	}))

	wsURL := "ws" + server.URL[4:]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}

	client := &models.Client{
		ID:   GenerateRandomID(),
		Conn: conn,
	}

	return client, server
}

func TestRemoveClientFromRooms(t *testing.T) {
	rm := NewRoomManager()
	client, server := createTestWebSocketConnection(t)
	defer server.Close()
	defer client.Conn.Close()

	room := rm.CreateRoom(client)
	if room.ClientCount() != 1 {
		t.Errorf("Room should have 1 client, got %d", room.ClientCount())
	}

	rm.RemoveClientFromRooms(client)

	if room.ClientCount() != 0 {
		t.Errorf("Room should have 0 clients after removal, got %d", room.ClientCount())
	}
}
