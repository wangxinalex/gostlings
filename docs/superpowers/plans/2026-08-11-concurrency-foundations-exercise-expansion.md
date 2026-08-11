# Concurrency Foundations Exercise Expansion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task with review checkpoints.

**Goal:** Expand `12_goroutines` to 10 sequential exercises and `13_channels` to 50 sequential exercises, with repeated high-frequency patterns, strict starter failure checks, and fast selective race verification.

**Architecture:** Each learner exercise is an independent Go `main` package with a deliberately incomplete `main.go`, focused `main_test.go`, and a matching runnable solution. Existing channel pipeline exercises move to `24_concurrency_patterns`; channel 45 consumes externally supplied tokens instead of constructing a ticker. `check.sh` validates the learner TODO gate before testing and consults one root race manifest.

**Tech Stack:** Go 1.26 standard library; `fmt`, `sync`, `time`, `errors`, `reflect`, `sort`, and `testing`; POSIX `sh` for checker utilities.

## Global Constraints

- Modify only `/Users/wangxinalex/SelfStudy/Rust/gostlings`; do not touch `working-progress`.
- Use `apply_patch` for source and documentation edits.
- Keep `exercises/<topic>/<exercise>` and `solutions/<topic>/<exercise>` in exact directory parity.
- Keep learner starters intentionally incomplete and require a `TODO:` seam in every new `main.go`.
- `check.sh` must reject an exercise target while its starter `main.go` still contains `TODO:`; only the internal `--verify-starter` mode may bypass that gate to prove the focused test itself is red.
- Every new exercise must fail against its pristine starter before its solution is written; a passing pristine starter is a test-design failure.
- Use explicit channels, closed outputs, state assertions, or bounded watchdogs for synchronization; do not use `time.Sleep` as the only proof of completion.
- `13_channels` may use raw stop channels and token channels, but must not import `context` or teach timer/ticker construction.
- Do not run `go test ./...` over the learner tree; learner packages are intentionally incomplete.
- `--race` is selective through `race.list`; `--race-all` is the explicit full audit mode; `--verify-starter` is used only by the starter verifier.
- Run `gofmt` after each code batch and commit each coherent batch.

## File map

- Modify `check.sh` to add the TODO gate, verifier-only bypass, selective race lookup, and `--race-all`.
- Create `race.list` with these initial shared-state paths relative to `exercises/`: `16_sync/sync1` through `16_sync/sync14`; later integration tasks append `24_concurrency_patterns/concurrency3`, `concurrency5`, `concurrency13`, `concurrency14`, `concurrency15`, and `concurrency18`.
- Create `scripts/verify_exercise_starters.sh` for red-phase validation of pristine starters.
- Create or replace `exercises/12_goroutines/goroutines1..10/{main.go,main_test.go}` and matching solutions.
- Create or replace `exercises/13_channels/channels1..50/{main.go,main_test.go}` and matching solutions.
- Move the former raw pipeline content out of `exercises/13_channels/channels19..21` during the channel renumbering; the integration plan creates its new home under `24_concurrency_patterns`.
- Update `exercises/12_goroutines/README.md`, `exercises/13_channels/README.md`, and the root `README.md`.

---

### Task 1: Add strict starter validation and selective race execution

**Files:**
- Modify: `check.sh`
- Create: `race.list`
- Create: `scripts/verify_exercise_starters.sh`

**Interfaces:**
- `sh check.sh <target> [--run-all] [--race] [--race-all] [--verify-starter]`
- `race.list`: one non-comment path per line, relative to `exercises/`, for example `16_sync/sync1`.
- `sh scripts/verify_exercise_starters.sh exercises/12_goroutines exercises/13_channels`

- [ ] **Step 1: Write the checker regression cases.**

Create a temporary fixture under `${TMPDIR:-/tmp}` containing one exercise with `main.go` that prints the expected value but still contains `// TODO: implement`, and one solution-like directory with no TODO. The expected checker behavior is: the exercise target fails before running `go test`; `--verify-starter` bypasses only the static gate so the same fixture can prove its focused test is red; the solution target runs normally.

