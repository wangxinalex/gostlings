# Cancellation, Synchronization, and Time Exercise Expansion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task with review checkpoints.

**Goal:** Expand `15_context` to 14 exercises, `16_sync` to 14 exercises, and `21_time` to 8 exercises with practical repetition, strict starter-failure validation, and selective race checks.

**Architecture:** Each chapter stays focused on its standard-library primitive. Context exercises use `context.Context` and deterministic cancellation gates; sync exercises use shared memory and `sync` primitives rather than channel protocols; time exercises own timer/ticker construction and cleanup. Integrated combinations are deferred to `24_concurrency_patterns`.

**Tech Stack:** Go 1.26 standard library; `context`, `sync`, `sync/atomic` only where explicitly required by the integrated chapter, `time`, `fmt`, `errors`, and `testing`.

## Global Constraints

- Modify only `/Users/wangxinalex/SelfStudy/Rust/gostlings`; do not touch `working-progress`.
- Use `apply_patch` for edits and keep exercise/solution directory parity.
- Every new or rewritten learner `main.go` contains a `TODO:` seam and fails the checker until that seam is solved.
- Run a red-phase test for every exercise before writing its solution; a pristine starter that passes is rejected as insufficiently specified.
- Context exercises do not teach raw fan-in/fan-out or channel ownership; sync exercises do not use channels as the primary synchronization lesson; time exercises do not teach general worker-pool protocols.
- Avoid `time.Sleep` as synchronization. Use gates, channels, state predicates, or bounded watchdogs.
- Add only shared-memory exercises with meaningful race coverage to `race.list`; `--race-all` remains the expensive full audit.
- Do not run `go test ./...` over intentionally incomplete learner packages.

## File map

- Modify or create `exercises/15_context/context1..14/{main.go,main_test.go}` and matching solutions.
- Modify or create `exercises/16_sync/sync1..14/{main.go,main_test.go}` and matching solutions.
- Modify or create `exercises/21_time/time1..8/{main.go,main_test.go}` and matching solutions.
- Update `race.list` with the exact entries `16_sync/sync1` through `16_sync/sync14`; do not add context or time paths because their tests use cancellation and timer state assertions rather than shared mutable state.
- Create or replace `exercises/15_context/README.md`, `exercises/16_sync/README.md`, and `exercises/21_time/README.md`.
- Modify `README.md` counts and descriptions for these three chapters.

---

### Task 1: Rebuild context exercises 1–7

