package board

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
