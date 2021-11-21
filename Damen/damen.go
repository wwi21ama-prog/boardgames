package damen

import (
	b "github.com/wwi21ama-prog/boardgames/board"
)

// Löst die Aufgabe, n Damen auf einem nxn-Schachbrett zu platzieren, die sich
// nicht gegenseitig schlagen können. Gibt das fertige Spielfeld aus.
// Falls es keine Lösung für die Größe n gibt, wird ein leeres Feld ausgegeben.
func Damen(n int) {
	board := b.MakeBoard(n, " ")
	solve(board, 0)
	b.PrintBoard(board)
}

// Löst die Aufgabe für das gegebene Spielfeld ab Zeile row.
// Liefert true, falls das Spiel gelöst werden konnte.
func solve(board [][]string, row int) bool {
	// Wenn die Zeile größer ist als die Zeilenzahl, dann ist das Spiel schon gelöst.
	if row >= len(board) {
		return true
	}

	// Die aktuelle Zeile durchgehen und ...
	for col := 0; col < len(board); col++ {
		// ... Eine Dame platzieren, sofern das erlaubt ist.
		// Wenn es nicht erlaubt ist, dann passiert in diesem Schleifendurchlauf gar
		// nichts mehr.
		if allowed(board, row, col) {
			board[row][col] = "*"

			// Versuchen, das Spiel ab der nächsten Zeile zu lösen.
			done := solve(board, row+1)

			// Wenn gelöst, dann sind auch wir hier fertig.
			if done {
				return true
			}

			// Wenn nicht gelöst, die gesetzte Dame wieder wegnehmen, damit es im nächsten
			// Schleifendurchlauf in der nächsten Zeile versucht werden kann.
			board[row][col] = " "
		}
	}
	// Wenn die Schleife durchläuft, ohne dass das Spiel gelöst wurde, ist es in der
	// aktuellen Zeile nicht (mehr) lösbar. Deshalb false liefern.
	return false
}

// Prüft, ob es erlaubt ist, in der angegebenen Zeile und Spalte eine Dame zu setzen.
func allowed(board [][]string, row, col int) bool {
	return rowAllowed(board, row) &&
		colAllowed(board, col) &&
		diag1Allowed(board, row, col) &&
		diag2Allowed(board, row, col)
}

// Prüft, ob es erlaubt ist, in der angegebenen Zeile eine Dame zu setzen.
func rowAllowed(board [][]string, row int) bool {
	return b.RowEquals(board, row, " ")
}

// Prüft, ob es erlaubt ist, in der angegebenen Spalte eine Dame zu setzen.
func colAllowed(board [][]string, col int) bool {
	return b.ColumnEquals(board, col, " ")
}

// Prüft, ob es erlaubt ist, in der angegebenen Diagonale von links oben nach rechts unten eine Dame zu setzen.
func diag1Allowed(board [][]string, row, col int) bool {
	return b.Diag1Equals(board, row, col, " ")
}

// Prüft, ob es erlaubt ist, in der angegebenen Diagonale von rechts oben nach links unten eine Dame zu setzen.
func diag2Allowed(board [][]string, row, col int) bool {
	return b.Diag2Equals(board, row, col, " ")
}
