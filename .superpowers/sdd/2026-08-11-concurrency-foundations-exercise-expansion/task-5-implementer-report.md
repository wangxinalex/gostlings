# Task 5 Implementer Report — Raw Channel Composition 21–30

## Commit

`90947d3 feat: add channel fan-in fan-out and worker pools`

## Scope

Replaced the old pre-expansion channel content at 21–24 and added channels 25–30. Each exercise has a TODO-bearing learner starter, focused behavioral tests, a matching solution, and an explicit beginner-facing hint. The ordered interfaces are:

| Exercise | Interface | Tested behavior |
| --- | --- | --- |
| 21 | `squareWorkers(workers int, jobs <-chan int) <-chan int` | Fan-out squares every job, starts every requested worker, and a coordinator closes output after all worker exits. |
| 22 | `merge(inputs ...<-chan int) <-chan int` | Fan-in forwards every input, closes for zero inputs, and does not ignore an initially blocked input. |
| 23 | `merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int` | Fan-in cancels both blocked input receive and blocked output send paths. |
| 24 | `squareWorkers(stop <-chan struct{}, workers int, jobs <-chan int) <-chan int` | A cancellable worker pool covers normal results plus blocked jobs receive and blocked results send. |
| 25 | `merge(inputs ...<-chan int) <-chan int` | Empty, already-closed, buffered, and still-open non-nil inputs terminate correctly. |
| 26 | `run(workers int, jobs []int) []int` | Basic pool processes every job, starts each requested worker, and returns for empty input. |
| 27 | `runOrdered(workers int, jobs []int) []int` | Indexed worker results restore original input order and empty input returns. |
| 28 | `run(workers int, jobs []job) error` | The first observed error is returned; stop closes once; every worker exits before return. |
| 29 | `startPool(workers int) (chan<- int, <-chan int)` | Directional jobs/results handles, full result stream, and coordinator-owned result close. |
| 30 | `runBounded(workers, buffer int, jobs []int) []int` | Bounded queue capacity, requested worker count, full result set, and empty input. |

## Implementation boundaries

- Solutions use raw data, stop, exit-acknowledgement, failure, and result channels only.
- Every multi-sender output has a separate coordinator which waits for raw worker/forwarder acknowledgements before it calls `close`.
- Cancellable merge and worker-pool operations select on `stop` beside both potentially blocking receives and sends.
- The basic, ordered, directional, and bounded pools each use a jobs producer that closes its jobs channel, workers that send results, and a coordinator that closes results after all worker acknowledgements.
- Exercise 27 sends indexed envelopes and writes returned values to the matching result slot. Exercise 30 creates `make(chan int, buffer)` after normalizing a negative capacity to zero.
- No task starter or solution imports `context`, `sync`, atomics, `time.Timer`, or `time.Ticker`; the only time usage is in tests as 500ms liveness watchdogs.

## Test design and red phase

Focused tests were written before the new implementations. The first run failed at the intended undefined replacement interfaces for all ten packages, including `squareWorkers`, both `merge` variants, each pool API, and the `job` error-pool seam.

The tests use channel gates rather than sleeps:

- Worker-start hooks require the requested number of workers, rejecting a single-worker implementation.
- Worker-exit hooks hold acknowledgement publication so tests can reject a worker that closes shared output early.
- Blocked-input tests wait for an explicit producer-send acknowledgement and then require the forwarded value, rejecting a merge that skips a source.
- Cancellable variants use a pre-send hook, then close `stop` while no output receiver exists; this deterministically tests cancellation of a blocked send.
- The error-pool test blocks both worker exit hooks until it has observed the one stop-close event, then rejects a return before the exit gates release.
- The bounded-pool test observes the real jobs-channel capacity and gates two workers, rejecting a queue capacity other than the requested bound and a single-worker implementation.

Every untouched starter retained `TODO:` and was rejected by the normal checker. The starter verifier additionally ran the focused tests with the TODO gate bypassed and confirmed that each starter fails behaviorally.

## Verification

