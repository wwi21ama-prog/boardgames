package main

import (
	"fmt"

	b "github.com/wwi21ama-prog/boardgames/board"
)

func main() {
	fmt.Println("Leeres TicTacToe-Board (Größe 3x3):")
	fmt.Println()

	board := b.MakeBoard(3, " ")
	b.PrintBoard(board)
}
