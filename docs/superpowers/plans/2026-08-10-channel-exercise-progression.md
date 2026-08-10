# Progressive Channel Exercises Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task with verification checkpoints.

**Goal:** Replace the current eight channel exercises with a sequential, frequency-weighted 24-exercise curriculum and publish the matching solutions and guide from a branch based on `origin/main`.

**Architecture:** Each exercise is an independent `main` package with a learner starter, deterministic tests, and a runnable solution. The chapter guide documents the dependency path and lifecycle rules; the root README only provides the chapter index. Existing channel directories are recreated in their new positions so the numeric order is the learning order.

**Tech Stack:** Go 1.26, standard library only, `sync`, `time`, `errors`, `fmt`, `sort`, and `sync/atomic` only where tests need concurrency observation.

## Global Constraints

- Modify only the main repository lineage in `/private/tmp/gostlings-channels`; do not modify `/Users/wangxinalex/SelfStudy/Rust/gostlings/working-progress`.
- Keep `exercises/13_channels/channelsN` and `solutions/13_channels/channelsN` in exact directory parity for `N=1..24`.
- Every solution package must compile, pass `go vet`, run through `check.sh`, and pass the race detector.
- Learner tests must use explicit channel completion or bounded timeouts, never `time.Sleep` as the only synchronization assertion.
- Preserve the repository's Go 1.26 requirement and use `sync.WaitGroup.Go` where it improves readability; explain the equivalent `Add`/`Done` lifecycle in the guide where relevant.
- Use `apply_patch` for source edits and commit after each coherent curriculum batch.

## File map

- Create or replace `exercises/13_channels/channels1..24/main.go` and `main_test.go`: learner starters and behavioral tests.
- Create or replace `solutions/13_channels/channels1..24/main.go`: runnable reference implementations.
- Create `exercises/13_channels/README.md`: chapter model, 24-exercise map, dependencies, timings, pitfalls, and production notes.
- Modify `README.md`: update the channel chapter count and link to its guide.
- Create `docs/superpowers/specs/2026-08-10-channel-exercise-progression-design.md`: already committed as `20e3b38`; it is the source design for this plan.

### Task 1: Rebuild foundational semantics exercises 1–5

**Files:**
- Replace: `exercises/13_channels/channels1..5/main.go`
- Replace: `exercises/13_channels/channels1..5/main_test.go`
- Replace: `solutions/13_channels/channels1..5/main.go`

**Interfaces and tests:**

- `channels1` keeps a `main()` output test for `hi`; starter performs the send synchronously, while the solution moves it into a goroutine.
- `channels2` keeps a `main()` output test for `1\n2\n`; starter uses an unbuffered channel, while the solution gives it capacity two.
- `channels3` keeps a `main()` output test for `1\n2\n3\n`; starter sends values but does not close the channel, while the solution closes after the last send.
- `channels4` defines `func read(ch <-chan int) (int, bool)`; tests require a closed channel to return `(0, false)` and a buffered zero value to return `(0, true)`, proving that a closed receive is not a real zero value.
- `channels5` defines `func generate(values ...int) <-chan int`; tests verify all values, output closure, and an empty input returning promptly. The solution owns and closes the output.

- [ ] Write all five starter files and tests with the signatures above.
- [ ] Write runnable solutions for all five packages.
- [ ] Run `gofmt -w exercises/13_channels solutions/13_channels` and focused tests for `channels1..5`.
- [ ] Run `go test -run '^$' ./exercises/13_channels/...` to compile every learner starter without executing intentionally incomplete tests.
- [ ] Run `go test -race ./solutions/13_channels/...` to verify every completed reference package.
- [ ] Commit `feat: rebuild foundational channel exercises`.

### Task 2: Add select, timeout, completion, and cancellation exercises 6–12

**Files:**
- Create or replace `exercises/13_channels/channels6..12/main.go`
- Create `exercises/13_channels/channels6..12/main_test.go`
- Create or replace `solutions/13_channels/channels6..12/main.go`

**Interfaces and tests:**

- `channels6`: `func receiveFast(fast, slow <-chan string) string`; fast is buffered and slow has no sender; test expects the fast value.
- `channels7`: `func await(ch <-chan string) string`; an empty channel must return `timed out` within a bounded test timeout.
- `channels8`: `func tryReceive(ch <-chan int) string`; test covers an empty channel returning `no value` and a ready channel returning its value.
- `channels9`: `func complete() <-chan struct{}`; test requires the returned done channel to close promptly.
- `channels10`: `func produce(stop <-chan struct{}) <-chan int`; test closes `stop` without receiving and requires the output to close, then separately checks normal values.
- `channels11`: `func run() string`; test expects `timed out` quickly and exercises the producer's cancellation and completion wait.
- `channels12`: `func startWorkers(count int, stop <-chan struct{}) <-chan struct{}`; test starts three workers, closes one stop channel, and requires the group-done channel to close.

- [ ] Add starter implementations that compile but fail their lifecycle assertions without revealing the complete pattern.
- [ ] Add deterministic tests with completion channels and maximum 500ms deadlock guards.
- [ ] Implement solutions using `select`, `close(done)`, cancellation-aware sends, and a `WaitGroup` coordinator.
- [ ] Compile learner packages with `go test -run '^$'` and run normal plus race checks for solutions 6–12.
- [ ] Commit `feat: add channel lifecycle and cancellation exercises`.

