// // handlers/room.go
// package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"hearthstone-clone-backend/models"
// 	"hearthstone-clone-backend/utils"

// 	"github.com/google/uuid"
// )

// var rooms = make(map[string]models.Room)

// // CreateRoom creates a new game room and generates an invite link
// func CreateRoom(w http.ResponseWriter, r *http.Request) {
// 	roomID := uuid.New().String()
// 	inviteLink := "http://localhost:8080/join-game?roomId=" + roomID

// 	room := models.Room{
// 		ID:         roomID,
// 		InviteLink: inviteLink,
// 		Players:    []models.Player{}, // Initialize empty player list
// 	}

// 	rooms[roomID] = room

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(room)
// }

// handlers/room.go
package handlers