Also create a fixture with a listed race path and an unlisted path. Record that `--race` adds `-race` only to the listed path, while `--race-all` adds it to both. Keep the fixture commands in the plan's eventual verification notes rather than committing generated files.

- [ ] **Step 2: Run the checker regression cases and observe failure.**

Run the fixture commands through the current `check.sh` and confirm the TODO fixture incorrectly reaches the test command, the verifier-only bypass is unavailable, and the selective race behavior is not available. This is the red phase for the checker change.

- [ ] **Step 3: Implement the TODO gate and verifier-only bypass.**

Before constructing the Go command, when the normalized target is under `exercises/`, inspect each selected `main.go`. If it contains `TODO:` and `--verify-starter` is absent, print `FAIL: <path> (starter still contains TODO)` and count it as a failure. Do not apply this gate to `solutions/`. Reject `--verify-starter` for `solutions/` and for targets outside the exercise tree.

- [ ] **Step 4: Implement manifest-based race selection.**

Normalize a selected path to its relative `topic/exercise` form. When `--race` is set, map the path to `race.list`; append `-race` only when an exact manifest entry exists. When `--race-all` is set, append `-race` unconditionally. Reject simultaneous `--race` and `--race-all` with usage text and exit status 2.

- [ ] **Step 5: Implement the starter verifier.**

The script must:

1. accept one or more topic directories;
2. enumerate numeric exercise directories in order;
3. require `main.go`, `main_test.go`, and a `TODO:` marker in `main.go`;
4. run `sh check.sh <exercise> --verify-starter` and require a non-zero result from compilation or its focused tests;
5. print the first unexpected passing exercise and exit 1.

This script is run only against a pristine starter snapshot before solutions are copied. It is not part of the learner's normal solve loop.

- [ ] **Step 6: Run the checker tests again.**

Run the fixture cases, `sh -n check.sh`, `sh -n scripts/verify_exercise_starters.sh`, and `git diff --check`. Expected result: the TODO fixture fails before tests, selective race behavior is visible in the command trace, and the full-race mode remains available.

- [ ] **Step 7: Commit the checker infrastructure.**

```sh
git add check.sh race.list scripts/verify_exercise_starters.sh
git commit -m "test: enforce incomplete exercises and selective race checks"
```

---

### Task 2: Rebuild the 10 goroutine exercises

**Files:**
- Create or replace: `exercises/12_goroutines/goroutines1..10/main.go`
- Create or replace: `exercises/12_goroutines/goroutines1..10/main_test.go`
- Create or replace: `solutions/12_goroutines/goroutines1..10/main.go`

**Interfaces and test contracts:**

| Exercise | Learner interface | Required red/green assertion |
| --- | --- | --- |
| 1 | `main()` launches one greeting task | Capture output with a bounded completion channel; pristine synchronous main exits before the worker prints. |
| 2 | `runLabels(labels []string) []string` | Every loop value is represented once; the starter captures the loop variable incorrectly. |
| 3 | `runWithArgs(labels []string) []string` | Explicit goroutine arguments preserve each label and empty input returns an empty slice. |
| 4 | `runWorkers(count int) []string` | All workers report completion before return; missing `Add`, `Done`, or `Wait` must fail. |
| 5 | `runJobs(jobs []int, work func(int) string) []string` | Dynamic job count and zero jobs complete without a sleep-based wait. |
| 6 | `runEach(jobs []int, visit func(int)) int` | Return the exact number of completed visits; completion is recorded once per started goroutine. |
| 7 | `runBatches(batches [][]int) [][]string` | Batch two cannot finish before batch one is joined; output preserves batch boundaries. |
| 8 | `runWorkersWithInput(jobs []int) []string` | Workers receive immutable values through parameters and return one result per job. |
| 9 | `runWithEarlyReturn(jobs []int, stopAt int) []string` | Every launched worker reaches deferred completion even when its branch returns early. |
| 10 | `reviewRun(jobs []int) []string` | Combine explicit parameters, dynamic launch, deferred completion, and final join. |

- [ ] **Step 1: Write the failing tests for exercises 1–5.**

Use real functions, not mocks. For example, the core test for exercise 4 must be shaped like:

