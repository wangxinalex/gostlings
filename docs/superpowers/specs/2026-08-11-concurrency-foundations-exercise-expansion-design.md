# Concurrency Foundations Exercise Expansion Design

## Goal

Expand the goroutine and channel learning path into a sequential, repetition-heavy curriculum: 10 goroutine exercises followed by 50 channel exercises. The learner should be able to complete each chapter from directory 1 upward without skipping ahead to a later concurrency concept.

## Scope and boundaries

This design covers `exercises/12_goroutines` and `exercises/13_channels`.

`12_goroutines` owns the mechanics of starting goroutines, passing values safely into them, and waiting for a known group to finish. Its exercises do not teach channel protocols or shared-memory protection as the main idea.

`13_channels` owns native channel semantics and coordination protocols: blocking, buffering, close ownership, `select`, raw stop channels, done/result separation, fan-out, fan-in, worker pools, request/reply, backpressure, semaphores, and token-channel rate limiting.

The channel chapter will not introduce `context.Context`, mutexes, atomics, timer/ticker construction, error wrapping, or context-aware pipeline design. Those responsibilities belong to the later chapters described in the companion designs. The current channel pipeline exercises 19–21 will move to `24_concurrency_patterns`, where they can be integrated with context and other primitives instead of duplicating a second pipeline curriculum.

## Spiral-learning rules

Every high-frequency channel skill appears in at least three forms:

1. A small isolated exercise that exposes the basic rule.
2. A variation where a different operation can block or where the input can be empty.
3. A composed exercise that adds cancellation, output closure, backpressure, ordering, or an error path.

Hints remain self-contained. They may name the earlier pattern being reused, but they must also identify the sender, receiver, closer, blocking operation, cancellation case, and required exit order for the current exercise.

## Sequential curriculum

### `12_goroutines`: 10 exercises

| Exercise | Focus | Progression |
| --- | --- | --- |
| 1 | Launch one goroutine | Move work behind `go` and observe completion. |
| 2 | Closure capture | Make each goroutine observe the intended loop value. |
| 3 | Explicit goroutine parameters | Pass the loop value as an argument instead of relying on capture. |
| 4 | Basic `WaitGroup` lifecycle | Add before launch, defer `Done`, wait before returning. |
| 5 | Wait for a dynamic number of tasks | Derive the group size from input and handle an empty input. |
| 6 | Per-task completion bookkeeping | Ensure every started goroutine reaches its completion path exactly once. |
| 7 | Batch execution | Run independent batches and wait for each batch before starting the next. |
| 8 | Worker function boundaries | Pass immutable job data into workers and keep the worker body small. |
| 9 | Completion on every branch | Use deferred completion to cover early returns without using sleeps. |
| 10 | Goroutine lifecycle review | Combine parameter passing, dynamic launch, and a final join over a small task set. |

The chapter may use `sync.WaitGroup` as the waiting mechanism because it is the prerequisite for later synchronization work, but it does not ask learners to protect shared mutable state. Output assertions use deterministic collection or a test-safe result path rather than scheduler timing.

### `13_channels`: 50 exercises

| Exercise | Focus | Existing source or role |
| --- | --- | --- |
| 1 | Unbuffered handoff | Existing channels1 |
| 2 | Buffered channel capacity | Existing channels2 |
| 3 | Close plus `range` | Existing channels3 |
| 4 | Comma-ok receive | Existing channels4 |
| 5 | Drain buffered values after close | New semantic reinforcement |
| 6 | Generator ownership | Existing channels5 |
| 7 | Directional channel API | New ownership reinforcement |
| 8 | Basic receive `select` | Existing channels6 |
| 9 | Non-blocking receive | Existing channels8 |
| 10 | Non-blocking send | New send-side `default` variation |
| 11 | Select timeout | Existing channels7 |
| 12 | Select over both receive and send | New relay variation |
| 13 | Done completion notification | Existing channels9 |
| 14 | Buffered one-shot result | New sender-lifecycle reinforcement |
| 15 | Separate result and done channels | New protocol distinction |
| 16 | Cancellable send | Existing channels10 |
| 17 | Cancellable receive and forward | New receive-side cancellation variation |
| 18 | Broadcast cancellation | Existing channels12 |
| 19 | Timeout, cancellation, and join | Existing channels11 |
| 20 | Coordinated shutdown | New stop → wait → done protocol |
| 21 | Basic fan-out | Existing channels13 |
| 22 | Basic fan-in | Existing channels14 |
| 23 | Cancellable fan-in | Existing channels15 |
| 24 | Cancellable fan-out | New worker receive/send variation |
| 25 | Fan-in with empty and already-closed inputs | New close-boundary reinforcement |
| 26 | Basic worker pool | Existing channels16 |
| 27 | Ordered worker-pool results | Existing channels17 |
| 28 | First-error worker-pool cancellation | Existing channels18 |
| 29 | Directional worker-pool API | New public API boundary |
| 30 | Worker-pool backpressure | New concurrent submission variation |
| 31 | First result wins | New cancellation of losing tasks |
| 32 | Request/reply with per-request response channels | New service-style protocol |
| 33 | Multi-producer close coordinator | New single-closer protocol |
| 34 | Worker-pool empty and zero-worker boundaries | New lifecycle reinforcement |
| 35 | Fan-out with cancellation and output closure | Repeated fan-out failure path |
| 36 | Fan-in of result envelopes | Repeated fan-in with explicit result status |
| 37 | Result collection with caller abandonment | Repeated buffered result and cancellation path |
| 38 | Done broadcast to multiple observers | Repeated close-as-broadcast pattern |
| 39 | Control messages beside data messages | Repeated multi-case `select` protocol |
| 40 | Relay with cancellable receive and send | Repeated two-sided leak prevention |
| 41 | Nil channel as a dynamic `select` switch | Existing channels24 |
| 42 | Downstream early exit and producer cancellation | Repeated backpressure/cleanup path |
| 43 | Buffered channel semaphore | Existing channels22 |
| 44 | Cancellable semaphore acquisition | New resource-control variation |
| 45 | Token-channel rate limiter | Existing channels23, refactored to consume external tokens |
| 46 | Cancellable token-channel rate limiter | New rate-limit cleanup variation |
| 47 | Graceful shutdown of a channel service | New integrated raw-channel lifecycle |
| 48 | Bounded, ordered, cancellable worker pool | New integrated composition |
| 49 | Bounded first-error worker pool | New integrated failure path |
| 50 | Raw-channel service capstone | Combines request/reply, workers, backpressure, timeout, cancellation, ordering, and close ownership |

