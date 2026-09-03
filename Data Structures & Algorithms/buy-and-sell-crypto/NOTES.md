# Best Time to Buy and Sell Stock — review

Five minutes, first submission, canonical form, fuzz-verified against
a brute-force reference on 5,000 random arrays.

The history is the interesting part. This was quiz question P3 the day
before, cold: the first answer was a bare "dynamic programming?", and
it took two prompts to reach "min price so far". One day later the
solve took five minutes with the reasoning written in the comment
before the code. That's spaced retrieval doing its job — the quiz miss
became the durable version.

The transferable frame, from the quiz discussion: this is the Two Sum
move with a single int instead of a map. Record the past compactly
(the running minimum), evaluate the present against it (profit if sold
today). One question about the past, one variable carrying its answer.

Polish, none of it worth a resubmit:

- `prices[i] <= lowest` — plain `<` does the same job, since updating
  on equal changes nothing.
- The loop could start at `i = 1`: day 0 can only produce profit 0,
  which `maxProfit`'s zero value already covers.
- Same-day buy/sell falls out as profit 0 because `lowest` updates
  before the profit check. That ordering is load-bearing — flipped, a
  strictly-decreasing array would still work, but only by accident of
  the previous minimum. Worth being able to say why the order is
  right.
- `maxProfit` the variable shadows `maxProfit` the function. Legal,
  and common on judges, but in real code the compiler would let you
  recurse by mistake and the name collision reads badly. `best` is the
  usual choice.
