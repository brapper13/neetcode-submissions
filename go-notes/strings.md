# Strings: Go's design and the C/Pascal history behind it

## What a Go string is

A Go string is a small header: a pointer to a byte array, and a length.
Two words, passed by value, cheap to copy. The bytes themselves are
immutable — the compiler and runtime rely on that to share them safely.

Because the length lives in the header, `len(s)` is O(1) and the bytes can
contain anything, including zero bytes. Nothing in the data is reserved to
mark the end.

## How C does it, and why

A C string is a bare `char*`. The end is marked in-band: a zero byte
(`NUL`) terminates the string. This was the cheap choice on early-70s
hardware — a string is just an array you can hand around as one pointer,
no second field to carry, and every routine scans until it hits zero.

The costs follow directly. `strlen` is O(n) because the length must be
discovered by scanning. And the data cannot contain a zero byte, because
the terminator is indistinguishable from content.

Plain C arrays don't carry a terminator — you pass the length alongside as
a separate parameter. Strings got special treatment because the culture
wanted to pass them as a single bare pointer, so the "where does it end"
information had to hide inside the data.

## How Pascal did it

Pascal put the length in front: first byte is the count, data follows.
Length is O(1) and any byte can appear in the data. The price was the
fixed-size length field — one byte caps a string at 255 characters.

Go's header design is the Pascal idea with the length moved out of the
data and given a full machine word.

## Why in-band termination is a security problem

When the end of the data is marked inside the data, two parsers can
disagree about where the data ends. The classic exploit: a certificate for
`paypal.com\0.evil.com`. The CA's tooling treated the name as
length-delimited and saw a subdomain of `evil.com`, so it issued the cert.
Browsers compared names as C strings, stopped at the zero byte, and saw
`paypal.com`. Same bytes, two parsers, two different names.

The same principle drove the encode/decode problem: length-prefix framing
(out-of-band) is safe, searching the data for a marker (in-band) is not.

## Bytes and runes

`s[i]` indexes raw bytes. `range s` decodes UTF-8 and yields runes.
`len(s)` counts bytes, not characters. Byte indexing and the `x - 'a'`
arithmetic are correct only under an ASCII constraint — state it when you
use them.

## Building strings without quadratic copies

### Why `+=` in a loop is quadratic

Strings are immutable, so `out += piece` can never append in place.
There is no spare room after `out`'s bytes, and even if there were,
writing into them would mutate a string someone else may share.

So every `+=` does the same three steps. Allocate a fresh array big
enough for both parts. Copy all of `out` into it. Copy `piece` after it.
The old array becomes garbage.

The copy of `piece` is fine — you pay that once per piece no matter
what. The killer is the copy of `out`, because `out` gets longer every
iteration. With n pieces of similar size, iteration 1 recopies 1 piece
worth of bytes, iteration 2 recopies 2, iteration 3 recopies 3. The
total is 1 + 2 + ... + n, which is n(n+1)/2 — O(n²). Joining 10,000
pieces of 10 bytes recopies about 500 MB to build a 100 KB string.

The bytes you copied in iteration k are thrown away in iteration k+1.
Almost all the work is copying data that is about to become garbage.

There is a second, smaller cost: n heap allocations and n dead arrays.
Each allocation takes time, and the dead arrays are work for the
garbage collector. Same lesson as the Group Anagrams slowdown — the
expensive thing in a hot loop is usually allocation, not function calls.

### How Builder avoids it

`strings.Builder` keeps one growable `[]byte` buffer with spare capacity
at the end. `WriteString` usually just copies the new piece into the
spare room — no allocation, nothing recopied.

When the buffer does fill up, it grows by roughly doubling. Doubling
means a piece can only be recopied when the buffer size crosses a power
of two, so each byte moves O(1) times on average instead of once per
remaining iteration. Total work for n pieces: O(n) bytes copied, and
only ~log n allocations. That's what "amortised O(1) append" means —
individual writes occasionally pay for a grow, but the average is
constant.

`String()` at the end hands you the buffer as a string without copying
it. This is safe only because Builder refuses to be copied and never
reuses the buffer afterwards — the immutability promise holds.

Build with `WriteString` and `WriteByte`, call `String()` once at the
end. (`bytes.Buffer` and appending to a `[]byte` work the same way;
Builder is just the string-shaped one.)

`*strings.Builder` implements `io.Writer` (it has a
`Write([]byte) (int, error)` method). That is why `fmt.Fprintf(&sb, ...)`
works: the whole `Fprint` family targets any `io.Writer` — a file, a
socket, an HTTP response, or a builder.

Pick the cheapest form that does the job:

1. `sb.WriteString(s)` / `sb.WriteByte(b)` when the pieces already exist.
   Fastest — no format parsing, no reflection.
2. `fmt.Fprintf(&sb, "%d#", n)` when you genuinely need formatting.
   It writes straight into the builder's buffer.
3. `sb.WriteString(fmt.Sprintf(...))` — never. It pays for `fmt` and
   then copies the result a second time. This exact shape is what
   staticcheck QF1012 ("Use fmt.Fprintf instead of WriteString") flags.

For a number followed by a delimiter, `strconv.Itoa(n)` plus
`WriteByte('#')` beats both `fmt` forms.
