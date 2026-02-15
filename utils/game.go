package utils

import (
	"encoding/json"
	"hearthstone-clone-backend/models"
	"log"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

// Config structures for loading cards.json
type CardConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ManaCost    int    `json:"manaCost"`
	Type        string `json:"type"`
}

type MinionConfig struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`
	Attack int    `json:"attack"`
	Health int    `json:"health"`
}

type CardsConfig struct {
	Cards   []CardConfig   `json:"cards"`
	Minions []MinionConfig `json:"minions"`
}

var loadedConfig *CardsConfig

// LoadCardsConfig loads the cards configuration from configs/cards.json
func LoadCardsConfig() (*CardsConfig, error) {
	if loadedConfig != nil {
		return loadedConfig, nil
	}

	file, err := os.Open("configs/cards.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config CardsConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	loadedConfig = &config
	return loadedConfig, nil
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

func createDeck() []models.Card {
	config, err := LoadCardsConfig()
	if err != nil {
		log.Printf("Error loading cards config: %v", err)
		return []models.Card{}
	}

	// Create a map of minion stats by cardID for quick lookup
	minionStats := make(map[string]MinionConfig)
	for _, minion := range config.Minions {
		minionStats[minion.CardID] = minion
	}

	// Build deck with multiple copies of each minion
	var cards []models.Card
	for _, cardConfig := range config.Cards {
		if cardConfig.Type == "minion" {
			// Add 2 copies of each minion to the deck
			for i := 0; i < 2; i++ {
				card := models.Card{
					ID:          cardConfig.ID,
					Name:        cardConfig.Name,
					Description: cardConfig.Description,
					ManaCost:    cardConfig.ManaCost,
					Type:        cardConfig.Type,
				}
				cards = append(cards, card)
			}
		}
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
