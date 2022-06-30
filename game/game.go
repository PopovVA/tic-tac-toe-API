package game

import (
	"errors"
	"sync"
)

const (
	FirstPlayer  = 1
	SecondPlayer = 2
)

var field = [3][3]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

var (
	ErrCellIsAlreadyTaken = errors.New("cell is already taken")
)

func canMove(column int, row int) error {
	cell := field[column][row]
	if cell > 0 {
		return ErrCellIsAlreadyTaken
	}
	return nil
}

func GetWinner() (player int) {
	var (
		winner int
		wg     sync.WaitGroup
	)

	wg.Add(3)

	// [1][1][1]
	go func() {
		for i := 0; i <= 2; i++ {
			rowSum := field[i][0] + field[i][1] + field[i][2]
			if rowSum == 3 {
				winner = FirstPlayer
			} else if rowSum == 6 {
				winner = SecondPlayer
			}
		}
		wg.Done()
	}()

	// [1]
	// [1]
	// [1]
	go func() {
		for i := 0; i <= 2; i++ {
			columnSum := field[0][i] + field[1][i] + field[2][i]
			if columnSum == 3 {
				winner = FirstPlayer
			} else if columnSum == 6 {
				winner = SecondPlayer
			}
		}
		wg.Done()
	}()

	//[1]
	//	[1]
	//		[1]

	//		[1]
	//	[1]
	//[1]
	go func() {
		leftDiagonalSum := field[0][0] + field[1][1] + field[2][2]
		rightDiagonalSum := field[2][2] + field[1][1] + field[0][0]

		if leftDiagonalSum == 3 || rightDiagonalSum == 3 {
			winner = FirstPlayer
		} else if leftDiagonalSum == 6 || rightDiagonalSum == 6 {
			winner = SecondPlayer
		}
		wg.Done()
	}()

	wg.Wait()

	return winner
}

func Move(player int, column int, row int) error {
	if err := canMove(column, row); err != nil {
		return err
	}
	field[column][row] = player
	return nil
}
