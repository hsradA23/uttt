package tt

import (
	"fmt"
	"log"
)

// A board is a single array
// the indexes correspond to the following
// spots on the board
// 0 1 2
// 3 4 5
// 6 7 8

const (
	Empty = iota
	Player1
	Player2
)

type Board struct {
	arr        []int
	lastPlayer int
	winner     int
}

func Initialize() Board {
	return Board{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, Empty, Empty}
}

func (board Board) Print_board() {
	for i, v := range board.arr {
		fmt.Print(v)
		if i == 2 || i == 5 || i == 8 {
			fmt.Println("")
		}
	}
}

func (board *Board) Move(player, position int) {
	if board.winner != Empty {
		log.Println("Cannot play on a board with a winner.")
		return
	}

	if !(player == Player1 || player == Player2) {
		log.Println("Invalid value for player")
		return
	}

	if position < 0 || position > 8 {
		log.Println(fmt.Sprintf("Invalid position: %d", position))
		return
	}

	if board.arr[position] != Empty {
		log.Println("Invalid move, spot is already filled.")
		return
	}

	if player == board.lastPlayer {
		log.Println("Same player cannot play again")
		return
	}

	board.arr[position] = player
	board.lastPlayer = player
	log.Println(fmt.Sprintf("Moving player %d to position %d", player, position))

	if winner := board.check_win(); winner != Empty {
		board.winner = winner
		log.Println("Player", winner, "has won the game.")
	}
}

func (board Board) check_win() int {
	// 0 1 2
	// 3 4 5
	// 6 7 8
	for i := 0; i < 3; i++ {
		// Horizontal wins
		if board.arr[i] == board.arr[i+1] && board.arr[i+1] == board.arr[i+2] {
			return board.arr[i]
		}
	}
	for i := 0; i < 3; i++ {
		// Vertical wins
		if board.arr[i] == board.arr[i+3] && board.arr[i+3] == board.arr[i+6] {
			return board.arr[i]
		}
	}
	if board.arr[0] == board.arr[4] && board.arr[4] == board.arr[8] {
		return board.arr[4]
	}
	if board.arr[2] == board.arr[4] && board.arr[4] == board.arr[6] {
		return board.arr[4]
	}
	return Empty
}
