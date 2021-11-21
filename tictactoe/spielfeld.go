package main

import "fmt"

func main() {
	spielfeld := makeBoard(3, " ")

	mainLoop(spielfeld)
}

// Abwechselnd Spieler X und O nach ihren Zügen fragen,
// die Züge machen und auswerten, ob jemand gewonnen hat.
func mainLoop(board [][]string) {
	currentPlayer := "X"

	// Solange der aktuelle Spieler nicht gewonnen hat...
	for !checkWinner(board, currentPlayer) {

		// Den Spieler wechseln

		// Den aktuellen Spieler nach seinem Zug fragen und ihn ausführen.
		move(board, currentPlayer)

	}
}

// Fragt den Spieler mit dem Zeichen 'char' nach seinem Zug
// (z.B. Zeile und Spalte) und führt ihn aus.
func move(board [][]string, char string) {
	fmt.Printf("Spieler %v, du bist dran.\n", char)

	var row, col int

	fmt.Print("Gib deine Zeile ein: ")
	fmt.Scanln(&row)
	fmt.Print("Gib deine Spalte ein: ")
	fmt.Scanln(&col)

	board[row][col] = char
}

func makeBoard(size int, initChar string) [][]string {
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
func printBoard(board [][]string) {
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

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Zeile drei mal 'char' enthält.
func checkRows(board [][]string, char string) bool {
	// Zeile 0 pruefen
	for _, row := range board {
		if checkList(row, char) {
			return true // Early Out
		}
	}
	return false
}

// Prüft eine einzelne Zeile, ob sie drei mal 'char' enthält.
func checkList(row []string, char string) bool {
	for _, element := range row {
		if element != char {
			return false
		}
	}
	return true
}

// Liefert die i-te Spalte des Spielfelds als Liste.
func getColumn(board [][]string, i int) []string {
	var result []string
	for _, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Spalte drei mal 'char' enthält.
// Funktioniert erstmal nur für quadratische Spielfelder.
func checkColumns(board [][]string, char string) bool {
	for i, _ := range board {
		if checkList(getColumn(board, i), char) {
			return true // Early Out
		}
	}
	return false
}

// Liefert die Diagonale von links oben nach rechts unten.
func getDiagonal1(board [][]string) []string {
	var result []string
	for i, _ := range board {
		result = append(result, board[i][i])
	}

	return result
}

// Liefert die Diagonale von rechts oben nach links unten.
func getDiagonal2(board [][]string) []string {
	var result []string
	l := len(board) - 1
	for i, _ := range board {
		result = append(result, board[i][l-i])
	}

	return result
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Diagonale drei mal 'char' enthält.
func checkDiagonals(board [][]string, char string) bool {
	if checkList(getDiagonal1(board), char) {
		return true
	}
	if checkList(getDiagonal2(board), char) {
		return true
	}

	return false
}

// Prüft, ob der Spieler mit dem Zeichen 'char' gewonnen hat.
func checkWinner(board [][]string, char string) bool {
	return checkRows(board, char) || checkColumns(board, char) || checkDiagonals(board, char)
}
