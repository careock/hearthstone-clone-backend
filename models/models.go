// models/models.go
package models

type GameState struct {
	Rooms []Room
}

type Card struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"` // "minion" or "spell"
	ManaCost int    `json:"mana_cost"`
}

type Player struct {
	ID     string `json:"id"`
	Health int    `json:"health"`
	Hand   []Card `json:"hand"`
	Board  []Card `json:"board"`
	Deck   Deck   `json:"deck"`
}

type Room struct {
	ID         string   `json:"id"`
	InviteLink string   `json:"invite_link"`
	Players    []Player `json:"players"` // Track players in the room
}

// Deck represents a collection of cards
type Deck struct {
	Cards []Card
}
