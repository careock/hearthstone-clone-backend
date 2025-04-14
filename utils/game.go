package utils

import (
	"hearthstone-clone-backend/models"
	"log"
	"math/rand"

	"github.com/google/uuid"
)

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

func createDeck() []models.Card {
	// Define the cards in the deck
	cards := []models.Card{
		{ID: "1", Name: "Card 1"},
		{ID: "1", Name: "Card 1"},
		{ID: "1", Name: "Card 1"},
	}

	// Shuffle the cards
	shuffledCards := make([]models.Card, len(cards))
	perm := rand.Perm(len(cards))
	for i, j := range perm {
		shuffledCards[i] = cards[j]
	}

	return shuffledCards
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

func StartGame(room *models.Room) {
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
	room.BroadcastGameState(gameState)
	log.Printf("gameState:")
	log.Println(gameState)
}
