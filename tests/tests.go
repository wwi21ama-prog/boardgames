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
		fmt.Println("Vollbesetztes TicTacToe-Board und einige von dessen Diagonalen ausgeben.")
		board := b.MakeNumberedBoard(3)
		b.PrintBoard(board)
		fmt.Printf("Hauptdiagonale 1: %v\n", b.GetPrincipalDiagonal1(board))
		fmt.Printf("Hauptdiagonale 2: %v\n", b.GetPrincipalDiagonal2(board))
		fmt.Printf("Diagonale nach rechts unten bei (1,0): %v\n", b.GetDiagonal1(board, 1, 0))
		fmt.Printf("Diagonale nach rechts unten bei (2,0): %v\n", b.GetDiagonal1(board, 2, 0))
		fmt.Printf("Diagonale nach links unten bei (0,1): %v\n", b.GetDiagonal2(board, 0, 1))
		fmt.Printf("Diagonale nach links unten bei (0,2): %v\n", b.GetDiagonal2(board, 0, 2))
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
