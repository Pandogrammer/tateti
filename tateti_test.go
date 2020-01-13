package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCross(t *testing.T) {
	game := Tateti()

	updatedGame := game.Cross(4)

	assert.Equal(t, X, updatedGame.positions[4])
}

func TestCircle(t *testing.T) {
	game := Tateti()

	updatedGame := game.Circle(4)

	assert.Equal(t, O, updatedGame.positions[4])
}

func TestOccupiedSlot(t *testing.T) {
	game, _ := Tateti().Move(O, 4)

	_, err := game.Move(X, 4)

	assert.Error(t, err, OccupiedSlot)
}

func TestLastTurn(t *testing.T) {
	game, _ := Tateti().Move(O, 4)

	assert.Equal(t, O, game.lastTurn)
}

func TestLastTurnError(t *testing.T) {
	game, _ := Tateti().Move(O, 4)

	_, err := game.Move(O, 5)

	assert.Error(t, err, OtherPlayerTurn)
}

func TestWinCondition(t *testing.T) {
	game := Tateti().Circle(3).Circle(4).Circle(5)

	game, _ = game.Move(X, 6)

	assert.Equal(t, O, game.winner)
}
