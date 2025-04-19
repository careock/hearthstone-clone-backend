// models/models.go
package models

type Card struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ManaCost    int    `json:"manaCost"`
	Type        string `json:"type"` // minion, spell
}

type Minion struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`
	Attack int    `json:"attack"`
	Health int    `json:"health"`
}

type Spell struct {
	ID     string `json:"id"`
	CardID string `json:"cardID"`
}

type GameState struct {
	ID            string
	RoomID        string   `json:"roomID"`
	CurrentPlayer string   `json:"currentPlayer"`
	TurnNumber    int      `json:"turnNumber"`
	Player1Hand   []Card   `json:"player1Hand"`
	Player2Hand   []Card   `json:"player2Hand"`
	Player1Deck   []Card   `json:"player1Deck"`
	Player2Deck   []Card   `json:"player2Deck"`
	Player1Board  []Minion `json:"player1Board"`
	Player2Board  []Minion `json:"player2Board"`
	Player1Mana   int      `json:"player1Mana"`
	Player2Mana   int      `json:"player2Mana"`
}

type PlayerGameState struct {
	ID            string   `json:"id"`
	RoomID        string   `json:"roomID"`
	CurrentPlayer string   `json:"currentPlayer"`
	TurnNumber    int      `json:"turnNumber"`
	MyHand        []Card   `json:"myHand"`
	OpponentHand  int      `json:"opponentHandSize"` // только количество карт
	MyDeck        int      `json:"myDeckSize"`       // только количество карт
	OpponentDeck  int      `json:"opponentDeckSize"` // только количество карт
	MyBoard       []Minion `json:"myBoard"`
	OpponentBoard []Minion `json:"opponentBoard"`
}
