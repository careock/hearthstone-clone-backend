package handlers

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"hearthstone-clone-backend/utils"
	"log"
)

func sendEvent(client *models.Client, event models.GameEvent) {
	message, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling event:", err)
		return
	}
	client.SendMessage(message)
}

func HandleCreateRoomEvent(client *models.Client, payload interface{}) {
	roomID := utils.GenerateRandomID()
	room := &models.Room{
		ID:      roomID,
		Clients: make(map[*models.Client]bool),
	}
	room.Clients[client] = true
	utils.Rooms[roomID] = room
	response := models.GameEvent{
		Type:    "roomCreated",
		Payload: roomID,
	}
	log.Println(utils.Rooms)
	log.Println(room)
	log.Printf("сlient:")
	log.Println(client)

	sendEvent(client, response)
}

func HandleJoinRoomEvent(client *models.Client, payload interface{}) {

	roomID := payload.(string)
	room := utils.Rooms[roomID]
	if room == nil {
		log.Printf("Room not found: %s", roomID)
		return
	}

	room.Clients[client] = true
	log.Printf("rooms:")
	log.Println(utils.Rooms)
	log.Printf("room:")
	log.Println(room)

	// Проверяем, что в комнате ровно 2 игрока
	if len(room.Clients) != 2 {
		log.Printf("Waiting for another player...")
		return
	}

	gameState := utils.StartGame(room)

	for clientInRoom := range room.Clients {
		playerState := utils.CreatePlayerGameState(gameState, clientInRoom.ID)
		sendEvent(clientInRoom, models.GameEvent{
			Type:    "gameState",
			Payload: playerState,
		})
	}
}

func HandlePlayCardEvent(client *models.Client, payload interface{}) {
	// Преобразуем payload в структуру
	payloadMap := payload.(map[string]interface{})
	playCardPayload := models.PlayCardPayload{
		CardID:    payloadMap["cardID"].(string),
		BoardSlot: int(payloadMap["boardSlot"].(float64)),
	}

	// Ищем активную игру, где этот клиент является текущим игроком
	var gameState *models.GameState
	for _, gs := range utils.GameStates {
		if gs.CurrentPlayer == client.ID {
			gameState = gs
			break
		}
	}

	if gameState == nil {
		log.Printf("No active game found for player: %s", client.ID)
		sendEvent(client, models.GameEvent{
			Type:    "error",
			Payload: "No active game found or not your turn",
		})
		return
	}

	// Играем карту
	updatedGameState, err := utils.PlayCard(gameState, client.ID, playCardPayload.CardID, playCardPayload.BoardSlot)
	if err != nil {
		log.Printf("Error playing card: %v", err)
		sendEvent(client, models.GameEvent{
			Type:    "error",
			Payload: err.Error(),
		})
		return
	}

	// Обновляем состояние игры
	utils.GameStates[updatedGameState.ID] = updatedGameState

	// Получаем комнату
	room := utils.Rooms[updatedGameState.RoomID]

	// Отправляем обновленное состояние каждому игроку
	for clientInRoom := range room.Clients {
		playerState := utils.CreatePlayerGameState(updatedGameState, clientInRoom.ID)
		sendEvent(clientInRoom, models.GameEvent{
			Type:    "gameState",
			Payload: playerState,
		})
	}
}
