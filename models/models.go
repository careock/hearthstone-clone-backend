// models/models.go
package models

type Card struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"` // "minion" or "spell"
	Attack   int    `json:"attack"`
	Health   int    `json:"health"`
	ManaCost int    `json:"mana_cost"`
}

type Player struct {
	ID     string `json:"id"`
	Health int    `json:"health"`
	Hand   []Card `json:"hand"`
	Board  []Card `json:"board"`
}

type Room struct {
	ID         string   `json:"id"`
	InviteLink string   `json:"invite_link"`
	Players    []Player `json:"players"` // Track players in the room
}
