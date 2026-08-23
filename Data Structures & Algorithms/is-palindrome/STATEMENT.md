# Valid Palindrome

Given a string `s`, return `true` if it reads the same forwards and
backwards once you ignore case and skip every non-alphanumeric character.
A string that is empty after cleaning counts as a palindrome.

**Example**

```
Input:  s = "Was it a car or a cat I saw?"
Output: true        (cleaned: "wasitacaroracatisaw")

Input:  s = "race a car"
Output: false       (cleaned: "raceacar")
```

**Constraints**

- ASCII input; letters, digits, punctuation, spaces.
- Length up to ~1000.

**Target:** O(n) time, O(1) space — two pointers converging from the ends,
skipping non-alphanumeric characters in place, comparing lowercased bytes.
Building a cleaned copy first also passes, but costs O(n) space.
