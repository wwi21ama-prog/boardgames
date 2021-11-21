package main

import (
	"fmt"

	b "github.com/wwi21ama-prog/boardgames/board"
	h "github.com/wwi21ama-prog/boardgames/helpers"
)

func main() {
	spielfeld := b.MakeBoard(3, " ")

	mainLoop(spielfeld)
}

// Abwechselnd Spieler X und O nach ihren Zügen fragen,
// die Züge machen und auswerten, ob jemand gewonnen hat.
func mainLoop(board [][]string) {
	currentPlayer := "O"
	// Anmerkung: Wir starten mit Spieler O, weil eigentlich spieler X starten soll.
	// Unten in der Schleife wird als erstes der Spieler gewechselt.

	// Solange der aktuelle Spieler nicht gewonnen hat...
	for !checkWinner(board, currentPlayer) {

		// Den Spieler wechseln
		// Dies geschieht hier vor dem Zug, damit nach Ende der Schleife currentPlayer noch
		// der Spieler ist, der zuletzt einen Zug gemacht hat.
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}

		// Den aktuellen Spieler nach seinem Zug fragen und ihn ausführen.
		move(board, currentPlayer)
	}
}

// Fragt den Spieler mit dem Zeichen 'char' nach seinem Zug
// (z.B. Zeile und Spalte) und führt ihn aus.
func move(board [][]string, char string) {
	b.PrintBoard(board)
	fmt.Printf("Spieler %v, du bist dran.\n", char)

	var row, col int

	fmt.Print("Gib deine Zeile ein: ")
	fmt.Scanln(&row)
	if row < 0 || row > 2 {
		fmt.Println("Die Zeile ist nicht gültig. Bitte noch einmal probieren.")
		move(board, char)
		return
	}
	fmt.Print("Gib deine Spalte ein: ")
	fmt.Scanln(&col)
	if col < 0 || col > 2 {
		fmt.Println("Die Spalte ist nicht gültig. Bitte noch einmal probieren.")
		move(board, char)
		return
	}
	if board[row][col] != " " {
		fmt.Println("Das Feld ist schon belegt. Bitte noch einmal probieren.")
		move(board, char)
	}
	board[row][col] = char
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Zeile drei mal 'char' enthält.
func checkRows(board [][]string, char string) bool {
	// Zeile 0 pruefen
	for _, row := range board {
		if h.AllElementsEqualTo(row, char) {
			return true // Early Out
		}
	}
	return false
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Spalte drei mal 'char' enthält.
// Funktioniert erstmal nur für quadratische Spielfelder.
func checkColumns(board [][]string, char string) bool {
	for i, _ := range board {
		if h.AllElementsEqualTo(b.GetColumn(board, i), char) {
			return true // Early Out
		}
	}
	return false
}

// Nimmt das Board als Parameter und liefert true, wenn
// es in einer Diagonale drei mal 'char' enthält.
func checkDiagonals(board [][]string, char string) bool {
	if h.AllElementsEqualTo(b.GetDiagonal1(board), char) {
		return true
	}
	if h.AllElementsEqualTo(b.GetDiagonal2(board), char) {
		return true
	}

	return false
}

// Prüft, ob der Spieler mit dem Zeichen 'char' gewonnen hat.
func checkWinner(board [][]string, char string) bool {
	return checkRows(board, char) || checkColumns(board, char) || checkDiagonals(board, char)
}
