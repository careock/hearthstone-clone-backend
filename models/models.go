// models/models.go
package models

type MessageType string

const (
	MessageTypeJoinRoom  MessageType = "join_room"
	MessageTypeLeaveRoom MessageType = "leave_room"
	MessageTypePlayCard  MessageType = "play_card"
	MessageTypeAttack    MessageType = "attack"
	MessageTypeEndTurn   MessageType = "end_turn"
	// Add more message types as needed
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}

type GameState struct {
	Rooms []Room
}

type Card struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ManaCost int    `json:"manaCost"`
	Attack   int    `json:"attack"`
	Health   int    `json:"health"`
	Text     string `json:"text"`
	Effect   Effect `json:"effect"`
}

type Effect struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	// Add more fields as needed
}

type Player struct {
	ID        string   `json:"id"`
	Health    int      `json:"health"`
	Mana      int      `json:"mana"`
	MaxMana   int      `json:"maxMana"`
	HeroPower string   `json:"heroPower"`
	Weapon    string   `json:"weapon"`
	Hand      []Card   `json:"hand"`
	Board     []Card   `json:"board"`
	Deck      []string `json:"deck"`
}

// Deck represents a collection of cards
type Deck struct {
	Cards []Card
}
