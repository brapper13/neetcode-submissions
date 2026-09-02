# Encode and Decode Strings

Design two functions: `Encode` packs a list of strings into one single
string, and `Decode` recovers the exact original list. Any character can
appear inside the strings, so a plain delimiter can be forged by the data.

**Example**

```
Input:  ["go", "neet"]
Encode: "12go14neet"     (this repo's length-of-length format)
Decode: ["go", "neet"]
```

**Constraints**

- Up to 100 strings, each up to 200 chars, ASCII.
- The list can be empty, and strings can be empty.

**Target:** O(total length) both ways. The design lesson: put the length
in front (out-of-band framing) instead of searching the data for a marker
(in-band signalling).

### Result
Complete. Took 27:20:75. had some snags but yeah.
The go compiler is telling me that doing string += string in a loop is inefficeint. What other way exists?
