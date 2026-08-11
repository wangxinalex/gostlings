# Integrated Concurrency Patterns Exercise Expansion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task with review checkpoints.

**Goal:** Expand `24_concurrency_patterns` to 18 sequential integration exercises that combine the earlier goroutine, channel, context, sync, time, atomic, and error skills without duplicating their introductory lessons.

**Architecture:** Rename the current `pipeline1`, `pipeline2`, and `atomic1` packages to `concurrency1`, `concurrency2`, and `concurrency3`, then add `concurrency4..18`. Each package exposes a narrow function, has focused normal and failure-path tests in the learner tree, and has a complete runnable solution. Pipelines and worker services use context-aware cancellation, one coordinator-owned output close, and explicit worker joins.

**Tech Stack:** Go 1.26 standard library; `context`, `errors`, `fmt`, `sort`, `sync`, `sync/atomic`, `time`, and `testing`.

## Global Constraints

- Modify only `/Users/wangxinalex/SelfStudy/Rust/gostlings`; do not touch `working-progress`.
- Use `apply_patch` for edits and keep exercise/solution parity.
- Every learner starter has a `TODO:` seam and is rejected by `check.sh` until solved.
- Run a red-phase test for every pristine starter before writing its solution; default returns, nil channels, and empty slices must not satisfy the focused tests.
- Each integrated exercise combines at least two earlier primitives and introduces one new lifecycle, failure, or resource-management decision.
- Every goroutine-starting exercise has a normal path and a cleanup/failure path in its tests.
- Use deterministic gates or injected event sources instead of `time.Sleep` as synchronization.
- Add only shared-state exercises with meaningful race coverage to `race.list`; keep `--race-all` for explicit full audits.
- Do not run `go test ./...` over learner packages.

## File map

- Move or replace: `exercises/24_concurrency_patterns/pipeline1` → `concurrency1`.
- Move or replace: `exercises/24_concurrency_patterns/pipeline2` → `concurrency2`.
- Move or replace: `exercises/24_concurrency_patterns/atomic1` → `concurrency3`.
- Create: `exercises/24_concurrency_patterns/concurrency4..18/{main.go,main_test.go}`.
- Mirror all 18 packages under `solutions/24_concurrency_patterns/`.
- Create: `exercises/24_concurrency_patterns/README.md`.
- Modify: `README.md` and `race.list`.

---

### Task 1: Move existing integration exercises and make their boundaries explicit

**Files:**
- Move: `exercises/24_concurrency_patterns/pipeline1` to `concurrency1`
- Move: `exercises/24_concurrency_patterns/pipeline2` to `concurrency2`
- Move: `exercises/24_concurrency_patterns/atomic1` to `concurrency3`
- Mirror moves under `solutions/24_concurrency_patterns/`
- Remove the former raw pipeline implementations from `exercises/13_channels/channels19..21` when the foundation plan replaces those numbered channel starters; carry their concepts only through the integrated `concurrency1..4` implementations.

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 1 | `generate(values ...int) <-chan int`, `sum(in <-chan int) int` | Producer closes output and reducer returns after close, including empty input. |
| 2 | `square(ctx context.Context, in <-chan int) <-chan int` | Context cancellation interrupts a blocked send and output closes on every exit. |
| 3 | `incrementConcurrently(workers, increments int) int64` | Atomic count equals the expected total after all workers join and race test is clean. |

- [ ] **Step 1: Move directories without changing behavior.**

Use `git mv` for both exercise and solution directories, then update package comments and focused test command paths from the old names to `concurrency1..3`.

- [ ] **Step 2: Add the strict TODO seam to each learner starter.**

Ensure each moved learner file contains a single clear implementation seam. The test must fail on the starter even if the function returns a legal zero value; for example, `sum` must receive a closed output containing values, and `square` must have a blocked-send cancellation test.

- [ ] **Step 3: Run the red phase.**

Run `sh scripts/verify_exercise_starters.sh exercises/24_concurrency_patterns` and confirm the three pristine moved starters fail or are rejected.

- [ ] **Step 4: Run the moved solutions.**

Run `go test -count=1 ./solutions/24_concurrency_patterns/concurrency1`, `concurrency2`, and `concurrency3`; run `go test -race` for `concurrency3`; run `go vet ./solutions/24_concurrency_patterns/...`.

- [ ] **Step 5: Update race metadata and commit the move.**

Add `24_concurrency_patterns/concurrency3` to `race.list` and commit:

```sh
git add exercises/24_concurrency_patterns/concurrency1 exercises/24_concurrency_patterns/concurrency2 exercises/24_concurrency_patterns/concurrency3 solutions/24_concurrency_patterns/concurrency1 solutions/24_concurrency_patterns/concurrency2 solutions/24_concurrency_patterns/concurrency3 race.list
git commit -m "refactor: sequence integrated concurrency exercises"
```

