# Review notes — valid-sudoku (submissions 3–4)

**Verdict:** correct, accepted. Three-scan structure with [9]int count
arrays. Sub-4 exists only to remove a shipped fmt.Println — self-caught
this time, which is the ritual starting to work.

## Credits

- Bounded-domain pattern fired unprompted ([9]int, three times) one problem
  after the maps miss. Learned → automatic.
- Digit indexing via board[i][j]-'0' clean throughout.

## Improvements (his own takeaways, applied back)

1. Count-then-scan → check-before-insert (takeaway 2). A duplicate is
   detectable the moment it appears: if count already positive, return
   false. Deletes the scan loops, exits at first crime.
2. The box walk (single loop, wrapping startColumn, manual carriage
   return) needs a trace to trust. Nested loops stepping by 3 need none
   (takeaway 3).
3. Do-over target: single pass over the board updating rows/cols/boxes
   together. Key identity to DERIVE, not memorise: box = (i/3)*3 + j/3.

## Optional flourish

[9]int arrays only answer seen/not-seen. Nine bits in one int does the
same. His own Apple Notes bitwise note covers the needed ops. Trading-firm
interviewers enjoy the bitmask version.

## Submissions 5 and 13 (post-review iterations)

Sub-5: check-before-insert applied to all three checks, scan-after loops
deleted. Sub-13: rows+cols merged into one scan via the transpose read
(board[i][j] and board[j][i] in the same loop), [9][9]bool arrays adopted.
Boxes still a separate odometer scan — the box-identity merge stays a
do-over item.

Costs observed: eight attempts between 5 and 13, almost certainly on the
transpose indexing. Cure prescribed: extract `d := int(board[i][j]-'0')-1`
into a named local per branch. The expression appears six times with two
index orders — every repetition is a chance to swap the wrong pair.

fmt.Println shipped a FIFTH time (column branch, sub-13). Escalated from
note to rule: grep the file for Println before every paste. No exceptions.

## Submission 15 — single pass reached

Boxes merged via box := (i/3)*3 + (j/3). Odometer walk deleted. One scan,
three [9][9]bool structures, exit at first duplicate. Canonical solution.
Full arc: three scans + count-then-check (sub-3) → check-before-insert
(sub-5) → rows+cols transpose merge (sub-13) → single pass (sub-15).

Comprehension check issued: derive the box formula, don't remember it.
Remaining do-over polish: extract d per branch (expression still appears
seven times), which also allows the combined check
`if rows[i][d] || boxes[box][d]`.
