package main

import (
	"errors"
)

type Symbol int

const (
	X     Symbol = 1
	O     Symbol = 2
	EMPTY Symbol = 0
)

const (
	OtherPlayerTurn string = "OtherPlayerTurn"
	OccupiedSlot    string = "OccupiedSlot"
)

var PossibleWins = [...][3]int{
	//horizontal
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},

	//vertical
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},

	//crossed
	{0, 4, 8},
	{6, 4, 2},
}

func (game tateti) Circle(position int) tateti {
	return game.putSymbol(O, position)
}

func (game tateti) Cross(position int) tateti {
	return game.putSymbol(X, position)
}

func (game tateti) Move(symbol Symbol, position int) (tateti, error) {
	_, err := game.validate(position, symbol)
	if err != nil {
		return game, err
	}

	updatedGame := game.putSymbol(symbol, position)
	updatedGame.winner = updatedGame.validateWinner()
	updatedGame.lastTurn = symbol
	return updatedGame, nil
}

func (game tateti) validate(position int, symbol Symbol) (tateti, error) {
	err := game.checkPosition(position)
	if err != nil {
		return game, err
	}
	err = game.checkTurn(symbol)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (game tateti) checkPosition(position int) error {
	if game.positions[position] != EMPTY {
		return errors.New(OccupiedSlot)
	}
	return nil
}

func (game tateti) putSymbol(symbol Symbol, position int) tateti {
	game.positions[position] = symbol
	return game
}

func (game tateti) checkTurn(symbol Symbol) error {
	if game.lastTurn == symbol {
		return errors.New(OtherPlayerTurn)
	}
	return nil
}

func (game tateti) validateWinner() Symbol {
	winner := EMPTY
	possibleWin := 0
	for winner == EMPTY && possibleWin < len(PossibleWins) {
		winner = game.validateLine(PossibleWins[possibleWin])
		possibleWin++
	}
	return winner
}

func (game tateti) validateLine(line [3]int) Symbol {
	var a = game.positions[line[0]]
	var b = game.positions[line[1]]
	var c = game.positions[line[2]]

	if a == b && b == c {
		return a
	}
	return EMPTY
}

func Tateti() tateti {
	return tateti{}
}

type tateti struct {
	positions [9]Symbol
	lastTurn  Symbol
	winner    Symbol
}
