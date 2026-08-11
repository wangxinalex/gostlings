# Task 7 Implementer Report — Advanced Raw Channel Practice 41–50

## Delivered

Added ten learner exercises, deterministic learner tests, and matching raw-channel solutions under `13_channels`.

| Exercise | Interface | Covered behavior |
| --- | --- | --- |
| 41 | `drain` | Nil-disables closed selected inputs after their buffered values drain. |
| 42 | `produce` | Stop releases an abandoned output send and closes the stream. |
| 43 | `parallel` | Buffered-token semaphore bounds active work and indexed results restore order. |
| 44 | cancellable `parallel` | Stop covers token acquisition and result publication. |
| 45 | `rateLimit` | Caller-supplied tokens gate one forwarded input each. |
| 46 | cancellable `rateLimit` | Stop interrupts token waits and blocked output sends. |
| 47 | `serve` | Per-request replies, result closure, and done-after-result worker join. |
| 48 | `runOrderedBounded` | Bounded workers, stable order, empty input, and cancellation. |
| 49 | `runFirstErrorBounded` | First job error stops production and waits active workers. |
| 50 | `runService` | Bounded request/reply service with stop-as-timeout signal, ordering, backpressure, and first-error propagation. |

Exercises 47 and 50 document the package-private request/response fields supplied by the task owner:

```go
type response struct { value int; err error }
type request struct { value int; reply chan response; err error }
type job struct { value int; err error }
```

Every starter retains a `TODO:` seam and an actionable hint. Solutions use raw channels only: no `context.Context`, sync primitives, atomics, timer/ticker construction, or later integrated patterns.

## Test design

Tests use gate channels, injected processing functions, lifecycle closures, and individual 500ms liveness guards. They do not use sleeps or package-timeout-only assertions. Coverage rejects nil/no-op implementations, repeated closed-channel zero values, ignored tokens or cancellation, excess active work, unordered output, swallowed job/request errors, duplicate replies, and closing outputs before worker exit.

The 47–50 capstone cases cover normal and empty inputs, blocked downstream/reply sends, cancellation before later work begins, one failing job/request, result and done closure, bounded worker starts, and stable ordered responses.

## Verification

All commands used Go 1.26.5 with:

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH
GOCACHE=/private/tmp/gostlings-go-cache
```

Fresh solution-backed learner tests passed twice through a temporary, deleted Go overlay mapping each learner `main.go` to its matching solution:

```sh
go test -count=2 -overlay=.solution-test-overlay.json ./exercises/13_channels/channels41 ... ./exercises/13_channels/channels50
```

Also passed:

```sh
go test -count=1 ./solutions/13_channels/channels41 ... ./solutions/13_channels/channels50
go vet ./exercises/13_channels/channels41 ... ./channels50 ./solutions/13_channels/channels41 ... ./channels50
sh scripts/verify_exercise_starters.sh exercises/13_channels/channels41 ... channels50
git diff --check
gofmt -d exercises/13_channels/channels41 ... channels50 solutions/13_channels/channels41 ... channels50
```

The normal checker rejects an untouched starter because its `TODO:` remains; the starter verifier additionally runs every starter's focused behavioral tests and confirms each fails rather than accidentally passing.

No new `race.list` entry or `-race` run was appropriate: the solutions use channel ownership rather than shared mutable state. Test hook variables are assigned before launch and restored only after the corresponding result/done/return confirms all workers exited, so they are not concurrently accessed mutably.

## Post-commit self-review

Commit: `6cbaf9c feat: complete advanced channel practice`

- A single producer or coordinator owns every shared output close.
- Worker exit acknowledgements precede all shared-result/done closure.
- Stop cases guard token waits, jobs receives, replies, and result sends.
- The first-error pools use capacity-one error channels and a coordinator-only internal cancel close, avoiding blocked error senders and double closes.
- Results that promise ordering carry and restore request/job indexes.
- The intentional Task 3 and Task 5 report deletions were not staged or restored.

No Task 7 concerns remain.

## Review fixes

- Replaced the `100ms` absence windows in channels43, 48, and 50 with
  permit-backed process gates. The first two calls consume fixed active permits;
  any third entry before release reports on an explicit overflow channel. Tests
  release the permitted calls, join normally, and then reject an observed
  overflow without using a timeout to establish absence.
- Strengthened channels47 shutdown coverage with a package-private
  `serveWorkerCount` test seam. A gated first worker holds one request active
  while a second request remains queued; after stop, the queued request has
  neither a reply nor a published result. The exported `serve` signature and
  channel-close ownership are unchanged.
- Added channels49 external-stop coverage: stop cannot return while a gated
  active call remains blocked, and after release it joins with no successful
  value or job error. Added a separate two-worker/three-successful-job
  saturation case using the same explicit overflow observation. The existing
  error-join test now dispatches its gated success before the failing job so its
  intended concurrent state does not depend on worker scheduling.

Fresh Go 1.26.5 checks with `GOCACHE=/private/tmp/gostlings-go-cache` passed:

```sh
go test -count=2 -overlay=.solution-test-overlay.json \
  ./exercises/13_channels/channels43 ./exercises/13_channels/channels47 \
  ./exercises/13_channels/channels48 ./exercises/13_channels/channels49 \
  ./exercises/13_channels/channels50
go test -race -count=1 -overlay=.solution-test-overlay.json \
  ./exercises/13_channels/channels43 ./exercises/13_channels/channels47 \
  ./exercises/13_channels/channels48 ./exercises/13_channels/channels49 \
  ./exercises/13_channels/channels50
```

The Task 7 starter verifier, `go vet` across learner and solution packages
41–50, `gofmt -d` for the changed Go files, and `git diff --check` also passed.
The overlay was removed after verification. The pre-existing Task 3 and Task 5
report deletions remain unstaged and untouched.
