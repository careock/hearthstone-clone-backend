package models

import (
	"testing"
)

func TestNewRoom(t *testing.T) {
	room := NewRoom("test-room-id")

	if room == nil {
		t.Fatal("NewRoom() returned nil")
	}
	if room.ID != "test-room-id" {
		t.Errorf("Room ID = %s, want test-room-id", room.ID)
	}
	if room.Clients == nil {
		t.Error("Room Clients map is nil")
	}
	if len(room.Clients) != 0 {
		t.Errorf("New room should have 0 clients, got %d", len(room.Clients))
	}
}

func TestRoomAddClient(t *testing.T) {
	room := NewRoom("test-room")
	client := &Client{ID: "test-client"}

	room.AddClient(client)

	if room.ClientCount() != 1 {
		t.Errorf("Room should have 1 client, got %d", room.ClientCount())
	}
}

func TestRoomRemoveClient(t *testing.T) {
	room := NewRoom("test-room")
	client := &Client{ID: "test-client"}

	room.AddClient(client)
	room.RemoveClient(client)

	if room.ClientCount() != 0 {
		t.Errorf("Room should have 0 clients after removal, got %d", room.ClientCount())
	}
}

func TestRoomSetGameState(t *testing.T) {
	room := NewRoom("test-room")
	gameState := &GameState{ID: "game-1", RoomID: "test-room"}

	room.SetGameState(gameState)

	if room.GetGameState() != gameState {
		t.Error("GameState not set correctly")
	}
}

func TestNewGameState(t *testing.T) {
	gameState := NewGameState("game-1", "room-1", "player-1")

	if gameState.ID != "game-1" {
		t.Errorf("GameState ID = %s, want game-1", gameState.ID)
	}
	if gameState.RoomID != "room-1" {
		t.Errorf("GameState RoomID = %s, want room-1", gameState.RoomID)
	}
	if gameState.CurrentPlayer != "player-1" {
		t.Errorf("GameState CurrentPlayer = %s, want player-1", gameState.CurrentPlayer)
	}
	if gameState.TurnNumber != 1 {
		t.Errorf("GameState TurnNumber = %d, want 1", gameState.TurnNumber)
	}
}
