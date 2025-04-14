package utils

import (
	"hearthstone-clone-backend/models"
	"math/rand/v2"
)

func DrawInitialHand(player *models.Client, deck []models.Card) {
	handSize := 3 // Adjust the initial hand size as needed
	player.Hand = make([]models.Card, handSize)
	for i := 0; i < handSize; i++ {
		if len(deck) > 0 {
			player.Hand[i] = deck[0]
			deck = deck[1:]
		}
	}
}

func CreateDeck() []models.Card {
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
