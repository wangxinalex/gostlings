# Sync: the lifecycle of shared memory

Work through 1–14 in order. The protagonists here are shared state and the
standard-library synchronization primitives; message flow, worker pipelines,
and request/reply compositions are left to channels and
`24_concurrency_patterns`.

| Exercise | Focus |
| ---: | --- |
| 1–2 | Mutex-protected counters and read-modify-write on a map |
| 3–5 | `sync.Once`, WaitGroup Add/Done/Wait ordering, and dynamic tasks |
| 6–7 | RWMutex cache, slow computation outside the lock, protected commit |
| 8 | Concurrent config initialization happens exactly once |
| 9–10 | Cond predicate loop, Signal and Close/Broadcast |
| 11–12 | `sync.Map` canonical LoadOrStore, `sync.Pool` borrow/return boundaries |
| 13 | Bounded shared state, waiting for capacity, and shutdown wake-up |
| 14 | Once initialization, active count, Finish, and service Wait |

All sync exercises get selective race coverage from `race.list`; semantic
tests still come before race checks, because "this schedule did not hit a
race" does not mean the implementation is correct. Cond must check the
predicate inside the lock and wait with a `for` loop.

```sh
sh check.sh exercises/16_sync --run-all
sh check.sh solutions/16_sync --run-all --race
```