```go
func TestRunWorkersWaitsForEveryWorker(t *testing.T) {

	got := runWorkers(4)
	if len(got) != 4 {
		t.Fatalf("runWorkers(4) returned %d completions, want 4", len(got))
	}
}
```

The tests for exercises 2 and 3 must sort only when order is not part of the contract; exercise 7 must assert batch order. Add empty-input tests for exercises 3 and 5.

- [ ] **Step 2: Run exercises 1–5 against pristine starters.**

Run `sh scripts/verify_exercise_starters.sh exercises/12_goroutines`. Confirm every selected starter fails and each failure identifies output completion, captured values, or missing joins.

- [ ] **Step 3: Implement the minimum solutions for exercises 1–5.**

Use explicit goroutine parameters, `sync.WaitGroup.Add` before each launch, `defer wg.Done()` inside each goroutine, and `wg.Wait()` before returning. Do not introduce channels or mutexes as the lesson's solution.

- [ ] **Step 4: Run the green tests for exercises 1–5.**

Run `for dir in exercises/12_goroutines/goroutines1 exercises/12_goroutines/goroutines2 exercises/12_goroutines/goroutines3 exercises/12_goroutines/goroutines4 exercises/12_goroutines/goroutines5; do go test -count=1 "./$dir" || exit 1; done` after applying the learner fixes, then run `go test -count=1 ./solutions/12_goroutines/...`. Confirm empty and batch-order assertions pass.

- [ ] **Step 5: Write the failing tests for exercises 6–10.**

Exercise 6 must count actual visits; exercise 7 must use a gate or an observable batch completion record, not a timing sleep; exercise 9 must verify early-return workers still decrement the join count; exercise 10 must cover zero and several jobs.

- [ ] **Step 6: Run the red phase and implement exercises 6–10.**

Run the starter verifier first. Then implement the smallest `WaitGroup` lifecycle that satisfies the tests, keeping shared state out of the exercise unless the test owns it safely.

- [ ] **Step 7: Run the green phase and format the chapter.**

Run `gofmt -w exercises/12_goroutines solutions/12_goroutines`, focused tests for all ten solution packages, and `go vet ./solutions/12_goroutines/...`. Do not add goroutine paths to `race.list` unless a final test genuinely observes shared mutable state.

- [ ] **Step 8: Commit the goroutine batch.**

```sh
git add exercises/12_goroutines solutions/12_goroutines
git commit -m "feat: expand goroutine fundamentals"
```

---

### Task 3: Renumber and rebuild channel exercises 1–10

**Files:**
- Move or replace: `exercises/13_channels/channels1..10/{main.go,main_test.go}`
- Move or replace: `solutions/13_channels/channels1..10/main.go`

**Interfaces and test contracts:**

| Exercise | Learner interface | Required assertions |
| --- | --- | --- |
| 1 | `main()` unbuffered handoff | Main completes and prints `hi`; the starter remains blocked. |
| 2 | `main()` buffered sends | Capacity two permits two sends before receives. |
| 3 | `main()` close plus range | Range prints all values and terminates only after close. |
| 4 | `read(ch <-chan int) (int, bool)` | Closed receive is `(0,false)` while buffered zero is `(0,true)`. |
| 5 | `drainClosed(ch <-chan int) []int` | Values sent before close are returned before the closed state ends the loop. |
| 6 | `generate(values ...int) <-chan int` | Producer owns output closure; empty input closes promptly. |
| 7 | `generate(values ...int) <-chan int` plus directional helper | Caller cannot send or close the returned channel; all values and closure are preserved. |
| 8 | `receiveFast(fast, slow <-chan string) string` | Either ready input is accepted; a silent input cannot block a ready one. |
| 9 | `tryReceive(ch <-chan int) string` | Ready receive returns a value; empty receive returns `no value` immediately. |
| 10 | `trySend(ch chan<- int, value int) bool` | Ready send returns true; full or unbuffered-without-receiver returns false immediately. |

- [ ] **Step 1: Move existing directories with `git mv` and create new tests before editing starters.**

Preserve current behavior for exercises 1–4, move old channels5 to new channels6, old channels6 to new channels8, and old channels8 to new channels9. Create new channels5, 7, and 10 with tests that directly exercise their function seam.

