# Valid Sudoku

Given a 9×9 Sudoku board, decide whether the filled cells are valid: no
digit repeats in any row, any column, or any of the nine 3×3 boxes. Empty
cells are `'.'` and are ignored. The board does not need to be solvable —
only the filled cells are judged.

**Example**

```
A board with two '5's in the same column → false
A board with no repeats in any row/column/box → true
```

**Constraints**

- Always 9×9. Cells are `'1'`–`'9'` or `'.'`.

**Target:** one pass with `[9][9]bool` seen-tables for rows, columns and
boxes. Box identity: `box := (i/3)*3 + j/3`.
