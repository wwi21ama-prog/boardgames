package board

import (
	"fmt"
	"math"

	h "github.com/wwi21ama-prog/boardgames/helpers"
)

// Liefert ein quadratisches Spielfeld der angegebenen Größe.
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

// Liefert ein quadratisches Spielfeld mit durchnummerierten Feldern.
// Ist zu Testzwecken nützlich und je nach Eingabemethode ggf. auch bei Spielen.
func MakeNumberedBoard(size int) [][]string {
	board := MakeBoard(size, " ")
	counter := 0
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = fmt.Sprintf("%v", counter)
			counter++
		}
	}
	return board
}

// Liefert eine Zeilen-Trennlinie für der Spielfeld mit der Länge length.
func dividerLine(length int) string {
	result := ""
	for i := 0; i < length-1; i++ {
		result += "---+"
	}
	result += "---"
	return result
}

// Das Spielfeld mit Trennlinien ausgeben.
func PrintBoard(board [][]string) {
	for i, row := range board {
		printRow(row)
		if i < len(board)-1 {
			fmt.Println(dividerLine(len(row)))
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

// Liefert die Hauptdiagonale von links oben nach rechts unten.
func GetPrincipalDiagonal1(board [][]string) []string {
	return GetDiagonal1(board, 0, 0)
}

// Liefert die Diagonale, die von links oben nach rechts unten durch die
// angegebene Zeile und Spalte geht.
func GetDiagonal1(board [][]string, row, col int) []string {
	var result []string

	// Minimum von row und col bestimmen
	m := int(math.Min(float64(row), float64(col)))

	// Wir berechnen nun mit r und c Versionen von row und col, die um das Minimum
	// reduziert sind. Eines der beiden ist Null, d.h. wir laufen soweit wie möglich
	// nach links oben. Die Position (r,c) sind die Zeile und Spalte des Elements
	// ganz links oben in der Diagonale.
	// Von hier aus läuft die Schleife nun maximal weit nach rechts unten.

	for r, c := row-m, col-m; r < len(board) && c < len(board[0]); r, c = r+1, c+1 {
		result = append(result, board[r][c])
	}

	return result
}

// Liefert die Hauptdiagonale von rechts oben nach links unten.
func GetPrincipalDiagonal2(board [][]string) []string {
	return GetDiagonal2(board, 0, len(board[0])-1)
}

// Liefert die Diagonale, die von rechts oben nach links unten durch die
// angegebene Zeile und Spalte geht.
func GetDiagonal2(board [][]string, row, col int) []string {
	var result []string

	// Ähnlicher Ansatz wie bei GetDiagonal1(), aber wir benutzen dieses mal nicht das
	// Minimum, sondern wir setzen für den Start bedingungslos die Zeile auf 0.
	// Die Schleife läuft durch alle Zeilen. Dabei kann es sein, dass die Spalte
	// zu klein (negativ) oder zu groß (größer als die Länge) ist. Dies wird in der
	// Schleife geprüft.

	for r, c := 0, col+row; r < len(board); r, c = r+1, c-1 {
		if c >= 0 && c < len(board[r]) {
			result = append(result, board[r][c])
		}
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
	return h.AllElementsEqualTo(GetPrincipalDiagonal1(board), s)
}

// Prüft, ob die Hauptdiagonale, die von rechts oben nach links unten läuft,
// ausschließlich 's' enthält.
func PrincipalDiag2Equals(board [][]string, s string) bool {
	return h.AllElementsEqualTo(GetPrincipalDiagonal2(board), s)
}
