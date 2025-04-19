package utils

import (
	"hearthstone-clone-backend/models"
)

func CreatePlayerGameState(gameState *models.GameState, playerID string) *models.PlayerGameState {
	isPlayer1 := (playerID == gameState.CurrentPlayer)

	playerState := &models.PlayerGameState{
		ID:            gameState.ID,
		RoomID:        gameState.RoomID,
		CurrentPlayer: gameState.CurrentPlayer,
		TurnNumber:    gameState.TurnNumber,
	}

	if isPlayer1 {
		playerState.MyHand = gameState.Player1Hand
		playerState.OpponentHand = len(gameState.Player2Hand)
		playerState.MyDeck = len(gameState.Player1Deck)
		playerState.OpponentDeck = len(gameState.Player2Deck)
		playerState.MyBoard = gameState.Player1Board
		playerState.OpponentBoard = gameState.Player2Board
	} else {
		playerState.MyHand = gameState.Player2Hand
		playerState.OpponentHand = len(gameState.Player1Hand)
		playerState.MyDeck = len(gameState.Player2Deck)
		playerState.OpponentDeck = len(gameState.Player1Deck)
		playerState.MyBoard = gameState.Player2Board
		playerState.OpponentBoard = gameState.Player1Board
	}

	return playerState
}
