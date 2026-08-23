# Defer: evaluation time, return values, implementation

## The two timing rules

`defer f(x)` evaluates `f` and its arguments immediately, at the defer
statement. Only the call itself is postponed to function return. So the
deferred call sees `x` as it was when the defer line ran.

A deferred *closure* is different: it captures variables, not values, so
when it finally runs it sees their final values.

```go
i := 1
defer fmt.Println(i)          // prints 1 — argument snapshotted now
defer func() { fmt.Println(i) }() // prints whatever i is at return
i = 2
```

Multiple defers run last-in, first-out.

## Mutating the return value

A function's return value is copied into an anonymous result slot before
the defers run, and the caller reads that slot. A deferred closure can
only change what the caller sees if it can reach the slot.

- **Named result** (`func f() (result int)`): the name *is* the slot. A
  deferred closure assigning to `result` changes the returned value. This
  is the standard way to turn a recovered panic into an error.
- **Unnamed result**: `return x` copies `x` into the slot, and the defer
  has no name for the slot. Mutating `x` afterwards changes nothing.
- **Pointer result**: the slot holds an address. A defer can't change
  which address is returned, but mutating the pointee changes what the
  caller finds there, because both refer to the same object.

## How it's implemented

The compiler picks the cheapest mechanism that works:

- **Open-coded defers** (Go 1.14+): when defers are known at compile time
  and not in loops, the calls are inlined at every return point, guarded
  by a bitmask of which defers activated. Near-zero cost.
- **Stack-allocated records**: otherwise, a `_defer` record on the stack,
  pushed onto a per-goroutine linked list. Return walks the list — that
  linked list is where LIFO comes from.
- **Heap-allocated records**: a defer inside a loop runs an unknown number
  of times, so its records go on the heap. This is why "defer in a loop"
  is both a cost smell and a resource-leak smell (nothing closes until the
  function exits).

`recover` only does anything when called directly inside a deferred
function during a panic — that's the hook that lets the named-result
pattern convert a panic into an ordinary return.