---

### Task 2: Add integrated pipeline and worker exercises 4–6

**Files:**
- Create: `exercises/24_concurrency_patterns/concurrency4..6/{main.go,main_test.go}`
- Create: `solutions/24_concurrency_patterns/concurrency4..6/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 4 | `pipeline(ctx context.Context, in <-chan int) <-chan int` | Multiple stages propagate the same context, stop receiving/sending on cancel, and close final output. |
| 5 | `runPool(ctx context.Context, workers, limit int, jobs []int) ([]int, error)` | Context-aware workers respect a semaphore limit, return all normal results, and join after cancellation. |
| 6 | `type job struct{ value int; fail bool }`, `type result struct{ value int; err error }`; `runPipeline(ctx context.Context, in <-chan job) (<-chan result, <-chan error)` | First error cancels upstream stages, one coordinator closes result, and error is observable. |

- [ ] **Step 1: Write tests for normal completion.**

Exercise 4 must transform three values through at least two stages and assert final closure. Exercise 5 must observe maximum active work with a test-side mutex and assert it never exceeds `limit`. Exercise 6 must assert all successful values before the failure policy and a non-nil error.

- [ ] **Step 2: Write failure-path tests and run red.**

Cancel while a stage has no downstream receiver; provide a failing job; use a 500ms watchdog only to diagnose a leak. The untouched starters must fail or be rejected before solutions are written.

- [ ] **Step 3: Implement context-aware stages and pool.**

Select on `ctx.Done()` during every input receive, semaphore acquisition, and output send. Use a `WaitGroup` coordinator to close outputs only after all senders exit. Return the first error through a capacity-one channel or a coordinator-owned field.

- [ ] **Step 4: Run normal, selective race, and vet checks.**

Run the three solution packages twice normally, run `go test -race` for `concurrency5` if its test observes shared active count, and run `go vet`.

- [ ] **Step 5: Commit the first integration batch.**

```sh
git add exercises/24_concurrency_patterns/concurrency4 exercises/24_concurrency_patterns/concurrency5 exercises/24_concurrency_patterns/concurrency6 solutions/24_concurrency_patterns/concurrency4 solutions/24_concurrency_patterns/concurrency5 solutions/24_concurrency_patterns/concurrency6 race.list
git commit -m "feat: add context-aware pipeline patterns"
```

---

### Task 3: Add timed batching and rate-limited workers 7–9

**Files:**
- Create: `exercises/24_concurrency_patterns/concurrency7..9/{main.go,main_test.go}`
- Create: `solutions/24_concurrency_patterns/concurrency7..9/main.go`
- Modify: `race.list` by adding `24_concurrency_patterns/concurrency5` if its bounded active-count assertion remains shared-state based; the final plan uses an injected counter with a mutex, so this entry is required.

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 7 | `batch(ctx context.Context, in <-chan int, flush <-chan time.Time, size int) <-chan []int` | Full batches flush immediately, timer events flush partial batches, and cancellation flushes or discards according to the documented policy before closing. |
| 8 | `runRateLimited(ctx context.Context, tokens <-chan struct{}, workers int, jobs []int) ([]int, error)` | Every started job consumes a token, workers stop on context cancellation, and output/error lifecycle closes cleanly. |
| 9 | `serve(ctx context.Context, jobs <-chan int) (<-chan int, <-chan struct{})` | Shutdown stops accepting work, waits for active work, closes results, then closes done. |

- [ ] **Step 1: Write deterministic injected-event tests.**

Exercise 7 receives explicit flush events instead of constructing a real ticker in the test. Exercise 8 receives explicit tokens. Exercise 9 uses a start gate and a caller cancellation to distinguish “stopped accepting” from “active work joined.”

- [ ] **Step 2: Run red and inspect each cleanup obligation.**

A batcher that never flushes a partial batch, a pool that blocks on token acquisition after cancel, or a service that closes results before workers finish must fail.

- [ ] **Step 3: Implement the minimum solutions.**

Keep a single batch owner, stop and drain timer resources only where the implementation creates them, select on context around every blocking operation, and close result/done signals in coordinator order.

- [ ] **Step 4: Verify and commit.**

Run normal focused tests, selective race if metrics require it, `go vet`, and commit:

```sh
git add exercises/24_concurrency_patterns/concurrency7 exercises/24_concurrency_patterns/concurrency8 exercises/24_concurrency_patterns/concurrency9 solutions/24_concurrency_patterns/concurrency7 solutions/24_concurrency_patterns/concurrency8 solutions/24_concurrency_patterns/concurrency9 race.list
git commit -m "feat: add timed batching and graceful service shutdown"
```

---

### Task 4: Add fan-in failures, ordered deadlines, and retry policy 10–12

**Files:**
- Create: `exercises/24_concurrency_patterns/concurrency10..12/{main.go,main_test.go}`
- Create: `solutions/24_concurrency_patterns/concurrency10..12/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 10 | `merge(ctx context.Context, sources ...<-chan result) <-chan result` | Multiple sources forward typed failures; one closer closes output after all sources stop. |
| 11 | `ordered(ctx context.Context, workers int, jobs []job) ([]int, error)` | Indexed results preserve order until a deadline or error stops remaining work. |
| 12 | `retry(ctx context.Context, attempts int, backoff <-chan time.Duration, work func() error) error` | Only retryable errors retry; context cancellation stops backoff waiting and work. |