**Files:**
- Replace: `exercises/15_context/context1..7/{main.go,main_test.go}`
- Replace: `solutions/15_context/context1..7/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 1 | `worker(ctx context.Context) string` | Cancelled context produces the cancellation result; the worker does not wait for the long branch. |
| 2 | `run() string` | Timeout result arrives through a capacity-one result channel and `cancel` is called. |
| 3 | `handler(ctx context.Context) string` | Typed key lookup returns the user value and missing value returns the documented fallback. |
| 4 | `classify(ctx context.Context) error` | `context.Canceled` and `context.DeadlineExceeded` are distinguished with `errors.Is`. |
| 5 | `runUntil(ctx context.Context, deadline time.Time) string` | `WithDeadline` stops work at the absolute deadline and the returned cancel function is called. |
| 6 | `startChildren(ctx context.Context, count int) <-chan struct{}` | Parent cancellation reaches every child and the group completion signal closes. |
| 7 | `cancelChild(parent context.Context) (parent context.Context, child context.Context)` | Child cancellation does not cancel the parent or a sibling derived from the parent. |

- [ ] **Step 1: Write the failing tests.**

Use a gate channel to hold workers in their work branch and cancel the context explicitly. A core test must look like:

```go
func TestWorkerStopsWhenContextIsCanceled(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := worker(ctx); got != "worker: canceled" {
		t.Fatalf("worker() = %q, want cancellation result", got)
	}
}
```

For exercise 7, wait on `child.Done()` and assert `parent.Err()` is nil after canceling the child.

- [ ] **Step 2: Run the red phase.**

Run `sh scripts/verify_exercise_starters.sh exercises/15_context` after each pristine batch is prepared. The TODO gate must reject untouched starters, and the focused tests must fail for an edited-but-incomplete implementation.

- [ ] **Step 3: Implement the minimum context solutions.**

Call the cancel function with `defer` when a child context is created. Select on `ctx.Done()` for cooperative stop; do not close `ctx.Done()` directly and do not replace a caller context with `context.Background()` inside a helper.

- [ ] **Step 4: Verify the green phase.**

Run `go test -count=1 ./solutions/15_context/...`, `go vet ./solutions/15_context/...`, and the matching learner tests only after applying their intended fixes.

- [ ] **Step 5: Commit the first context batch.**

```sh
git add exercises/15_context/context1 exercises/15_context/context2 exercises/15_context/context3 exercises/15_context/context4 exercises/15_context/context5 exercises/15_context/context6 exercises/15_context/context7 solutions/15_context/context1 solutions/15_context/context2 solutions/15_context/context3 solutions/15_context/context4 solutions/15_context/context5 solutions/15_context/context6 solutions/15_context/context7
git commit -m "feat: expand context cancellation fundamentals"
```

---

### Task 2: Add context propagation, cooperative workers, and cancel causes 8–14

**Files:**
- Create: `exercises/15_context/context8..14/{main.go,main_test.go}`
- Create: `solutions/15_context/context8..14/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 8 | `receive(ctx context.Context, in <-chan int) (int, error)` | Blocked input returns `ctx.Err()` after cancellation. |
| 9 | `send(ctx context.Context, out chan<- int, value int) error` | Blocked output returns `ctx.Err()` after cancellation. |
| 10 | `runWorkers(ctx context.Context, count int) <-chan struct{}` | All workers observe cancellation and the final done channel closes after every worker returns. |
| 11 | `startIfActive(ctx context.Context, work func()) bool` | Already-canceled context prevents work from starting. |
| 12 | `lookup(ctx context.Context, source func(context.Context) string) string` | The same context is passed through the helper chain and cancellation remains observable. |
| 13 | `runWithCause(ctx context.Context) error` | `WithCancelCause` preserves a sentinel cause retrievable through `context.Cause`. |
| 14 | `runRequest(ctx context.Context, workers int) (string, error)` | Deadline, propagation, cooperative stop, result delivery, and cleanup all hold together. |

- [ ] **Step 1: Write failure-path tests before implementation.**

Exercise 8 must use an unbuffered input with no sender; exercise 9 must use an unbuffered output with no receiver; exercise 10 must expose a completion channel; exercise 11 must use an atomic test-side counter only if it is listed in `race.list`, otherwise use a channel gate and a boolean result.

- [ ] **Step 2: Run the red phase and inspect failure reasons.**

An implementation that only checks cancellation before entering a loop must fail the blocked receive/send tests. An implementation that returns an error but leaves workers running must fail the completion assertion.

- [ ] **Step 3: Implement exercises 8–14.**

Use `select { case value := <-in: ...; case <-ctx.Done(): ... }` for receives and the analogous send select. Use one coordinator to close a group completion channel after all workers exit. Use `context.WithCancelCause` only in exercise 13; keep cause classification separate from error-string formatting.

- [ ] **Step 4: Verify normal behavior and repeated cancellation.**

Run each solution package twice with `go test -count=1`, then run the cancellation tests repeatedly in a shell loop. Confirm no test depends on a fixed sleep.

- [ ] **Step 5: Update the context guide and commit.**

Document raw channel cancellation as the prerequisite and context cancellation as the standard request-scoped API. Then commit:

```sh
git add exercises/15_context/context8 exercises/15_context/context9 exercises/15_context/context10 exercises/15_context/context11 exercises/15_context/context12 exercises/15_context/context13 exercises/15_context/context14 solutions/15_context/context8 solutions/15_context/context9 solutions/15_context/context10 solutions/15_context/context11 solutions/15_context/context12 solutions/15_context/context13 solutions/15_context/context14 exercises/15_context/README.md
git commit -m "feat: deepen context propagation and cancellation"
```

---

### Task 3: Rebuild sync exercises 1–7

**Files:**
- Replace: `exercises/16_sync/sync1..7/{main.go,main_test.go}`
- Replace: `solutions/16_sync/sync1..7/main.go`
- Modify: `race.list` to include `16_sync/sync1` through `16_sync/sync7`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 1 | `main()` protected counter | Exactly 1000 increments; race detector remains clean. |
| 2 | `main()` protected map | Exactly 10000 entries; all writes use the same mutex. |
| 3 | `initConfig()` called by many goroutines | Initialization output/state occurs exactly once. |
| 4 | `waitForWorkers(count int) int` | Empty and non-empty groups return only after every worker reports done. |
| 5 | `runTasks(jobs []int) int` | All known tasks are added before waiting; dynamic input and empty input are correct. |
| 6 | `readWriteCache` with `Get` and `Put` | Concurrent readers do not race with exclusive writers and values remain consistent. |
| 7 | `updateState` with a protected snapshot | Slow work occurs outside the lock; the committed state is still race-free. |

- [ ] **Step 1: Write real race-sensitive tests.**

For exercise 1, run 100 goroutines with 10 increments each and assert 1000. For exercise 2, assert length and selected key/value pairs. For exercise 6, launch readers and writers, wait, then assert a stable final value.

- [ ] **Step 2: Run red tests without `-race` and with `-race`.**

The starter must fail the semantic assertion even if a particular scheduler run does not expose a race. The race command is additional evidence, not the only correctness test.

- [ ] **Step 3: Implement mutex, Once, WaitGroup, and RWMutex solutions.**

Lock around the read-modify-write operation, use `defer Unlock` in goroutine paths, call `once.Do` exactly around initialization, add before launch, defer `Done`, and wait only after all adds. Use `RLock`/`RUnlock` for cache reads and `Lock`/`Unlock` for writes.

- [ ] **Step 4: Run the green tests and race subset.**

Run `go test -count=1 ./solutions/16_sync/...`, `go test -race ./solutions/16_sync/...`, and `go vet ./solutions/16_sync/...`.

- [ ] **Step 5: Commit the sync foundation batch.**

```sh
git add exercises/16_sync/sync1 exercises/16_sync/sync2 exercises/16_sync/sync3 exercises/16_sync/sync4 exercises/16_sync/sync5 exercises/16_sync/sync6 exercises/16_sync/sync7 solutions/16_sync/sync1 solutions/16_sync/sync2 solutions/16_sync/sync3 solutions/16_sync/sync4 solutions/16_sync/sync5 solutions/16_sync/sync6 solutions/16_sync/sync7 race.list
git commit -m "feat: expand mutex and waitgroup practice"
```

---

### Task 4: Add advanced sync primitives 8–14

**Files:**
- Create: `exercises/16_sync/sync8..14/{main.go,main_test.go}`
- Create: `solutions/16_sync/sync8..14/main.go`
- Modify: `race.list` to include `16_sync/sync8` through `16_sync/sync14`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 8 | `config() string` using `sync.Once` | Concurrent callers receive one initialized value and initialization count is one. |
| 9 | `queue.Push`/`queue.Pop` using `sync.Cond` | Pop waits while predicate is false and resumes after signal. |
| 10 | `queue.Close` using `sync.Cond.Broadcast` | All waiters wake when the queue closes and return the documented closed state. |
| 11 | `registry` using `sync.Map` | `Load`, `Store`, and `LoadOrStore` preserve one canonical value per key. |
| 12 | `reuseBuffer` using `sync.Pool` | Objects can be returned and reused, but correctness does not depend on a particular reuse. |
| 13 | `boundedState` using mutex plus condition signaling | Producers wait at the limit and consumers release capacity without races. |
| 14 | `serviceState` combining Once, protected state, wait, and shutdown | Initialization occurs once, final state is stable after join, and no goroutine remains waiting. |

