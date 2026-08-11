# Integrated Concurrency Patterns Exercise Expansion Design

## Goal

Expand `24_concurrency_patterns` to 18 exercises that combine the previously taught goroutine, channel, context, sync, time, atomic, and error primitives in realistic service-style workflows. These exercises are the integration layer, not a second introduction to channel or context syntax.

## Scope and boundaries

Every exercise in this chapter must combine at least two previously taught primitives and must expose a lifecycle, failure, or resource-management decision that a single-primitive chapter cannot cover alone.

The chapter owns context-aware pipelines, bounded and cancellable worker pools, error propagation, batch flushing, graceful shutdown, metrics, and end-to-end resource cleanup. It may use `sync.WaitGroup`, `sync/atomic`, timers, tickers, raw channels, `context`, and standard errors as implementation tools.

It will not repeat the isolated rules for closing a channel, using `Mutex`, creating a `Ticker`, or calling `WithCancel`; each exercise will link those prerequisites in its hint and test the combination instead.

## Spiral-learning rules

The integrated sequence repeats the most important production failure modes:

- cancellation is tested while receiving, sending, waiting for work, and waiting for results;
- output closure is tested after normal completion, cancellation, and first error;
- bounded concurrency is tested with both a semaphore and a worker queue;
- error propagation is tested alongside cancellation and cleanup;
- time-based flush and rate control are tested with injected or controllable clocks where possible;
- metrics are checked after all workers have joined, not through a race-prone read during execution.

## Sequential curriculum

| Exercise | Focus | Required combination |
| --- | --- | --- |
| 1 | Channel-owned generator and reducer | Directional channels + close + aggregation. |
| 2 | Context-aware pipeline stage | Context cancellation + channel send/receive + output closure. |
| 3 | Atomic worker counter | `atomic.Int64` + goroutine launch + `WaitGroup`. |
| 4 | Multi-stage context pipeline | Multiple channel stages + propagated context + cancellation on both sides. |
| 5 | Bounded context worker pool | Context + semaphore channel + worker results + join. |
| 6 | Error-cancellable pipeline | Error channel + first error + context cancellation + output close. |
| 7 | Timed batcher | Input channel + timer flush + maximum batch size + final flush. |
| 8 | Context-aware rate-limited pool | Token channel or ticker + context + bounded workers + clean stop. |
| 9 | Graceful service shutdown | Stop accepting work + drain or cancel active work + wait + close results. |
| 10 | Fan-in with typed failures | Multiple producers + result envelopes + one closer + error classification. |
| 11 | Ordered concurrent transform with deadline | Indexed results + context deadline + partial failure policy. |
| 12 | Retryable task runner | Context + timer backoff + error classification + final cancellation. |
| 13 | Lazy shared configuration | `sync.Once` + context-aware initialization + stable result publication. |
| 14 | Atomic metrics for a worker service | Atomic counters + worker pool + completion barrier + snapshot after join. |
| 15 | Backpressure and load shedding | Bounded queue + non-blocking submission + explicit rejected-work result. |
| 16 | Multi-stage buffered pipeline | Buffer sizing + cancellation + error propagation + stage ownership. |
| 17 | Shutdown with in-flight deadlines | Context deadline + timer cleanup + worker join + result closure. |
| 18 | Concurrency capstone | Context, channels, bounded workers, ordered results, error cancellation, metrics, timeout, and graceful shutdown. |

The former raw channel pipeline exercises from `13_channels/channels19–21` are adapted into exercises 1–4 here. Their tests and solutions are rewritten to make the integration layer explicit rather than preserving a duplicate raw-channel lesson.

## Interfaces and scenario rules

Functions should expose narrow, testable boundaries such as `run(ctx, jobs)`, `stage(ctx, in)`, or `serve(ctx, requests)`. A caller-provided context controls cancellation; a result channel is closed only by the coordinator that owns all senders. Errors are returned or sent as structured values; the chapter does not rely on comparing unstable error strings unless the exercise explicitly teaches a user-facing message.

Time-dependent exercises accept durations or injected event sources where that makes tests deterministic. Tests may use a bounded timeout as a deadlock guard, but they must also assert the relevant result, error, counter, or closed output.

Every new starter contains an explicit `TODO:` implementation seam, and the checker rejects any learner target that still contains that marker. Before implementation, a red-phase verifier must run every pristine integrated starter and require it to fail for the intended lifecycle, error, cancellation, or bounded-concurrency behavior. This prevents a permissive test or a legal zero value from making an untouched capstone appear complete.

Integrated tests must include at least one failure-path assertion whenever an exercise starts background work: cancellation must be observed, the first error must be propagated, all owned outputs must close, and metrics must be read only after the worker group has joined. A passing example program without these assertions is insufficient.

## Documentation and verification

Create `exercises/24_concurrency_patterns/README.md` with a prerequisite map from chapters 12, 13, 15, 16, and 21. Each row must state which earlier primitive is being combined and which cleanup obligation is new.

Update the root README count and description for this chapter. Keep the chapter after `21_time` so learners encounter all primitive APIs before the integrated scenarios.

Verify all solutions with `go test`, `go vet`, and `sh check.sh solutions --run-all`. Use the shared `race.list` manifest for `sh check.sh solutions --run-all --race`, and reserve `--race-all` for the full audit. Integrated exercises with shared metrics, mutable registries, or atomic state are manifest entries; pure channel/timer behavior is validated by its ordinary focused tests. Focused tests must include cancellation and failure paths for every exercise that starts a goroutine.