- [ ] **Step 1: Write table-driven tests for success, failure, and cancellation.**

Exercise 10 must include an empty source and a source that returns an error envelope. Exercise 11 must include a slow job after a deadline and assert no leaked result sender. Exercise 12 must cover success after a retry, exhausted retryable error, permanent error, and cancellation during backoff.

- [ ] **Step 2: Run red tests and enforce semantic error assertions.**

Use `errors.Is` for sentinel classification in tests; do not compare only formatted strings. An untouched starter returning nil must fail on a sentinel failure case.

- [ ] **Step 3: Implement error propagation, ordered collection, and cancellable retry.**

Use a coordinator for fan-in closure, indexed result storage for order, and select between backoff events and `ctx.Done()`. Do not introduce a third-party retry or errgroup package.

- [ ] **Step 4: Verify and commit.**

Run focused tests twice, race only if the implementation has shared mutable metrics, vet all three packages, and commit:

```sh
git add exercises/24_concurrency_patterns/concurrency10 exercises/24_concurrency_patterns/concurrency11 exercises/24_concurrency_patterns/concurrency12 solutions/24_concurrency_patterns/concurrency10 solutions/24_concurrency_patterns/concurrency11 solutions/24_concurrency_patterns/concurrency12
git commit -m "feat: add integrated error and deadline patterns"
```

---

### Task 5: Add shared initialization and atomic service metrics 13–15

**Files:**
- Create: `exercises/24_concurrency_patterns/concurrency13..15/{main.go,main_test.go}`
- Create: `solutions/24_concurrency_patterns/concurrency13..15/main.go`
- Modify: `race.list` to include `concurrency13`, `concurrency14`, and `concurrency15`.

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 13 | `loadConfig(ctx context.Context, load func(context.Context) string) string` | `sync.Once` initializes once, cancellation is returned consistently, and all callers see one stable result. |
| 14 | `runMeasured(ctx context.Context, workers int, jobs []int) (completed, canceled int64)` | Atomic counters equal post-join worker outcomes and race detector is clean. |
| 15 | `type job struct{ value int }`, `type result struct{ value int; accepted bool }`; `submit(ctx context.Context, queue chan<- job, result <-chan result, capacity int) error` | Bounded submission either enqueues or returns an explicit rejected/canceled result without blocking forever. |

- [ ] **Step 1: Write race-sensitive tests first.**

Run 20 concurrent config callers and assert the loader count is one. Run multiple workers with successful and canceled jobs, then read metrics only after a done signal. Fill a bounded queue without a consumer and require `submit` to return the documented rejection or cancellation.

- [ ] **Step 2: Run red tests with and without race.**

The semantic tests must fail independently of race scheduling. Add these exact paths to `race.list` because their purpose includes shared state or bounded counters.

- [ ] **Step 3: Implement Once, atomic metrics, and load shedding.**

Use `sync.Once` around initialization, `atomic.Int64` for counters, and a non-blocking or context-aware send for the bounded queue. Never read live metrics while workers are still mutating them.

- [ ] **Step 4: Verify and commit.**

Run `go test -count=1`, `go test -race`, `go vet`, and commit:

```sh
git add exercises/24_concurrency_patterns/concurrency13 exercises/24_concurrency_patterns/concurrency14 exercises/24_concurrency_patterns/concurrency15 solutions/24_concurrency_patterns/concurrency13 solutions/24_concurrency_patterns/concurrency14 solutions/24_concurrency_patterns/concurrency15 race.list
git commit -m "feat: add integrated metrics and backpressure patterns"
```

---

### Task 6: Add buffered pipelines, deadline shutdown, and capstone 16–18