- [ ] **Step 2: Run red tests and the starter verifier.**

Run `sh scripts/verify_exercise_starters.sh exercises/13_channels` for the first ten directories. A starter that prints the expected example but still has `TODO:` must fail the checker before the Go test runs.

- [ ] **Step 3: Implement solutions for the first ten exercises.**

Use only native channel operations, close ownership, directional return types, and `select` with `default`. Do not introduce `context`, a ticker, a mutex, or a busy polling loop.

- [ ] **Step 4: Verify and commit the semantic batch.**

Run `gofmt`, focused normal tests, `go vet ./solutions/13_channels/...`, and the starter verifier on a clean copy of the starters. Commit:

```sh
git add exercises/13_channels solutions/13_channels
git commit -m "feat: strengthen channel semantics and select basics"
```

---

### Task 4: Implement channel timeout, result, and cancellation exercises 11–20

**Files:**
- Create or replace: `exercises/13_channels/channels11..20/{main.go,main_test.go}`
- Create or replace: `solutions/13_channels/channels11..20/main.go`

**Interfaces and test contracts:**

| Exercise | Learner interface | Required assertions |
| --- | --- | --- |
| 11 | `await(ch <-chan string) string` | Ready result wins; silent input returns `timed out` within a bounded deadline. |
| 12 | `relay(stop <-chan struct{}, in <-chan int) <-chan int` | Relay handles receive and send readiness and closes output on input close or stop. |
| 13 | `complete() <-chan struct{}` | Returned done channel closes promptly and carries no data. |
| 14 | `runAsync(work func() int) <-chan int` | Worker can publish one result into a capacity-one result channel even before caller receives. |
| 15 | `runWithDone(done chan struct{}) <-chan int` | Result delivery and completion notification are separate; both close or become observable in the documented order. |
| 16 | `produce(stop <-chan struct{}) <-chan int` | A blocked output send is interrupted by stop and output closes. |
| 17 | `forward(stop <-chan struct{}, in <-chan int) <-chan int` | A blocked receive and a blocked send both respond to stop. |
| 18 | `startWorkers(count int, stop <-chan struct{}) <-chan struct{}` | One close broadcasts to every worker and done closes after all exit. |
| 19 | `run(done chan struct{}) string` | Timeout closes stop, waits for done, and returns `timed out`; success also cleans up. |
| 20 | `shutdown(stop chan struct{}, workers int) <-chan struct{}` | A coordinator closes stop once, waits for all workers, then closes done exactly once. |

- [ ] **Step 1: Write tests for timeout and relay behavior.**

Use buffered ready inputs, silent unbuffered inputs, and a `time.After` watchdog only around the test. For exercise 12, assert both normal input closure and cancellation while the output has no receiver.

- [ ] **Step 2: Run the red phase.**

Run each starter individually and confirm the untouched implementation fails or is rejected by the TODO gate. Fix tests if a default return value can satisfy them without touching the intended select or close seam.

- [ ] **Step 3: Implement solutions for exercises 11–15.**

Use `select` for timeout and two-way relay, `defer close(done)` for completion, a capacity-one result channel for a one-shot result, and separate channels for result data versus lifecycle notification.

- [ ] **Step 4: Run normal tests for exercises 11–15.**

Run `go test -count=1` on each exercise and solution package, then run the focused tests twice to expose scheduling-sensitive assertions.

- [ ] **Step 5: Write and run red tests for exercises 16–20.**

Each cancellation test must close stop without receiving output and require the returned output or done channel to close within 500ms. Exercise 20 must include a second observer or repeated shutdown call that exposes double-close bugs.

- [ ] **Step 6: Implement cancellation and shutdown solutions.**

Every potentially blocking send and receive gets a stop case. One coordinator owns output/done closure; workers never close a shared signal that another worker may also close.

- [ ] **Step 7: Verify, format, and commit.**

Run `gofmt`, focused tests, `go vet`, and the starter verifier. Commit:

```sh
git add exercises/13_channels solutions/13_channels
git commit -m "feat: add channel cancellation and shutdown practice"
```

---

### Task 5: Implement channel composition exercises 21–30

