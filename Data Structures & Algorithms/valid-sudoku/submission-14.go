func isValidSudoku(board [][]byte) bool {
	// check rows
	var rows, cols [9][9]bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				if rows[i][int(board[i][j] - '0') - 1] {
					return false
				}
				rows[i][int(board[i][j]-'0') - 1] = true
			}
			if board[j][i] != '.' {
				if cols[i][int(board[j][i] - '0') - 1] {
					return false
				}
				cols[i][int(board[j][i]-'0') - 1] = true
			}
		}
	}

	// check sub-box
	var startRow int
	var startColumn int
	for startRow < len(board) {
		var boxDups [9]int
		for i := startRow; i < startRow+3; i++ {
			for j := startColumn; j < startColumn+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				if boxDups[int(board[i][j]-'0')-1] > 0 {
					return false
				}
				boxDups[int(board[i][j]-'0')-1]++
			}
		}
		startColumn += 3
		if startColumn >= len(board) {
			startColumn = 0
			startRow+=3
		}
	}
	return true
}
