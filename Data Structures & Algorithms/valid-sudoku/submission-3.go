func isValidSudoku(board [][]byte) bool {
	// check rows
	for i := 0; i < len(board); i++ {
		var rowDups [9]int
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			rowDups[int(board[i][j]-'0')-1]++
		}
		for _, item := range rowDups {
			if item > 1 {
				return false
			}
		}
	}
	// check column
	for i := 0; i < len(board); i++ {
		var colDups [9]int
		for j := 0; j < len(board[0]); j++ {
			if board[j][i] == '.' {
				continue
			}
			colDups[int(board[j][i]-'0')-1]++
		}
		for _, item := range colDups {
			if item > 1 {
				return false
			}
		}
	}
	
	// check sub-box
	var startRow int
	var startColumn int
	for startRow < len(board) {
		fmt.Println(startRow, startColumn)
		var boxDups [9]int
		for i := startRow; i < startRow+3; i++ {
			for j := startColumn; j < startColumn+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				boxDups[int(board[i][j]-'0')-1]++
			}
		}
		for _, item := range boxDups {
			if item > 1 {
				return false
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
