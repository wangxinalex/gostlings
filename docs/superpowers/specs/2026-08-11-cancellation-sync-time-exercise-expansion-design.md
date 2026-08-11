# Cancellation, Synchronization, and Time Exercise Expansion Design

## Goal

Expand `15_context`, `16_sync`, and `21_time` into focused, repeated practice tracks that teach the standard-library primitive owned by each chapter. Each track starts with one concept, repeats it with a different failure or lifecycle boundary, and ends with a small composed exercise without taking over the integrated patterns chapter.

## Ownership boundaries

### `15_context`

This chapter teaches `context.Context` as a request-scoped cancellation, deadline, and value carrier. It uses channels only as the observation mechanism for `ctx.Done()`; it does not re-teach raw channel ownership, fan-in, or worker-pool closure.

### `16_sync`

This chapter teaches shared-memory synchronization with the `sync` package. Its main assertions concern protected state, lock choice, one-time initialization, condition waiting, or reusable synchronization objects. Channels may coordinate tests, but they are not the lesson's primary mechanism.

### `21_time`

This chapter teaches the `time` API: durations, timers, ticker creation, stopping, resetting, deadlines, and periodic work. It may read timer/ticker channels because that is the API surface, but it does not teach general channel protocols or worker-pool design.

`24_concurrency_patterns` consumes these primitives in integrated scenarios. It owns context-aware pipelines and cross-package concurrency design; the three chapters here remain focused on one primitive at a time.

## Spiral-learning rules

Each important standard-library behavior appears in a minimum of three forms:

1. Direct API use with the happy path.
2. A boundary case such as empty input, child cancellation, timeout, or repeated initialization.
3. A composed use where cleanup or propagation must be correct on every branch.

Hints name the exact API operation, the state transition it causes, and the cleanup action that must follow. They do not merely say “use context” or “add a mutex.”

## Sequential curriculum

### `15_context`: 14 exercises

| Exercise | Focus | Progression |
| --- | --- | --- |
| 1 | `WithCancel` basics | Existing cancellation signal exercise. |
| 2 | `WithTimeout` | Existing deadline-like cancellation exercise with a buffered result. |
| 3 | Typed context values | Existing request-scoped value exercise. |
| 4 | Inspect `ctx.Err` | Distinguish `Canceled` from `DeadlineExceeded` after `Done` closes. |
| 5 | `WithDeadline` | Use an absolute deadline and always call the returned cancel function. |
| 6 | Parent-to-child propagation | Cancel a parent and observe multiple descendants stopping. |
| 7 | Child-only cancellation | Prove that canceling a child does not cancel its parent or sibling. |
| 8 | Context-aware receive | Stop waiting for input when `ctx.Done()` becomes ready. |
| 9 | Context-aware send | Stop a blocked result send when the caller cancels. |
| 10 | Context-aware worker group | Start several workers, propagate cancellation, and wait for all to return. |
| 11 | Cancellation before work starts | Check `ctx.Err` or `ctx.Done` before expensive work begins. |
| 12 | Deadline through a call chain | Pass the same context through helper functions without replacing it. |
| 13 | `WithCancelCause` | Preserve a structured cancellation reason and retrieve it with `context.Cause`. |
| 14 | Context lifecycle review | Combine deadline, propagation, cooperative stop, result delivery, and cleanup. |

The value exercise remains focused on typed keys and lookup behavior. It does not become a general data-passing exercise. The cancellation exercises use deterministic gates and result channels rather than fixed sleeps.

### `16_sync`: 14 exercises

| Exercise | Focus | Progression |
| --- | --- | --- |
| 1 | Mutex-protected counter | Existing basic mutual exclusion exercise. |
| 2 | Mutex-protected map | Existing concurrent map write exercise. |
| 3 | `sync.Once` | Existing one-time initialization exercise. |
| 4 | `WaitGroup` lifecycle | Add, defer `Done`, and wait for an empty or non-empty group. |
| 5 | Dynamic task joining | Add all known tasks before waiting and avoid calling `Add` concurrently with `Wait`. |
| 6 | `RWMutex` read-heavy cache | Allow concurrent readers and exclusive writers. |
| 7 | Minimal critical sections | Copy protected state under lock, do slow work after unlock, then commit safely. |
| 8 | Lazy initialization with `Once` | Repeat one-time setup with a returned value and multiple callers. |
| 9 | `sync.Cond` wait/signal | Wait in a predicate loop and signal after state changes. |
| 10 | `sync.Cond` broadcast | Wake a group of waiters after a shared state transition. |
| 11 | `sync.Map` registry | Use `Load`, `Store`, and `LoadOrStore` for a concurrent registry. |
| 12 | `sync.Pool` reuse | Get and put temporary objects without treating the pool as durable storage. |
| 13 | Mutex-protected bounded state | Combine a counter, a limit, and condition signaling without racing. |
| 14 | Synchronization review | Combine initialization, protected state, waiting, and shutdown in one small service. |

The chapter does not teach atomic counters as its main pattern; atomic metrics and lock-free counters remain in `24_concurrency_patterns`. Tests use race detection and deterministic state assertions.

### `21_time`: 8 exercises

| Exercise | Focus | Progression |
| --- | --- | --- |
| 1 | Layout formatting | Existing reference-time layout exercise. |
| 2 | Basic ticker loop | Existing periodic tick exercise with explicit stop. |
| 3 | Duration arithmetic | Convert and compare durations without unit mistakes. |
| 4 | One-shot timer | Read a timer channel and stop the timer on early exit. |
| 5 | Timer in a `select` | Distinguish result, cancellation, and timeout branches. |
| 6 | Safe timer reset | Stop, drain when necessary, reset, and reuse a timer. |
| 7 | Ticker lifecycle | Stop a ticker on every return path and close a completion signal. |
| 8 | Periodic task boundary | Combine ticker events, work duration, and a final deadline without leaking the ticker. |

The time chapter owns construction and cleanup of timers and tickers. The channel chapter's rate limiter receives external tokens so learners can focus on channel flow there.

## Exercise artifact and test contract

Each learner directory has a starter file and focused tests; each solution directory has a runnable reference implementation. Tests must verify the semantic result and the relevant cleanup state. Watchdog timeouts are allowed as deadlock protection, but a test must also observe a channel, state transition, or returned error that proves the intended behavior.

Every new starter contains a `TODO:` marker at its implementation seam. The exercise checker must reject an `exercises/...` target before running tests while that marker remains, preventing a no-op starter with a valid default result from passing. Before implementation, a red-phase verifier runs each pristine starter and requires a compile failure or focused-test failure that corresponds to the stated task. If the pristine starter passes, the test or starter is not sufficiently strict and must be corrected.

Tests must cover the API's failure or cleanup contract, not only the happy-path return value: canceled contexts must be observable, locks must protect the asserted state, timers and tickers must stop on all exits, and empty or repeated calls must exercise the intended boundary. The later solution must pass after the TODO seam is changed; leaving the marker in place is not considered a valid solution.

The expanded primitive chapters must pass the existing solution tests, race tests, vet checks, format checks, and directory-parity checks. Chapter guides must list prerequisite bands and link forward to the integrated patterns chapter without duplicating its capstones.

Race execution follows the repository-level `race.list` manifest. The ordinary checker runs every primitive solution normally; `--race` adds the detector only to exercises whose tests exercise shared mutable state, especially the expanded `16_sync` set. Context propagation and time API tests normally rely on deterministic cancellation and cleanup assertions instead of paying the race-detector cost. `--race-all` remains available for a full audit.
