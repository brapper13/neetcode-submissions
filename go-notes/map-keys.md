# Map keys and comparability

## The rule

A map key must support `==`. Comparable: numbers, strings, booleans,
pointers, channels, interfaces, arrays of comparable elements, and structs
whose fields are all comparable. Not comparable: slices, maps, and
functions. One slice field makes a whole struct unusable as a key.

## Why slices can't be keys

A slice is a header — pointer, length, capacity — over a shared, mutable
backing array. For `==` on slices, Go would have to pick one of two
meanings, and both break:

- **Header identity**: two slices are equal only if they point at the same
  memory. Then two slices with identical contents compare unequal, which
  is not what anyone reading `a == b` expects.
- **Content equality**: compare element by element. This reads sensibly
  but corrupts maps. The map hashes the key's contents on insert and files
  it in a bucket. Mutate the backing array through another slice header
  afterwards, and the stored key's contents no longer match its hash — the
  entry is filed under a hash it no longer has, and lookups miss it.

Go refuses to choose. Slices only compare to `nil`, and cannot be keys.

## Why arrays can

An array is a value, not a header. Its size is part of its type, and the
map copies it on insert. Nothing can alias the copy inside the map, so
content comparison is safe and stable. That's the Group Anagrams trick:
`map[[26]int][]string` — the letter counts themselves are the key.

## How `==` on arrays actually works

Element by element, in order, stopping at the first mismatch. Length is
never checked at runtime, because length is part of the type —
`[26]int` and `[25]int` are different types and comparing them is a
compile error.

The compiler is free to do better than a literal loop. For element
types whose equality is bitwise (ints, bytes), it calls the runtime's
memequal, which compares memory in word-sized chunks like memcmp. A
`[26]int` compare is a handful of 8-byte comparisons with an early
exit. This is why comparing count arrays per step (Permutation in
String, `countS1 == local`) is cheap.

The trap: an array compare is O(array size) every time it runs. Fine
at 26 or 128 slots in a loop. Worth a thought at a million.

## What pointers as keys mean

Pointers are comparable, but by identity: two pointers are equal only if
they point at the same object. Two objects with equal contents are two
different keys. Mutating the pointee is safe and does not move the key,
because the key is the address, not the contents. Use a pointer key when
you mean "this specific object", never when you mean "anything that looks
like this".