### Task 3: Implement fan-out, fan-in, and worker-pool exercises 13–18

**Files:**
- Create or replace `exercises/13_channels/channels13..18/main.go`
- Create `exercises/13_channels/channels13..18/main_test.go`
- Create or replace `solutions/13_channels/channels13..18/main.go`

**Interfaces and tests:**

- `channels13`: `func squareWorkers(workers int, jobs <-chan int) <-chan int`; close jobs, collect all squared values, and require results to close after workers exit.
- `channels14`: `func merge(inputs ...<-chan int) <-chan int`; verify all values and output closure, including zero inputs.
- `channels15`: `func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int`; verify values in the normal path and prompt output closure when stop closes while an input never sends.
- `channels16`: `func run(workers int, jobs []int) []int`; verify one square per job and prompt completion for empty jobs.
- `channels17`: `func runOrdered(workers int, jobs []int) []int`; verify results remain in input order despite parallel workers.
- `channels18`: `type job struct { value int; fail bool }` and `func run(workers int, jobs []job) error`; verify a failing job returns an error, all workers exit, and no second close of the cancellation channel occurs.

- [ ] Write tests first for output closure, empty input, cancellation, order, and first-error behavior.
- [ ] Implement fan-out with one output closer, fan-in with one `WaitGroup` closer, and cancellable forwarding on both receive and send.
- [ ] Implement worker-pool order restoration by carrying an index and worker-pool error cancellation with `sync.Once` or a single coordinator-owned close.
- [ ] Compile learner packages with `go test -run '^$'`, run solutions 13–18 under `-race`, and run each new solution with `check.sh`.
- [ ] Commit `feat: add channel composition exercises`.

### Task 4: Implement pipeline and bounded-concurrency exercises 19–24

**Files:**
- Create `exercises/13_channels/channels19..24/main.go`
- Create `exercises/13_channels/channels19..24/main_test.go`
- Create `solutions/13_channels/channels19..24/main.go`

**Interfaces and tests:**

- `channels19`: `func square(in <-chan int) <-chan int`; verify transformed values and output closure.
- `channels20`: `func transform(in <-chan int, fn func(int) int) <-chan int` and `func pipeline(in <-chan int) <-chan int`; compose two stages and verify final values and closure.
- `channels21`: `func pipeline(stop <-chan struct{}, in <-chan int) <-chan int`; verify normal composition and prompt output closure when input blocks and stop closes.
- `channels22`: `func parallel(limit int, jobs []int, work func(int) int) []int`; test results and observe active work with a mutex, requiring the maximum active count not to exceed `limit`.
- `channels23`: `func rateLimit(ticks <-chan time.Time, in <-chan int) <-chan int`; feed explicit ticks and jobs so tests do not sleep, and require one output per ticked job followed by closure.
- `channels24`: `func drain(first, second <-chan int) []int`; close each input after buffered values and verify both streams are drained without repeated zero values.

- [ ] Add tests that make pipeline and bounded-concurrency lifecycle failures observable without goroutine-count heuristics.
- [ ] Implement close propagation for each pipeline stage, cancellation on both receive and send, semaphore tokens via a buffered channel, ticker gating, and nil assignment after a selected input closes.
- [ ] Compile learner packages with `go test -run '^$'`, run solutions 19–24 with `-race`, and inspect solution lifecycles with `go test -timeout 5s`.
- [ ] Commit `feat: add pipeline and bounded concurrency exercises`.

### Task 5: Rewrite chapter documentation and root index

**Files:**
- Create: `exercises/13_channels/README.md`
- Modify: `README.md`

- [ ] Document the sender/receiver/closer ownership rule and the difference between data, completion, cancellation, and result channels.
- [ ] Add the exact 24-row exercise map with frequency and difficulty bands, plus “what the previous band prepared” guidance.
- [ ] Explain the timing of exercises 9–12, 14–18, and 19–21 in prose and show why close order matters.
- [ ] Add practical notes for `context.Context`, `WaitGroup.Go`, race detection, and goroutine-leak diagnosis.
- [ ] Replace the root README channel entry with count 24 and a link to `exercises/13_channels/README.md`.
- [ ] Commit `docs: guide progressive channel practice`.

### Task 6: Run repository-level verification and prepare the PR

**Files:**
- Verify all changed files; no additional source changes expected.

- [ ] Run `git diff --check` and `gofmt -d $(git ls-files '*.go')`.
- [ ] Run `go test ./solutions/...` and `go vet ./solutions/...` with Go 1.26.5.
- [ ] Run `sh check.sh solutions --run-all` and `sh check.sh solutions --run-all --race`.
- [ ] Run `go test -run '^$' ./exercises/13_channels/...` and `go test -race ./solutions/13_channels/...`.
- [ ] Verify directory parity with the CI `find`/`diff` command.
- [ ] Review the complete diff for accidental changes outside channel exercises, docs, and root README.
- [ ] Push `agent/expand-channel-practice` and create a draft PR with a summary, testing evidence, and note that `working-progress` was left untouched.
