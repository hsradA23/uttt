package main

import (
	"uttt/src/tt"
)

func main() {
	board := tt.Initialize()
	board.Move(tt.Player1, 0)
	board.Move(tt.Player2, 4)
	board.Move(tt.Player1, 1)
	board.Move(tt.Player2, 5)
	board.Move(tt.Player1, 2)
	board.Move(tt.Player2, 6)
}
