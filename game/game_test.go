package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	//Valid move
	field = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	err := Move(1, 1, 2)
	assert.NoError(t, err, "Must be no error when move is valid")

	//Invalid move
	field = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	err = Move(1, 0, 0)
	assert.Error(t, err, "Must be error when move is invalid")
}

func TestGetWinner(t *testing.T) {
	// Vertical winner
	Move(1, 0, 0)
	Move(1, 1, 0)
	Move(1, 2, 0)
	winner := GetWinner()
	assert.Equal(t, FirstPlayer, winner, "Must be a vertical victory")

	field = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	// Horizontal winner
	Move(1, 0, 0)
	Move(1, 0, 1)
	Move(1, 0, 2)
	winner = GetWinner()
	assert.Equal(t, FirstPlayer, winner, "Must be a horizontal victory")

	field = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	// Diagonal winner
	Move(1, 0, 0)
	Move(1, 1, 1)
	Move(1, 2, 2)
	winner = GetWinner()
	assert.Equal(t, FirstPlayer, winner, "Must be a diagonal victory")

	field = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	// Has no winner
	Move(1, 0, 0)
	Move(1, 0, 1)
	Move(1, 1, 0)
	noWinner := GetWinner()
	assert.Equal(t, 0, noWinner, "Must be no victory")
}
