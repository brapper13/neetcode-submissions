# Review notes — is-palindrome (submissions 3, 10, 11, 13)

**Verdict:** canonical two-pointer solution reached in sub-13, and it
survived a 10-case adversarial run locally. First problem of the Two
Pointers block, driven from O(n)-space to O(1)-space through three shapes.
Submissions 4–9 were lost to the private-repo sync outage.

## Progression

- **sub-3** — correct, O(n) space: lowercase the whole string, build a
  cleaned copy with `strings.Builder`, then mirror-check the copy. Two
  redundancies: the `'A' <= b && b <= 'Z'` arm was dead code after
  `ToLower`, and the mirror loop ran the full length, checking every pair
  twice.
- **sub-10/11** — two pointers, skipping junk in place. sub-11 only
  removed the dead uppercase arm. Both carried the bug below.
- **sub-13** — the fix (`for i < j`), per-byte lowercasing at comparison
  time, and both checks extracted into named helpers.

## The bug that mattered: accepted but wrong

sub-11's outer loop was `for i < len(s)/2`. That bound is the midpoint of
the *original* string, but after junk-skipping the pointers don't move
symmetrically — left-heavy junk pushes `i` past the midpoint while real
characters between `i` and `j` are still unchecked.

Failing input: `"@@@@axya"`. The first iteration compares the outer `a`s,
then `i` is already 5 and `5 < 4` fails — `x` and `y` are never compared,
and the function returns `true` for a non-palindrome.

**NeetCode accepted sub-11.** The judge's tests never front-load junk, so
a wrong program got a green tick. After the pointers themselves became the
loop state, the correct condition was the one they already carry:
`for i < j` — "while there is anything left between us".

## What sub-13 does right

- Bounds check before the character test in the skip loops
  (`i < j && !isAlphaNumeric(s[i])`) — guard first, then index.
- `toLowerCase` on the two bytes being compared, not on the string.
  `s[i]` yields a mutable *copy* of a byte, so string immutability is no
  obstacle, and no second string is ever built. That is the whole O(1).
- `b + 'a' - 'A'` instead of a bare `+ 32` — the offset trick with its
  meaning visible.
- Helpers for the two predicates. Function calls cost nothing (takeaway
  12); the duplicated six-clause condition did.

## Remaining polish

- `isAlphaNumeric` ends with `if cond { return true } return false` —
  return the condition directly: `return 'a' <= b && ... `.
- gofmt writes `len(s) - 1`, not `len(s)-1`. One `gofmt` pass locally.

## Idiom notes from the Google breaks

- There is no `isalnum` for bytes in the stdlib. The manual range check
  *is* the idiomatic ASCII answer; `unicode.IsLetter` / `unicode.IsDigit`
  exist for runes.
- `strings.Builder`: `WriteByte` / `WriteString`, then `String()`. The
  alternative — `+=` on a string in a loop — recopies the whole string
  every iteration.
