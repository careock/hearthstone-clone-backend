// models/models.go

package models

type Player struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Health int    `json:"health"`
	Mana   int    `json:"mana"`
	Hand   []Card `json:"hand"`
	Board  []Card `json:"board"`
}

type Card struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Attack int    `json:"attack"`
	Health int    `json:"health"`
	Cost   int    `json:"cost"`
}

type Room struct {
	ID          string   `json:"id"`
	Players     []Player `json:"players"`
	InviteLink  string   `json:"inviteLink"`
	GameStarted bool     `json:"gameStarted"`
	CurrentTurn string   `json:"currentTurn"`
	Winner      string   `json:"winner"`
}
