func isValidSudoku(board [][]byte) bool {
	// check rows
	var rows, cols, boxes [9][9]bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// we're giving each cell a box number from 0 to 9
			// position of the box horizonally is based on the col/3
			// position of the box vertically is based on the row but we're 
			// dividing by 3 then multiplying by 3 so we only get the answer 			 // in multiples of 3.
			box := (i/3)*3 + (j/3)
			if board[i][j] != '.' {
				if rows[i][int(board[i][j] - '0') - 1] {
					return false
				}
				if boxes[box][int(board[i][j] - '0') - 1] {
					return false
				}

				boxes[box][int(board[i][j] - '0') - 1] = true
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
	return true
}