All Go commands used:

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH
GOCACHE=/private/tmp/gostlings-go-cache
```

Successful solution-backed learner verification (using a temporary Go overlay mapping each learner `main.go` to the matching solution) was:

```sh
go test -count=1 -overlay=.solution-test-overlay.json \
  ./exercises/13_channels/channels21 ... ./exercises/13_channels/channels30
go test -count=3 -race -overlay=.solution-test-overlay.json \
  ./exercises/13_channels/channels21 ... ./exercises/13_channels/channels30
```

All ten focused packages passed in the normal run and all three repeated race runs. The overlay was deleted before commit.

Additional successful checks:

```sh
go test -count=1 ./solutions/13_channels/...
go vet ./solutions/13_channels/...
sh scripts/verify_exercise_starters.sh exercises/13_channels
# Normal checker was also required to reject each of channels21 through channels30.
rg -n 'context\.Context|sync\.|sync/|atomic\.|time\.(NewTimer|NewTicker)|time\.(Timer|Ticker)' \
  exercises/13_channels/channels{21,22,23,24,25,26,27,28,29,30}/main.go \
  solutions/13_channels/channels{21,22,23,24,25,26,27,28,29,30}/main.go
git diff --check
```

The forbidden-API scan produced no matches; all verification commands exited successfully.

## Self-review

- Confirmed every interface and sequence matches the Task 5 table, including the intentional replacement of obsolete 21–24 material.
- Confirmed every starter has a concrete TODO seam and an actionable hint that names the blocking and close-owner responsibilities.
- Confirmed every normal multi-sender output is closed by a single coordinator after raw exit acknowledgements; no worker or forwarder closes a shared output.
- Confirmed cancellation covers both directions for exercises 23 and 24.
- Confirmed exercises 21, 24, 26, 27, 29, and 30 reject single-worker implementations; exercise 27 rejects completion-order output; exercise 28 rejects swallowed errors and premature returns; exercise 30 checks queue capacity.
- Confirmed the repeated race run covers merge and all pool behavior, including the synchronization test hooks.

## Concerns

None.

## Review Fixes — 2026-08-11

### Scope

- Added package-private `onForwarderExit` hooks to channels22 and channels25 starters and solutions. Each forwarder invokes the hook immediately before its raw exit acknowledgement. The new lifecycle tests block every forwarder at that point, prove the merged output is still open, then release and drain it.
- Added a package-private `onWorkerExit` hook to channels26 starters and solution. The solution invokes it before publishing the worker exit acknowledgement. The test blocks all three requested worker exits, proves `run` cannot return, then releases and checks all squared results.
- Reworked channels27's requested-worker test with per-value gates. It waits for all input values to start, releases `3`, `1`, then `4`, confirms that deliberate non-input completion sequence, and still requires input-order output.
- Added the channels28 one-worker regression where a successful job comes before an error job, proving the later error is inspected while retaining the existing first-error, stop-once, and join coverage.
- Added a package-private `onBoundedProcessStart(value int)` hook to channels30 starters and solution, called immediately after receipt from the bounded jobs channel. The bounded-pool test blocks the first two actual processing calls, observes both process-entry hooks, rejects a third entry before release, and separately checks capacity is exactly one.

All new test-hook assignments use `t.Cleanup` to restore the previous hook. The affected implementations remain raw-channel-only and retain their public exercise APIs.

### Verification

Commands used Go 1.26.5 with `GOCACHE=/private/tmp/gostlings-go-cache`:

```sh
go test -count=1 -overlay=.solution-test-overlay.json \
  ./exercises/13_channels/channels22 \
  ./exercises/13_channels/channels25 \
  ./exercises/13_channels/channels26 \
  ./exercises/13_channels/channels27 \
  ./exercises/13_channels/channels28 \
  ./exercises/13_channels/channels30
go test -count=3 -overlay=.solution-test-overlay.json [same six packages]
go test -count=3 -race -overlay=.solution-test-overlay.json [same six packages]
sh scripts/verify_exercise_starters.sh exercises/13_channels
go vet ./solutions/13_channels/...
git diff --check
```

All commands succeeded. The temporary solution overlay was removed before commit. The forbidden-API scan found no use of `context`, `sync`, atomics, `time.Timer`, or `time.Ticker` in channels21–30 starters or solutions.

### Concerns

None.
