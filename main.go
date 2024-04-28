package main

import (
	"fmt"
	"uttt/src/tt"
	utttengine "uttt/src/uttt-engine"
)

func main() {
	// board := tt.Initialize()
	// board.Move(tt.Player1, 0)
	// board.Move(tt.Player2, 4)
	// board.Move(tt.Player1, 1)
	// board.Move(tt.Player2, 5)
	// board.Move(tt.Player1, 2)
	// board.Move(tt.Player2, 6)
	game := utttengine.Initialize()
	player := tt.Player1
	for {
		j, k := 0, 0
		fmt.Println("Enter board and position for player", player)
		fmt.Scan(&j, &k)
		game.Move(player, j, k)
		if player == tt.Player1 {
			player = tt.Player2
		} else {
			player = tt.Player1
		}
	}
}