The former channels19–21 pipeline exercises are moved to `24_concurrency_patterns`. The former channels23 rate limiter remains in this chapter only as a token-channel exercise: the learner supplies the event channel; `21_time` owns the construction and cleanup of `time.Ticker`.

## Exercise artifact contract

Every exercise directory contains:

- `main.go` with `Concept`, `Task`, `Expected behavior`, and an actionable `Hint`;
- `main_test.go` with value assertions and lifecycle assertions appropriate to the exercise;
- no import or dependency on the matching solution.

Every solution directory contains a complete runnable `main.go` with the same public function names and behavior as the exercise. Exercise and solution directory numbers must remain in exact parity.

## Strict starter-failure gate

An untouched learner starter must never be accepted as solved. Every new starter `main.go` must contain at least one `TODO:` marker at the exact implementation seam, and `check.sh` must reject an `exercises/...` target while that marker remains. This static gate runs before the package test, so a placeholder that accidentally prints the expected output or returns a legal zero value cannot pass by accident.

The implementation plan must also run a red-phase verifier over every new exercise before any solution is copied: each pristine starter must either fail to compile or fail its focused test, and the failure must identify the intended behavior. A test that passes against the untouched starter is a design failure and must be rewritten before implementation continues.

The normal checker rejects the TODO marker before running the package. The verifier may use an internal `--verify-starter` mode that bypasses only this static rejection, runs the focused test, and requires a non-zero result; that mode is not valid for solutions or for normal learner validation.

The focused tests must exercise the TODO seam directly and include at least one edge or lifecycle assertion. A solution is not accepted because its example `main` prints the expected text alone; the test must also prove closure, ordering, cancellation, bounded progress, or the relevant synchronization boundary.

## Testing and verification

Learner tests use explicit completion channels, closed outputs, and bounded watchdog timeouts. They do not use `runtime.NumGoroutine` as the sole leak assertion and do not use `time.Sleep` as synchronization.

The foundation batch must pass:

- `gofmt -d` over all tracked Go files;
- `go test ./solutions/...`;
- `go vet ./solutions/...`;
- `sh check.sh solutions --run-all`;
- `sh check.sh solutions --run-all --race`;
- focused `go test -race` runs only for goroutine and channel solution packages listed in `race.list`;
- directory parity checks used by CI.

The foundation batch must also verify the starter contract with the dedicated starter verifier and must show that the modified solution passes only after the TODO seam is changed. The verifier is a development-quality gate, not a package-wide test of the intentionally incomplete tree.

## Check performance

The expanded tree must not run the race detector for every package by default. Add a root-level `race.list` manifest whose entries are paths relative to `exercises/`; entries identify exercises with meaningful shared-memory race coverage, such as mutable state protected by a mutex or atomic counter. The checker maps the same relative path to `solutions/` automatically.

Keep the following behavior:

- `sh check.sh solutions --run-all` runs every solution normally;
- `sh check.sh solutions --run-all --race` runs every solution, but adds `-race` only for manifest entries;
- `sh check.sh solutions --run-all --race-all` adds `-race` to every solution for release or diagnostic audits;
- a direct exercise path uses the same manifest decision, so a non-race-sensitive package is not slowed down accidentally.

The manifest is reviewed whenever a new exercise is added. Pure channel lifecycle tests, read-only context propagation, and timer/ticker API tests stay out of the manifest unless they intentionally inspect shared mutable state. The normal test suite remains the correctness gate for deadlocks, channel closure, cancellation, and timer cleanup; race detection is not treated as a substitute for those assertions.

## Documentation changes

Rewrite `exercises/13_channels/README.md` as a 50-exercise map with prerequisite bands, repeated-skill references, ownership diagrams, cancellation ordering, and a clear handoff to `15_context`, `21_time`, and `24_concurrency_patterns`.

Add or update `exercises/12_goroutines/README.md` with the 10-exercise path and an explicit handoff to channel-based completion in the next chapter.

Update the root `README.md` counts and links. The root index must show the expanded counts and the chapter order must still be sequential for learners.
