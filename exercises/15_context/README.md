# Context: from cancellation signals to request cleanup

Work through 1–14 in order; do not skip. Every exercise keeps an explicit
`TODO:` seam; the focused test checks the real blocking point, context
propagation, and cleanup rather than only a literal result.

| Exercise | Focus |
| ---: | --- |
| 1–2 | `WithCancel`/`WithTimeout`, `Done`, capacity-one result, defer cancel |
| 3–5 | Typed value, `errors.Is`, absolute deadline and cancel |
| 6–7 | Parent/child propagation, child isolation, done only after every child exits |
| 8–9 | Cancelable blocked receive/send |
| 10–12 | Worker join, pre-cancellation checks, passing the same context through a helper chain |
| 13 | `WithCancelCause` and `context.Cause` |
| 14 | Request-level composition of deadline, workers, result, cancellation, and final join |

The Hints give each exercise's minimal contract: blocking operations must also
listen on `ctx.Done()`, functions that derive a context must release it, and
completion signals must wait for every worker to return. Raw-channel
fan-in/fan-out belongs to `13_channels`; shared-memory primitives belong to
`16_sync`.

```sh
sh check.sh exercises/15_context --run-all
sh scripts/verify_exercise_starters.sh exercises/15_context
```