**Files:**
- Create: `exercises/13_channels/channels21..30/{main.go,main_test.go}`
- Create: `solutions/13_channels/channels21..30/main.go`

**Interfaces and test contracts:**

| Exercise | Learner interface | Required assertions |
| --- | --- | --- |
| 21 | `squareWorkers(workers int, jobs <-chan int) <-chan int` | One result per job; output closes after all workers stop. |
| 22 | `merge(inputs ...<-chan int) <-chan int` | All inputs forward; zero inputs close output. |
| 23 | `merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int` | Both input receive and output send are cancellable. |
| 24 | `squareWorkers(stop <-chan struct{}, workers int, jobs <-chan int) <-chan int` | Workers stop before a blocked jobs receive or result send strands them. |
| 25 | `merge(inputs ...<-chan int) <-chan int` | Empty, nil-free, already-closed, and buffered inputs all terminate correctly. |
| 26 | `run(workers int, jobs []int) []int` | Basic pool processes every job and empty input returns. |
| 27 | `runOrdered(workers int, jobs []int) []int` | Indexed results restore input order after parallel processing. |
| 28 | `run(workers int, jobs []job) error` | First error is returned, stop closes once, and all workers join. |
| 29 | `startPool(workers int) (chan<- int, <-chan int)` | Caller can submit jobs through a send-only handle and range a receive-only result stream. |
| 30 | `runBounded(workers, buffer int, jobs []int) []int` | Submission and processing make progress under a bounded queue without unbounded buffering. |

- [ ] **Step 1: Write focused fan-out/fan-in tests.**

Collect results with a select plus a 500ms deadline, sort only where completion order is intentionally unspecified, and always assert the output channel closes. Include zero inputs and a blocked input for exercises 22–25.

- [ ] **Step 2: Run the red phase and correct permissive tests.**

An implementation that returns a nil output, drops values, closes output from a worker, or ignores cancellation must fail a focused test. The starter verifier must reject any untouched TODO starter.

- [ ] **Step 3: Implement exercises 21–25.**

Use one coordinator-owned output close for every multi-sender pattern, and add stop cases around both directions in cancellable variants.

- [ ] **Step 4: Implement and test exercises 26–30.**

Use a jobs producer that closes jobs, a worker group that sends results, and a separate coordinator that closes results after joining workers. Exercise 29 must expose directional types in the function signature; exercise 30 must use a bounded channel or explicit bounded submission protocol.

- [ ] **Step 5: Verify and commit.**

Run normal and repeated focused tests, `go vet ./solutions/13_channels/...`, and the starter verifier. Commit:

```sh
git add exercises/13_channels solutions/13_channels
git commit -m "feat: add channel fan-in fan-out and worker pools"
```

---

### Task 6: Implement channel service protocol exercises 31–40

**Files:**
- Create: `exercises/13_channels/channels31..40/{main.go,main_test.go}`
- Create: `solutions/13_channels/channels31..40/main.go`

**Interfaces and test contracts:**

| Exercise | Learner interface | Required assertions |
| --- | --- | --- |
| 31 | `firstResult(tasks []func(<-chan struct{}) string) string` | First completed result is returned; stop is broadcast and every task observes it before return. |
| 32 | `type request struct{ value int; reply chan int }`; `serve(requests <-chan request) <-chan struct{}` | Each request receives its own response; server stops after request input closes. |
| 33 | `collect(sources []<-chan int) <-chan int` | Multiple senders finish, one coordinator closes output, and no sender closes it independently. |
| 34 | `run(workers int, jobs []int) []int` | Zero workers and empty jobs have explicit, non-deadlocking behavior. |
| 35 | `fanOut(stop <-chan struct{}, jobs <-chan int, workers int) <-chan int` | Cancellation stops workers and closes output after all sends finish. |
| 36 | `mergeResults(inputs ...<-chan result) <-chan result` | Result envelopes preserve value/status and output closes for empty inputs. |
| 37 | `collectOrStop(stop <-chan struct{}, work <-chan int) <-chan int` | A caller that stops receiving can close stop and allow the producer to finish. |
| 38 | `watch(done <-chan struct{}) <-chan string` | Multiple observers can receive the same close-as-broadcast completion event. |
| 39 | `type command int` with `pause` and `resume`; `serveCommands(stop <-chan struct{}, commands <-chan command, jobs <-chan int) <-chan int` | Data and control cases are selected without starving shutdown. |
| 40 | `relay(stop <-chan struct{}, in <-chan int) <-chan int` | Both receive and send paths are cancellable under a blocked downstream. |

