package main

import "testing"

func b(rows ...string) [][]byte {
	out := make([][]byte, len(rows))
	for i, r := range rows {
		out[i] = []byte(r)
	}
	return out
}

func TestIsValidSudoku(t *testing.T) {
	empty := b(".........", ".........", ".........", ".........", ".........",
		".........", ".........", ".........", ".........")
	valid := b("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
		"7...2...6", ".6....28.", "...419..5", "....8..79")
	rowDup := b("11.......", ".........", ".........", ".........", ".........",
		".........", ".........", ".........", ".........")
	colDup := b("1........", ".........", ".........", ".........", ".........",
		".........", ".........", ".........", "1........")
	boxDup := b("1........", ".1.......", ".........", ".........", ".........",
		".........", ".........", ".........", ".........")
	sparse := b("1........", "...1.....", ".........", ".........", ".........",
		".........", ".........", ".........", ".........")
	// row dup where the two cells sit in DIFFERENT boxes and columns —
	// catches code whose row check is accidentally a box check
	rowDupFar := b("1.......1", ".........", ".........", ".........", ".........",
		".........", ".........", ".........", ".........")
	// box dup at opposite corners of box 0 — different row AND column
	boxCorners := b("..1......", ".........", "1........", ".........", ".........",
		".........", ".........", ".........", ".........")
	// dup of '9' in the last box — catches box-formula off-by-ones
	lastBox := b(".........", ".........", ".........", ".........", ".........",
		".........", "......9..", ".........", "........9")
	// '5' at (0,3) and its transpose (3,0): different row, column, and box,
	// so the board is VALID — code that mixes board[r][c] with board[c][r]
	// tends to reject it
	transposeOK := b("...5.....", ".........", ".........", "5........", ".........",
		".........", ".........", ".........", ".........")
	// '7' at (3,5) and (5,3): transposes of each other AND both in box 4 —
	// invalid, but only the box check can see it
	transposeBox := b(".........", ".........", ".........", ".....7...", ".........",
		"...7.....", ".........", ".........", ".........")
	// the full valid board with '5' added at (8,0): row 8 and box 6 stay
	// clean, but column 0 already has a '5' at the top
	buriedColDup := b("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
		"7...2...6", ".6....28.", "...419..5", "5...8..79")

	cases := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{"empty", empty, true},
		{"valid full", valid, true},
		{"row dup", rowDup, false},
		{"col dup", colDup, false},
		{"box dup", boxDup, false},
		{"sparse valid", sparse, true},
		{"row dup across boxes", rowDupFar, false},
		{"box dup at corners", boxCorners, false},
		{"dup in last box", lastBox, false},
		{"transposed pair is valid", transposeOK, true},
		{"transposed pair in same box", transposeBox, false},
		{"col dup buried in full board", buriedColDup, false},
	}
	for _, c := range cases {
		if got := isValidSudoku(c.board); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}
