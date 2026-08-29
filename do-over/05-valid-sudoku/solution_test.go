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
	}
	for _, c := range cases {
		if got := isValidSudoku(c.board); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}