- [ ] **Step 1: Write tests for first-result cancellation and request/reply.**

Use task functions that block on a per-test gate or stop channel, not sleeps. Assert the returned result, each reply value, server completion, and that canceled tasks can exit.

- [ ] **Step 2: Implement exercises 31–34.**

Use a capacity-one winner/error channel so the winning worker cannot block after the caller returns; use a coordinator for source closure; define and document explicit zero-worker behavior instead of allowing a deadlock.

- [ ] **Step 3: Write and run red tests for exercises 35–40.**

For every output stream, include a normal close test and an early-stop test. Exercise 39 must send a stop command while jobs are blocked and require output closure.

- [ ] **Step 4: Implement exercises 35–40.**

Keep output ownership in one goroutine, use `select` around each potentially blocking operation, and set closed input variables to nil only in the exercise that explicitly teaches that technique later.

- [ ] **Step 5: Verify and commit.**

Run focused tests twice, race only if the final implementation uses shared mutable state listed in `race.list`, and commit:

```sh
git add exercises/13_channels solutions/13_channels
git commit -m "feat: add channel service protocol practice"
```

---

### Task 7: Implement channel control, limits, and capstones 41–50

**Files:**
- Create: `exercises/13_channels/channels41..50/{main.go,main_test.go}`
- Create: `solutions/13_channels/channels41..50/main.go`
- Modify: `race.list` only for exercises with shared mutable state in their tests or solutions.

**Interfaces and test contracts:**

| Exercise | Learner interface | Required assertions |
| --- | --- | --- |
| 41 | `drain(first, second <-chan int) []int` | Closed inputs are set to nil after draining; no repeated zero values. |
| 42 | `produce(stop <-chan struct{}) <-chan int` | Producer exits when consumer abandons output and stop closes. |
| 43 | `parallel(limit int, jobs []int, work func(int) int) []int` | Buffered token capacity limits active work and output order is restored. |
| 44 | `parallel(stop <-chan struct{}, limit int, jobs []int, work func(int) int) ([]int, bool)` | Cancellation can interrupt token acquisition and active result publication. |
| 45 | `rateLimit(tokens <-chan struct{}, in <-chan int) <-chan int` | Each input consumes one externally supplied token; input closure closes output. |
| 46 | `rateLimit(stop <-chan struct{}, tokens <-chan struct{}, in <-chan int) <-chan int` | Stop interrupts both token receive and output send. |
| 47 | `serve(stop <-chan struct{}, jobs <-chan request) (<-chan response, <-chan struct{})` | Shutdown stops new work, waits workers, closes results, then closes done. |
| 48 | `runOrderedBounded(stop <-chan struct{}, workers int, jobs []int) ([]int, bool)` | Bounded workers, ordered output, cancellation, and closure all hold. |
| 49 | `runFirstErrorBounded(stop <-chan struct{}, workers int, jobs []job) ([]int, error)` | One error cancels production, waits workers, and returns without a double close. |
| 50 | `runService(stop <-chan struct{}, workers int, requests []request) ([]response, error)` | Request/reply, bounded workers, backpressure, timeout signal, cancellation, order, error, and close ownership are all observable. |

- [ ] **Step 1: Write failing tests for nil-channel and producer-abandonment cases.**

Exercise 41 must include two already-closed inputs and verify exactly four values. Exercise 42 must stop without draining output and require output closure within 500ms.

- [ ] **Step 2: Implement and verify exercises 41–46.**

Use nil assignment only after comma-ok reports closure. Use buffered token channels as semaphores and receive tokens from the caller rather than constructing a ticker. Stop cases must surround token receives and output sends.

- [ ] **Step 3: Write capstone tests before capstone solutions.**

