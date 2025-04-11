// handlers/eventhandlers.go
package handlers

import (
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
)

func handleJoinRoomEvent(client *utils.Client, payload interface{}) {
	roomID := payload.(string)
	room := utils.GetRoomByID(roomID)
	if room == nil {
		room = models.NewRoom()
		utils.AddRoom(room)
	}

	player := &models.Player{
		ID:      client.ID,
		Health:  30,
		Mana:    0,
		MaxMana: 0,
		Deck:    utils.BuildDeck(),
	}

	room.Players = append(room.Players, player)
	client.Player = player

	if len(room.Players) == 2 {
		startGame(room)
	}
}
