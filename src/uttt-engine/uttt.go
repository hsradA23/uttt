package utttengine

import (
	"log"
	"uttt/src/tt"
)

type Game struct {
	boards     []tt.Board
	lastPlayer int
	nextBoard  int
}

func Initialize() Game {
	game := Game{}
	for i := 0; i < 9; i++ {
		game.boards = append(game.boards, tt.Initialize())
	}
	game.nextBoard = -1
	game.lastPlayer = tt.Empty
	return game
}

func (game *Game) Move(player, board, position int) {
	if player == game.lastPlayer {
		log.Println("Same player cannot play twice in a row.")
		return
	}

	if board < 0 || board > 8 {
		log.Println("Invlid value for board.")
		return
	}

	if game.nextBoard != -1 && board != game.nextBoard {
		log.Println("Can only play on the board", game.nextBoard)
		return
	}

	if game.boards[board].Winner != tt.Empty {
		log.Println("Cannot move on a finished board")
		return
	}

	log.Println("Board", board)
	game.boards[board].Move(player, position)

	if game.boards[position].Winner == tt.Empty {
		game.nextBoard = position
	} else {
		game.nextBoard = -1
	}
	game.lastPlayer = player
}
