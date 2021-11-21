package board

import (
	"fmt"

	h "github.com/wwi21ama-prog/boardgames/helpers"
)

func MakeBoard(size int, initChar string) [][]string {
	// Definieren einer 2D-Slice
	var board [][]string

	for i := 0; i < size; i++ {
		var row []string
		for j := 0; j < size; j++ {
			row = append(row, initChar)
		}
		board = append(board, row)
	}
	return board
}

// Das Spielfeld mit Trennlinien ausgeben.
func PrintBoard(board [][]string) {
	for i, row := range board {
		printRow(row)
		if i < len(board)-1 {
			fmt.Println("---+---+---")
		}
	}
}

// Eine Zeile des Spielfelds mit Trennzeichen ausgeben.
func printRow(row []string) {
	for j := 0; j < len(row); j++ {
		fmt.Print(" " + row[j] + " ")
		if j < len(row)-1 {
			fmt.Print("|")
		}
	}
	fmt.Println()
}

// Liefert die i-te Spalte des Spielfelds als Liste.
func GetColumn(board [][]string, i int) []string {
	var result []string
	for _, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Liefert die Diagonale von links oben nach rechts unten.
func GetDiagonal1(board [][]string) []string {
	var result []string
	for i, _ := range board {
		result = append(result, board[i][i])
	}

	return result
}

// Liefert die Diagonale von rechts oben nach links unten.
func GetDiagonal2(board [][]string) []string {
	var result []string
	l := len(board) - 1
	for i, _ := range board {
		result = append(result, board[i][l-i])
	}

	return result
}

// Prüft, ob die i-te Zeile des Boards ausschließlich 's' enthält.
func RowEquals(board [][]string, i int, s string) bool {
	return h.AllElementsEqualTo(board[i], s)
}

// Prüft, ob die i-te Spalte des Boards ausschließlich 's' enthält.
func ColumnEquals(board [][]string, i int, s string) bool {
	return h.AllElementsEqualTo(GetColumn(board, i), s)
}

// Prüft, ob die Hauptdiagonale, die von links oben nach rechts unten läuft,
// ausschließlich 's' enthält.
func PrincipalDiag1Equals(board [][]string, s string) bool {
	return h.AllElementsEqualTo(GetDiagonal1(board), s)
}

// Prüft, ob die Hauptdiagonale, die von rechts oben nach links unten läuft,
// ausschließlich 's' enthält.
func PrincipalDiag2Equals(board [][]string, s string) bool {
	return h.AllElementsEqualTo(GetDiagonal2(board), s)
}
