# Review notes — string-encode-and-decode (submissions 4, 7)

**Verdict:** submission 7 is correct and survives both breaking inputs that
killed submission 4. The design is his own invention, not the canonical one.

## The design

Canonical: `4#neet`, decoder scans digits until the marker.
His: `#14neet`. One marker byte, one byte for "how many digits the length
has", the length digits, then the data. The decoder never scans. It reads
fixed-width metadata and jumps. This is a length-of-length header, the same
shape as many binary protocols. Legitimate and arguably cleaner than
canonical.

## Fixes for the do-over

1. The `#` is skipped, never read, never verified. Delete it or verify it.
2. Remove the defensive digit check inside the length loop. The encoder
   wrote exactly `digits` digit bytes. Parse by construction.
3. `int(encoded[j] - '0')` replaces Sprintf + Atoi for single digits.
4. Decode: `encoded[j : j+length]` replaces the byte-by-byte Builder copy.
   Substrings are free. New header, same bytes.
5. Encode: this is where strings.Builder belongs. `output +=` in a loop is
   the O(n²) rebuild. Also strconv.Itoa over Sprintf.
6. fmt.Println shipped in a submission again. Second offence.

## Submission 12 (post-review iteration)

Four of six fixes applied. Marker deleted, defensive check deleted, byte
digit conversion, slices in decode. Format is now `12go14neet`. The length
parse via `Atoi(encoded[j:j+digits])` improves on the suggested per-digit
loop. Still open: Builder in Encode, and fmt.Println shipped a third time
(now a workflow fix: grep for Println pre-submit). Decode is do-over ready.

## Credit

- Receiver shadow from sub-4 fixed (`for _, str := range strs`).
- Empty list needs no sentinel now. It encodes to "" naturally.
- Empty string elements work (`#10`), verified by trace.

## History thread

Sub-4 was a C string (in-band marker, forgeable by data). Sub-7 is a Pascal
string, evolved (length ahead of data, plus a fixed-width length-of-length).
Discussion along the way covered C's null terminator, decay, the (pointer,
length) convention, Go's string header, and the null-prefix certificate
attack as the cost of two parsers disagreeing.
