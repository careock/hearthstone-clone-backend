package utils

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"log"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

var Rooms = make(map[string]*models.Room)

func FindRoomByID(roomID string) *models.Room {
	return Rooms[roomID]
}

func GenerateRandomID() string {
	return uuid.New().String()
}

func drawInitialHand(gameState *models.GameState) {
	handSizePlayer1 := 3
	handSizePlayer2 := 3

	// Тянем карты для первого игрока
	gameState.Player1Hand = make([]models.Card, handSizePlayer1)
	for i := 0; i < handSizePlayer1; i++ {
		if len(gameState.Player1Deck) > 0 {
			randomIndex := rand.Intn(len(gameState.Player1Deck))
			card := gameState.Player1Deck[randomIndex]
			gameState.Player1Hand[i] = card
			gameState.Player1Deck = append(gameState.Player1Deck[:randomIndex], gameState.Player1Deck[randomIndex+1:]...)
		}
	}

	// Тянем карты для второго игрока
	gameState.Player2Hand = make([]models.Card, handSizePlayer2)
	for i := 0; i < handSizePlayer2; i++ {
		if len(gameState.Player2Deck) > 0 {
			randomIndex := rand.Intn(len(gameState.Player2Deck))
			card := gameState.Player2Deck[randomIndex]
			gameState.Player2Hand[i] = card
			gameState.Player2Deck = append(gameState.Player2Deck[:randomIndex], gameState.Player2Deck[randomIndex+1:]...)
		}
	}

}

// loadCardConfig загружает конфигурацию карт из JSON файла
func loadCardConfig() (*models.CardConfig, error) {
	file, err := os.ReadFile("configs/cards_config.json")
	if err != nil {
		return nil, err
	}

	var config models.CardConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	log.Println(config)
	return &config, nil
}

func createDeck() []models.Card {
	config, err := loadCardConfig()
	if err != nil {
		log.Printf("Error loading card config: %v", err)
		return []models.Card{} // Возвращаем пустую колоду в случае ошибки
	}

	// Создаем колоду из всех карт (по 2 копии каждой)
	deck := make([]models.Card, 0, len(config.Cards)*2)
	for _, card := range config.Cards {
		// Добавляем две копии каждой карты
		deck = append(deck, card, card)
	}

	// Перемешиваем колоду
	shuffledDeck := make([]models.Card, len(deck))
	perm := rand.Perm(len(deck))
	for i, j := range perm {
		shuffledDeck[i] = deck[j]
	}
	log.Println(shuffledDeck)
	return shuffledDeck
}

func selectRandomPlayer(room *models.Room) string {
	// Преобразуем карту клиентов в слайс
	clients := make([]*models.Client, 0, len(room.Clients))
	for client := range room.Clients {
		clients = append(clients, client)
	}

	// Выбираем случайного игрока
	randomIndex := rand.Intn(len(room.Clients))
	currentPlayer := clients[randomIndex]

	return currentPlayer.ID
}

func StartGame(room *models.Room) *models.GameState {
	currentPlayer := selectRandomPlayer(room)

	//объявляем начальное состояние игры (в этой комнате)
	gameState := &models.GameState{
		ID:            GenerateRandomID(),
		RoomID:        room.ID,
		CurrentPlayer: currentPlayer,
		TurnNumber:    1,
		Player1Deck:   createDeck(),
		Player2Deck:   createDeck(),
	}
	drawInitialHand(gameState)
	log.Printf("gameState:")
	log.Println(gameState)
	return gameState
}

// PlayCardResult результат попытки сыграть карту
type PlayCardResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func PlayCard(gameState *models.GameState, cardID string, playerID string, boardSlot int) *PlayCardResult {
	var playerHand *[]models.Card
	var playerBoard *[]models.Card

	// Определяем, чья очередь и проверяем, что это текущий игрок
	if playerID != gameState.CurrentPlayer {
		return &PlayCardResult{
			Success: false,
			Message: "It's not your turn",
		}
	}

	// Определяем руку и доску игрока
	if playerID == gameState.CurrentPlayer {
		playerHand = &gameState.Player1Hand
		playerBoard = &gameState.Player1Board
	} else {
		playerHand = &gameState.Player2Hand
		playerBoard = &gameState.Player2Board
	}

	// Ищем карту в руке игрока
	var cardIndex int = -1
	var playedCard models.Card
	for i, card := range *playerHand {
		if card.ID == cardID {
			cardIndex = i
			playedCard = card
			break
		}
	}

	if cardIndex == -1 {
		return &PlayCardResult{
			Success: false,
			Message: "Card not found in hand",
		}
	}

	switch playedCard.Type {
	case "minion":
		// Проверяем, есть ли место на доске
		if len(*playerBoard) >= 7 {
			return &PlayCardResult{
				Success: false,
				Message: "Board is full",
			}
		}

		// Проверяем валидность позиции
		if boardSlot < 0 || boardSlot > len(*playerBoard) {
			return &PlayCardResult{
				Success: false,
				Message: "Invalid board position",
			}
		}

		// Добавляем миньона на доску в указанную позицию
		*playerBoard = append((*playerBoard)[:boardSlot], append([]models.Card{playedCard}, (*playerBoard)[boardSlot:]...)...)

		// Удаляем карту из руки
		*playerHand = append((*playerHand)[:cardIndex], (*playerHand)[cardIndex+1:]...)

		return &PlayCardResult{
			Success: true,
			Message: "Minion summoned successfully",
		}
	}

	return &PlayCardResult{
		Success: true,
		Message: "Card played successfully",
	}
}
