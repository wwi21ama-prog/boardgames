package main

import (
	"fmt"

	b "github.com/wwi21ama-prog/boardgames/board"
)

func main() {
	{
		fmt.Println("Leeres TicTacToe-Board (Größe 3x3):")
		board := b.MakeBoard(3, " ")
		b.PrintBoard(board)
		fmt.Println()
		fmt.Println()
	}
	{
		fmt.Println("Vollbesetztes TicTacToe-Board und dessen Diagonalen ausgeben.")
		board := b.MakeNumberedBoard(3)
		b.PrintBoard(board)
		fmt.Printf("Diagonale 1: %v\n", b.GetPrincipalDiagonal1(board))
		fmt.Printf("Diagonale 2: %v\n", b.GetPrincipalDiagonal2(board))
		fmt.Println()
		fmt.Println()
	}
	{
		fmt.Println("Leeres Schachbrett (Größe 8x8):")
		board := b.MakeBoard(8, " ")
		b.PrintBoard(board)
		fmt.Println()
		fmt.Println()
	}
}
