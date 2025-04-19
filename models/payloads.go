package models

type GameEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type PlayCardPayload struct {
	CardID    string `json:"cardID"`
	BoardSlot int    `json:"boardSlot"`
}
