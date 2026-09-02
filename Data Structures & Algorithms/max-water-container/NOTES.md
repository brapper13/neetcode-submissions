# Container With Most Water — review

Canonical greedy two-pointer on the first submission, including the
correct move rule (advance the shorter side).

Polish notes (not resubmitted):

- `(j-i) * min(heights[i], heights[j])` is computed twice. Name it:
  `area := ...`, or fold the comparison into
  `maxArea = max(maxArea, area)`.
- Same `continue`/`else` redundancy as Two Sum II — one shape, not
  both.
- Uses the built-in `min`, which worked — and that's diagnostic. `min`
  is a Go 1.21 builtin, so the toolchain is modern; the `slices`
  package failing means NeetCode pre-imports a fixed package set
  (submissions carry no import lines) and `slices` isn't in it.
  Builtins need no import, packages do.

The proof sketch behind the move rule is worth rehearsing aloud: fix
the shorter line. Any container it forms with a *nearer* opposite wall
has smaller width and height still capped by the same shorter line, so
its area cannot beat the current one. Discarding the shorter line
discards no candidate answers.