- [ ] **Step 1: Write tests for Once and Cond predicates.**

For exercise 9, call `Pop` before `Push` in a goroutine and use a channel to prove it is waiting; then push and require the value. For exercise 10, start two waiters, close the queue, and require both to return.

- [ ] **Step 2: Write tests for sync.Map, Pool, and bounded state.**

The Pool test must assert object contents after `Get`; it must not require a specific reuse because the runtime may clear the pool. The bounded-state test must assert that active count never exceeds the bound after all work joins.

- [ ] **Step 3: Run red tests and implement exercises 8–14.**

Always wait in a `for` loop around a condition predicate. Hold the condition lock while checking and changing state. Use `LoadOrStore` for canonical registry publication and make shutdown wake all condition waiters.

- [ ] **Step 4: Run normal, race, and vet checks.**

Run focused tests with `-count=1`, `go test -race ./solutions/16_sync/...`, and `go vet ./solutions/16_sync/...`. Repeat Cond and bounded-state tests at least five times.

- [ ] **Step 5: Update the sync guide and commit.**

Document why the chapter uses shared memory while channels own message flow, and include a table mapping each primitive to its common production use and cleanup rule. Commit:

```sh
git add exercises/16_sync/sync8 exercises/16_sync/sync9 exercises/16_sync/sync10 exercises/16_sync/sync11 exercises/16_sync/sync12 exercises/16_sync/sync13 exercises/16_sync/sync14 solutions/16_sync/sync8 solutions/16_sync/sync9 solutions/16_sync/sync10 solutions/16_sync/sync11 solutions/16_sync/sync12 solutions/16_sync/sync13 solutions/16_sync/sync14 exercises/16_sync/README.md race.list
git commit -m "feat: add advanced sync primitives"
```

---

### Task 5: Rebuild time exercises 1–4

