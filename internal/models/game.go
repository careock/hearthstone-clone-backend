package models

// Card represents a generic card in the game
type Card struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ManaCost    int    `json:"manaCost"`
	Type        string `json:"type"`
}

// Minion represents a minion card with attack and health
type Minion struct {
	Card
	Attack int `json:"attack"`
	Health int `json:"health"`
}

// GameEvent represents an event sent between client and server
type GameEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// GameState represents the current state of a game
type GameState struct {
	ID            string   `json:"id"`
	RoomID        string   `json:"roomID"`
	CurrentPlayer string   `json:"currentPlayer"`
	TurnNumber    int      `json:"turnNumber"`
	Player1Hand   []Card   `json:"player1Hand"`
	Player2Hand   []Card   `json:"player2Hand"`
	Player1Deck   []Card   `json:"player1Deck"`
	Player2Deck   []Card   `json:"player2Deck"`
	Player1Board  []Minion `json:"player1Board"`
	Player2Board  []Minion `json:"player2Board"`
}

// NewGameState creates a new game state for a room
func NewGameState(id, roomID, currentPlayer string) *GameState {
	return &GameState{
		ID:            id,
		RoomID:        roomID,
		CurrentPlayer: currentPlayer,
		TurnNumber:    1,
		Player1Hand:   make([]Card, 0),
		Player2Hand:   make([]Card, 0),
		Player1Deck:   make([]Card, 0),
		Player2Deck:   make([]Card, 0),
		Player1Board:  make([]Minion, 0),
		Player2Board:  make([]Minion, 0),
	}
}
