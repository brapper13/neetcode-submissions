# Go notes

Deep dives written up from coaching discussions, one topic per file. These
are the "why does Go work this way" companions to the problem-driven rules
in [TAKEAWAYS.md](../TAKEAWAYS.md). New topics get a file here as they come
up.

- [strings.md](strings.md) — what a Go string is, and the C/Pascal history
  that explains why.
- [map-keys.md](map-keys.md) — what can be a map key, why slices can't,
  and what pointers as keys really mean.
- [defer.md](defer.md) — when arguments are evaluated, how defer can
  mutate a return value, and how it's implemented.
- [closures.md](closures.md) — what a closure captures, and why that
  sometimes means a heap allocation.
- [map-internals.md](map-internals.md) — buckets, load factor, growth,
  and why pre-sizing matters.