For exercises 47–50, tests must cover: normal completion, empty requests, a blocked downstream, cancellation before all jobs start, one failing job, result closure, done closure, and stable ordered responses. Use deterministic gates and injected functions instead of sleeps.

- [ ] **Step 4: Run the red phase for exercises 47–50.**

Run the starter verifier and record a failure for every capstone. A capstone that merely returns an empty slice or nil error must fail because tests require per-request responses and lifecycle closure.

- [ ] **Step 5: Implement the minimum capstone solutions.**

Use one owner for every close, a capacity-one error/winner channel where a caller may stop receiving, explicit worker join before result close, and indexed result storage for ordering. Do not use `context`, ticker construction, or shared state without a listed race reason.

- [ ] **Step 6: Verify and commit.**

Run focused normal tests twice, selective race tests from `race.list`, `go vet`, and starter verification on a pristine copy. Commit:

```sh
git add exercises/13_channels solutions/13_channels race.list
git commit -m "feat: complete advanced channel practice"
```

---

### Task 8: Rewrite foundation chapter guides and root index

**Files:**
- Create or replace: `exercises/12_goroutines/README.md`
- Create or replace: `exercises/13_channels/README.md`
- Modify: `README.md`

- [ ] **Step 1: Document the 10 goroutine progression.**

List every exercise in order, state that WaitGroup is the completion primitive here, and explain that channel protocols begin in the next chapter.

- [ ] **Step 2: Document the 50 channel progression.**

Include the ownership model, repeated-skill matrix, difficulty bands, exact 1–50 table, cancellation diagrams, strict hint expectations, and the handoff of raw pipeline work to `24_concurrency_patterns`.

- [ ] **Step 3: Document checker behavior.**

Explain that an exercise target with a remaining `TODO:` is rejected, `--race` uses `race.list`, and `--race-all` is the expensive full audit. Show commands for one exercise, all solutions, and the starter verifier.

- [ ] **Step 4: Update root counts and links.**

Change `12_goroutines` to 10 and `13_channels` to 50, update the solution count text, and link the two chapter guides.

- [ ] **Step 5: Verify docs and commit.**

Run `git diff --check`, `gofmt -d $(git ls-files '*.go')`, and commit:

```sh
git add exercises/12_goroutines/README.md exercises/13_channels/README.md README.md
git commit -m "docs: guide expanded goroutine and channel practice"
```

---

### Task 9: Foundation verification checkpoint

**Files:**
- Verify: `check.sh`, `race.list`, `scripts/verify_exercise_starters.sh`, `exercises/12_goroutines`, `exercises/13_channels`, `solutions/12_goroutines`, `solutions/13_channels`.

- [ ] **Step 1: Verify layout parity.**

Run:

```sh
exercise_dirs="$(find exercises -mindepth 2 -maxdepth 2 -type d | sort)"
solution_dirs="$(find solutions -mindepth 2 -maxdepth 2 -type d | sed 's#^solutions/#exercises/#' | sort)"
diff -u <(printf '%s\n' "$exercise_dirs") <(printf '%s\n' "$solution_dirs")
```

- [ ] **Step 2: Verify the pristine starter gate.**

Run `sh scripts/verify_exercise_starters.sh exercises/12_goroutines exercises/13_channels`. Expected: every starter is rejected or fails its focused test before any solution code is introduced.

- [ ] **Step 3: Run the normal solution suite.**

Run `go test ./solutions/12_goroutines/... ./solutions/13_channels/...`, `go vet ./solutions/12_goroutines/... ./solutions/13_channels/...`, and `sh check.sh solutions/12_goroutines --run-all` plus the matching channels command.

- [ ] **Step 4: Run selective and full race checks.**

Run `sh check.sh solutions/12_goroutines --run-all --race`, `sh check.sh solutions/13_channels --run-all --race`, and only after the fast checks pass run `sh check.sh solutions/12_goroutines --run-all --race-all` and the matching channel command.

- [ ] **Step 5: Review the diff.**

Confirm no raw pipeline exercise remains in `13_channels`, every hint names the blocking and closing obligations, no new starter has a default path that can pass after the TODO gate is removed, and no files under `working-progress` changed.
