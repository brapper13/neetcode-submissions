package main

func isValidSudoku(board [][]byte) bool {
	// your rewrite: single pass, extract the repeated index expression
	// no duplicates in a row
	// no duplicates in a column
	// common number theory - assume a cell has index (i, j), the box number for the call is (i/3)*3 + (j/3)
	// cells are numbers from 1 to 9, we'll store the hints in 3 9x9 grids
	var rowCheck [9][9]bool
	var colCheck [9][9]bool
	var gridCheck [9][9]bool
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == '.' {
				continue
			}
			grid := (row/3)*3 + (col / 3)
			value := board[row][col] - '0' - 1
			if rowCheck[row][value] || colCheck[col][value] || gridCheck[grid][value] {
				return false
			}
			rowCheck[row][value] = true
			colCheck[col][value] = true
			gridCheck[grid][value] = true
		}
	}
	return true
}
