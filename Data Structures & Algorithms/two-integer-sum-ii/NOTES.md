# Two Sum II — review

Canonical converging two-pointer on the first submission. The palindrome
pattern transferred cleanly: `i`/`j` from the ends, `for i < j`, the
pointers are the loop state.

Polish notes (not resubmitted):

- `numbers[i] + numbers[j]` is computed three times per iteration. Name
  it once: `sum := numbers[i] + numbers[j]`.
- `if ... { j--; continue } else { i++ }` — the `continue` and the
  `else` do the same job. Pick one shape; `if/else if/else` reads
  cleanest here.

Worth saying in an interview: this problem is *why* the O(1)-space
constraint exists — it forbids the Two Sum hashmap and forces you to
spend the sortedness instead.
