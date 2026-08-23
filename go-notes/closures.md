# Closures and escape analysis

## What a closure is

A closure is a function value bundled with references to the variables it
uses from its enclosing scope. It captures the *variables*, not their
values at creation time. Read or write the variable later through the
closure and you see the live, current variable — the same one the
enclosing function sees.

That is exactly why a deferred closure sees final values while a deferred
plain call sees snapshots (see [defer.md](defer.md)).

## Escape analysis

Local variables normally live on the stack and die with the frame. A
captured variable can outlive the function that declared it — the closure
might be called after the function returns. The compiler's escape analysis
detects this and moves such variables to the heap.

So closures are the common reason a seemingly-local variable allocates.
The variable didn't get bigger — its lifetime got longer than the stack
frame allows. `go build -gcflags='-m'` shows these decisions.

## The loop-variable footnote

Before Go 1.22, a `for` loop had one loop variable reused across
iterations, and every closure created in the loop captured that single
variable — the classic "all my goroutines printed the last value" bug.
Since Go 1.22 each iteration gets a fresh variable, and the bug is gone.
Worth knowing both behaviours: interviewers still ask about the old one.
