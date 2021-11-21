package damen

import (
	b "github.com/wwi21ama-prog/boardgames/board"
)

// Löst die Aufgabe, n Damen auf einem nxn-Schachbrett zu platzieren, die sich
// nicht gegenseitig schlagen können. Gibt das fertige Spielfeld aus.
// Falls es keine Lösung für die Größe n gibt, wird ein leeres Feld ausgegeben.
func Damen(n int) {
	board := b.MakeBoard(n, " ")
	b.PrintBoard(board)
}
