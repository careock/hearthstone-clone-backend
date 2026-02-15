package service

import (
	"encoding/json"
	"hearthstone-clone-backend/internal/models"
	"log"
	"math/rand"
	"os"
)

// CardConfig represents a card definition from cards.json
type CardConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ManaCost    int    `json:"manaCost"`
	Type        string `json:"type"`
}

// MinionConfig represents a minion definition from cards.json
type MinionConfig struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`
	Attack int    `json:"attack"`
	Health int    `json:"health"`
}

// CardsConfig represents the structure of cards.json
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

// CreateDeck creates a shuffled deck of cards
func CreateDeck() []models.Card {
	config, err := LoadCardsConfig()
	if err != nil {
		log.Printf("Error loading cards config: %v", err)
		return []models.Card{}
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

// SelectRandomPlayer selects a random player from the room
func SelectRandomPlayer(room *models.Room) string {
	clients := make([]*models.Client, 0)
	for client := range room.Clients {
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return ""
	}

	randomIndex := rand.Intn(len(clients))
	return clients[randomIndex].ID
}

// DrawInitialHand draws initial hands for both players
func DrawInitialHand(gameState *models.GameState, handSize int) {
	// Draw cards for player 1
	for i := 0; i < handSize && len(gameState.Player1Deck) > 0; i++ {
		randomIndex := rand.Intn(len(gameState.Player1Deck))
		card := gameState.Player1Deck[randomIndex]
		gameState.Player1Hand = append(gameState.Player1Hand, card)
		gameState.Player1Deck = append(gameState.Player1Deck[:randomIndex], gameState.Player1Deck[randomIndex+1:]...)
	}

	// Draw cards for player 2
	for i := 0; i < handSize && len(gameState.Player2Deck) > 0; i++ {
		randomIndex := rand.Intn(len(gameState.Player2Deck))
		card := gameState.Player2Deck[randomIndex]
		gameState.Player2Hand = append(gameState.Player2Hand, card)
		gameState.Player2Deck = append(gameState.Player2Deck[:randomIndex], gameState.Player2Deck[randomIndex+1:]...)
	}
}

// StartGame initializes a new game for the room
func StartGame(room *models.Room) *models.GameState {
	currentPlayer := SelectRandomPlayer(room)

	gameState := models.NewGameState(
		GenerateRandomID(),
		room.ID,
		currentPlayer,
	)
	gameState.Player1Deck = CreateDeck()
	gameState.Player2Deck = CreateDeck()

	DrawInitialHand(gameState, 3)

	room.SetGameState(gameState)

	log.Printf("Game started: %+v", gameState)
	return gameState
}