**Files:**
- Create: `exercises/24_concurrency_patterns/concurrency16..18/{main.go,main_test.go}`
- Create: `solutions/24_concurrency_patterns/concurrency16..18/main.go`
- Modify: `race.list` by adding the exact entries `24_concurrency_patterns/concurrency13`, `24_concurrency_patterns/concurrency14`, `24_concurrency_patterns/concurrency15`, and `24_concurrency_patterns/concurrency18`.

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 16 | `bufferedPipeline(ctx context.Context, in <-chan int, buffer int) <-chan int` | Stage buffers improve decoupling but cancellation still interrupts both directions and output closes. |
| 17 | `shutdown(ctx context.Context, workers int, jobs []job) ([]result, error)` | In-flight work observes deadline, timers stop, workers join, and result output closes. |
| 18 | `runService(ctx context.Context, workers, limit int, requests []request) ([]response, error)` | Request/reply, bounded workers, ordered responses, first-error cancellation, timeout, metrics, and graceful shutdown are all observable. |

- [ ] **Step 1: Write the capstone tests before implementation.**

For every final exercise, include normal completion, empty input, cancellation before work starts, cancellation while output is blocked, one permanent error, output closure, and final worker completion. Use deterministic work gates and injected functions.

- [ ] **Step 2: Run the red phase.**

Run the starter verifier and ensure a zero slice/nil error implementation fails on response count, error, and done assertions. Check that no test passes merely because no work was started.

- [ ] **Step 3: Implement the integrated solutions.**

Use one context per request run, one coordinator per output close, a bounded queue or semaphore for the limit, indexed responses for order, a capacity-one error channel, and a final join before returning. Stop every timer/ticker created by the implementation.

- [ ] **Step 4: Run fast verification and selective race.**

Run normal focused tests twice, `go test -race` only for manifest-listed final packages, and `go vet ./solutions/24_concurrency_patterns/...`.

- [ ] **Step 5: Commit the final integration batch.**

```sh
git add exercises/24_concurrency_patterns/concurrency16 exercises/24_concurrency_patterns/concurrency17 exercises/24_concurrency_patterns/concurrency18 solutions/24_concurrency_patterns/concurrency16 solutions/24_concurrency_patterns/concurrency17 solutions/24_concurrency_patterns/concurrency18 race.list
git commit -m "feat: complete integrated concurrency capstones"
```

---

### Task 7: Write the integrated chapter guide and root index update

**Files:**
- Create: `exercises/24_concurrency_patterns/README.md`
- Modify: `README.md`

- [ ] **Step 1: Document prerequisites.**

State that chapters 12, 13, 15, 16, and 21 provide the primitives used here. Link each group of exercises to the exact earlier skill it combines.

- [ ] **Step 2: Document the 18-exercise sequence.**

For every row, state the new lifecycle obligation: context propagation, output closure, first-error cancellation, timer cleanup, bounded submission, atomic metric timing, or graceful shutdown.

- [ ] **Step 3: Document testing and race policy.**

Explain normal checks, selective `--race`, `--race-all`, starter red verification, and why deadlock tests use explicit completion signals plus bounded watchdogs.

- [ ] **Step 4: Update root README.**

Set `24_concurrency_patterns` to 18, link the guide, and keep it after `21_time` in the recommended progression.

- [ ] **Step 5: Format and commit.**

Run `git diff --check` and commit:

```sh
git add exercises/24_concurrency_patterns/README.md README.md
git commit -m "docs: guide integrated concurrency patterns"
```

---

### Task 8: Integrated verification checkpoint

**Files:**
- Verify: `exercises/24_concurrency_patterns`, `solutions/24_concurrency_patterns`, `race.list`, `check.sh`, and root documentation.

- [ ] **Step 1: Verify layout parity.**

Run the CI directory parity command and confirm exactly `concurrency1..18` exist under both trees.

- [ ] **Step 2: Verify pristine starters.**

Run `sh scripts/verify_exercise_starters.sh exercises/24_concurrency_patterns`; every starter must be rejected or fail its focused test.

- [ ] **Step 3: Run package tests and vet.**

Run `go test ./solutions/24_concurrency_patterns/...`, `go vet ./solutions/24_concurrency_patterns/...`, and `sh check.sh solutions/24_concurrency_patterns --run-all`.

- [ ] **Step 4: Run selective race and full audit.**

Run `sh check.sh solutions/24_concurrency_patterns --run-all --race`; inspect that only manifest entries use `-race`. Run `--race-all` after the selective suite passes.

- [ ] **Step 5: Review cross-chapter boundaries.**

Confirm no integrated exercise is merely a renamed raw channel drill, every context-aware stage checks both receive and send cancellation, every error path joins workers, and no timer/ticker remains running after return.
