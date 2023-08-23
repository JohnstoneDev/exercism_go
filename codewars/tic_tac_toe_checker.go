package main

import (
	"fmt"
)

// checks is a play(X || 0) has won the game (X == 1), (O == 2)
func CheckWin(board [3][3]int, play int) bool {
	// check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == play && board[i][1] == play && board[i][2] == play {
			return true
		}
	}

	// check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == play && board[1][i] == play && board[2][i] == play {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == play && board[1][1] == play && board[2][2] == play {
		return true
	}

	if board[0][2] == play && board[1][1] == play && board[2][0] == play {
		return true
	}

	return false
}

// checks that the game is complete (counts the number of empty spaces)
func IsFinished(board [3][3]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// checks the outcome of the game returns -1 if the board is not yet finished
// && no one has won, 1 if X won, 2 if 0 won, 0 if it's a draw
func IsSolved(board [3][3]int) int {
	playerXWin := CheckWin(board, 1);
	playerOWin := CheckWin(board, 2)

	if playerXWin {
		return 1
	} else if playerOWin {
		return 2
	} else if IsFinished(board) {
		return 0
	}

	return -1
}


func main() {
	win := [3][3]int{
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0},
	}

	unfinished := [3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 1, 0},
	}

	draw := [3][3]int{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 1},
	}

	fmt.Println(IsSolved(win), "<= Expecting 1")
	fmt.Println(IsSolved(unfinished), "<= Expecting -1")
	fmt.Println(IsSolved(draw), "<= Expecting 0")
}