**Files:**
- Replace: `exercises/21_time/time1..4/{main.go,main_test.go}`
- Replace: `solutions/21_time/time1..4/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 1 | `main()` layout formatting | Exact reference-time layout prints the documented date. |
| 2 | `main()` ticker loop | Exactly two ticks print and ticker stops before return. |
| 3 | `durationReport(d time.Duration) string` | Duration arithmetic preserves units and boundary values. |
| 4 | `waitTimer(stop <-chan struct{}, d time.Duration) <-chan time.Time` | Timer produces one event and is stopped when the caller cancels early. |

- [ ] **Step 1: Write tests for formatting and duration boundaries.**

Use a fixed UTC `time.Time`, explicit duration values, and output capture. For exercise 2, use a short controlled interval only for the solution example; the test must count exact output and cap total runtime.

- [ ] **Step 2: Write timer lifecycle tests.**

Test the normal timer event and the cancellation/early-return path. Assert the returned done signal closes; do not inspect internal runtime goroutine counts.

- [ ] **Step 3: Run red tests, implement solutions, and verify.**

Use the Go reference time layout, `time.Duration` constants, `time.NewTimer`, `defer timer.Stop()`, and a bounded tick loop. Do not add these packages to `race.list` unless a later test introduces shared mutable state.

- [ ] **Step 4: Commit the time foundation batch.**

```sh
git add exercises/21_time/time1 exercises/21_time/time2 exercises/21_time/time3 exercises/21_time/time4 solutions/21_time/time1 solutions/21_time/time2 solutions/21_time/time3 solutions/21_time/time4
git commit -m "feat: expand timer and duration fundamentals"
```

---

### Task 6: Add timer reset, ticker cleanup, and periodic boundaries 5–8

**Files:**
- Create: `exercises/21_time/time5..8/{main.go,main_test.go}`
- Create: `solutions/21_time/time5..8/main.go`

**Interfaces and test contracts:**

| Exercise | Interface | Required assertion |
| --- | --- | --- |
| 5 | `awaitOrCancel(result <-chan string, timeout time.Duration) string` | Result, cancellation, and timeout branches produce distinct outcomes. |
| 6 | `reuseTimer(gates <-chan struct{}) int` | Timer is stopped/drained before reset and produces one event per gate. |
| 7 | `runTicker(stop <-chan struct{}, ticks <-chan time.Time) <-chan struct{}` | Ticker-driven work stops on every exit and done closes. |
| 8 | `periodic(ctx context.Context, interval, deadline time.Duration) int` | Periodic work stops at context/deadline and ticker cleanup is deterministic. |

- [ ] **Step 1: Write failing select and timer-reset tests.**

Inject result, cancel, and tick channels where possible. For exercise 6, send explicit gate events and assert the count; do not sleep to wait for a timer.

- [ ] **Step 2: Run the red phase.**

Confirm a timer that is not stopped/drained, a ticker that is not stopped, or a loop that ignores cancellation fails a focused assertion or watchdog.

- [ ] **Step 3: Implement and verify exercises 5–8.**

Use `time.NewTimer`, `Stop`, conditional channel drain, `Reset`, `time.NewTicker`, `defer ticker.Stop`, and context/deadline only for the final integration within the time chapter. Keep rate limiting itself in channels as an injected token stream.

- [ ] **Step 4: Update the time guide and root index.**

Document timer versus ticker, safe reset order, stop-on-all-paths, and the handoff to channel token streams and integrated patterns.

- [ ] **Step 5: Commit the time batch.**

```sh
git add exercises/21_time/time5 exercises/21_time/time6 exercises/21_time/time7 exercises/21_time/time8 solutions/21_time/time5 solutions/21_time/time6 solutions/21_time/time7 solutions/21_time/time8 exercises/21_time/README.md
git commit -m "feat: add timer reset and periodic cleanup practice"
```

---

### Task 7: Update root curriculum and run the primitive-chapter checkpoint

**Files:**
- Modify: `README.md`
- Verify: `race.list`, `check.sh`, `scripts/verify_exercise_starters.sh`, all changed primitive chapters.

- [ ] **Step 1: Update root topic counts.**

Set `15_context` to 14, `16_sync` to 14, and `21_time` to 8. Explain that `24_concurrency_patterns` follows these primitive chapters.

- [ ] **Step 2: Run starter red verification.**

Run `sh scripts/verify_exercise_starters.sh exercises/15_context exercises/16_sync exercises/21_time`. Every pristine starter must be rejected or fail its focused test.

- [ ] **Step 3: Run normal solutions and vet.**

Run `go test ./solutions/15_context/... ./solutions/16_sync/... ./solutions/21_time/...`, `go vet ./solutions/15_context/... ./solutions/16_sync/... ./solutions/21_time/...`, and `sh check.sh solutions/15_context --run-all`, `sh check.sh solutions/16_sync --run-all`, `sh check.sh solutions/21_time --run-all`.

- [ ] **Step 4: Run selective race checks.**

Run `sh check.sh solutions/15_context --run-all --race`, `sh check.sh solutions/16_sync --run-all --race`, and `sh check.sh solutions/21_time --run-all --race`. Confirm only manifest-listed sync packages receive `-race`.

- [ ] **Step 5: Run the full audit only after fast checks pass.**

Run `sh check.sh solutions/15_context --run-all --race-all`, `sh check.sh solutions/16_sync --run-all --race-all`, and `sh check.sh solutions/21_time --run-all --race-all`.

- [ ] **Step 6: Verify formatting and parity, then commit the curriculum update.**

Run `gofmt -d $(git ls-files '*.go')`, `git diff --check`, and the CI directory parity command. Commit:

```sh
git add README.md exercises/15_context/README.md exercises/16_sync/README.md exercises/21_time/README.md race.list
git commit -m "docs: guide expanded context sync and time practice"
```
