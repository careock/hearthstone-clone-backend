package handlers

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hearthstone-clone-backend/models"

	"github.com/google/uuid"
)

// NewDeck generates a new deck of cards
func NewDeck() (*models.Deck, error) {
	cards := []models.Card{
		{ID: uuid.New().String(), Name: "Fireball", ManaCost: 4, Type: "spell"},
		{ID: uuid.New().String(), Name: "Frostbolt", ManaCost: 2, Type: "spell"},
		{ID: uuid.New().String(), Name: "Grizzly Bear", ManaCost: 3, Type: "minion"},
		{ID: uuid.New().String(), Name: "Fire Elemental", ManaCost: 5, Type: "minion"},
	}

	// Shuffle the deck using crypto/rand
	if err := shuffleDeck(cards); err != nil {
		return nil, fmt.Errorf("failed to shuffle deck: %w", err)
	}

	return &models.Deck{Cards: cards}, nil
}

// shuffleDeck shuffles the slice of cards using crypto/rand
func shuffleDeck(cards []models.Card) error {
	for i := range cards {
		j, err := randomInt(i + 1)
		if err != nil {
			return err
		}
		cards[i], cards[j] = cards[j], cards[i]
	}
	return nil
}

// randomInt generates a random integer between 0 and max using crypto/rand
func randomInt(max int) (int, error) {
	var num uint32
	err := binary.Read(rand.Reader, binary.LittleEndian, &num)
	if err != nil {
		return 0, err
	}
	return int(num) % max, nil
}

func startGame(roomID string) {
	NewDeck()
}

func DrawCard() {
}
