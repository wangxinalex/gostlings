# Concurrency patterns: integrated concurrency problem set

After finishing the goroutine, channel, context, sync, and time chapters,
practice 1–18 in order. Every exercise here combines at least two learned
primitives and adds one new lifecycle or failure strategy; do not jump back
to this chapter to repeat basic channel-closing exercises.

| Exercise | Integrated ability |
| ---: | --- |
| 1–3 | Generator/reducer, cancellable stage, atomic counter + join |
| 4–6 | Multi-stage cancel, bounded pool, typed result/error pipeline |
| 7–9 | Size/time batching, token rate limit, graceful service shutdown |
| 10–12 | Fan-in coordinator, ordered results, cancellable retry/backoff |
| 13–15 | Context-aware Once, atomic metrics, bounded load shedding |
| 16–18 | Buffered cancellation, deadline shutdown, request/reply capstone |

The shared checking order is: stop receiving or delivering → let blocking
operations exit → join workers → let the single coordinator close the output.
Errors in the exercises use `errors.Is` semantics; results must have
order/count/closure assertions, and you cannot pass with empty slices,
ignored errors, or early returns.

```sh
sh check.sh exercises/24_concurrency_patterns --run-all
sh scripts/verify_exercise_starters.sh exercises/24_concurrency_patterns
```
