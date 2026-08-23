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
