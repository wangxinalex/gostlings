# Task 3 Implementer Report — Channel Fundamentals 1–10

## Scope completed

Rebuilt the first graded channel block as ten raw-channel exercises and matching solutions. The progression is now:

1. unbuffered handoff;
2. a capacity-two buffered channel;
3. producer close plus `range`;
4. comma-ok receiving;
5. draining buffered values after close with `drainClosed`;
6. a producer-owned generator;
7. the same generator pattern with receive-only ownership exposed through `generate(values ...int) <-chan int` and `receiveAll(ch <-chan int)`;
8. selecting a ready input without waiting for a silent input;
9. nonblocking receive; and
10. nonblocking send.

The requested moves were applied: former channels 5, 6, and 8 are now 6, 8, and 9 respectively. The previous timeout, done-notification, and cancellation content at 7, 9, and 10 was displaced because Task 4 replaces exercises 11–20 with that material.

## Learner starters and tests

- Every starter in channels 1–10 retains an explicit `TODO:` seam.
- Existing 1–4 behavior was retained.
- New exercise 5 tests that `drainClosed` preserves sent values, including a buffered zero, before it observes closure.
- New exercise 7 uses `receiveAll(ch <-chan int)` as the directional helper and verifies ordered values plus closure for both populated and empty generators.
- New exercise 10 verifies successful buffered delivery, refusal of a full buffer without overwriting it, and immediate refusal for an unbuffered channel with no receiver.
- Exercise 9’s empty-channel assertion now uses a completion channel and a bounded watchdog, so an accidentally blocking receive is rejected deterministically rather than relying on the package-wide test timeout.
- Hints give concrete next actions (`value, ok := <-ch`, `defer close(out)`, and `select` with `default`) while keeping the implementation work with the learner.

The production starters and solutions for 1–10 use only raw channel operations, channel close ownership, directional channel types, and `select` with `default`. No `context`, sync primitive, timer/ticker, or integrated concurrency pattern was added.

## Red-phase evidence

Before adding the three new starters, their focused test packages failed at the intended missing seams:

- channels5: `drainClosed` undefined;
- channels7: `generate` and `receiveAll` undefined;
- channels10: `trySend` undefined.

After adding the TODO-bearing starters, this completed successfully:

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH \
GOCACHE=/private/tmp/gostlings-go-cache \
sh scripts/verify_exercise_starters.sh exercises/13_channels
```

That verifier requires every starter to retain `TODO:` and requires every focused starter test to fail when the normal TODO gate is bypassed. Normal `check.sh` calls for channels 1, 5, and 10 also rejected the starters before test execution.

## Green-phase and static verification

All checks used Go 1.26.5 at the requested PATH and `GOCACHE=/private/tmp/gostlings-go-cache`.

The learner tests were run in a disposable clean copy of the current starters with each matching solution overlaid. All ten packages passed with `go test -count=1`.

```sh
go test -count=1 \
  ./exercises/13_channels/channels1 ... ./exercises/13_channels/channels10
go test -count=1 \
  ./solutions/13_channels/channels1 ... ./solutions/13_channels/channels10
go vet ./solutions/13_channels/...
sh scripts/verify_exercise_starters.sh exercises/13_channels
git diff --check
```

The solution package command compiled all ten packages successfully, `go vet` completed without diagnostics across all channel solutions, and the final starter verifier completed successfully.

## Self-review

- Confirmed exact function signatures: `read(ch <-chan int) (int, bool)`, `drainClosed(ch <-chan int) []int`, both generator signatures, `receiveFast(fast, slow <-chan string) string`, `tryReceive(ch <-chan int) string`, and `trySend(ch chan<- int, value int) bool`.
- Confirmed every 1–10 exercise and solution directory has the expected source/test parity.
- Confirmed all first-ten starters contain `TODO:` and every solution omits them.
- Checked production source for forbidden `context`, `sync`, timer, and ticker APIs; none are present in the first block.
- Checked formatting and whitespace with `gofmt` and `git diff --check`.
- Reviewed each nonblocking test’s negative case: an empty receive, full send, and unbuffered send without receiver cannot pass through a no-op or a blocking channel operation.

## Commit

`4bf8227 feat: strengthen channel semantics and select basics`

## Concerns

None.

---

## Review fix pass — channels 1, 6, and 8

Applied the focused Task 3 review fixes without renumbering exercises or changing learner-facing APIs.

- `channels1` now parses `main.go` with Go's AST and requires the `chan string` allocation to be unbuffered and its send to occur inside a goroutine. The existing output assertion remains in place.
- `channels6` now invokes `generate` through a buffered result channel from a goroutine and establishes a 500 ms watchdog before waiting for the returned output. A synchronous-send implementation therefore fails locally instead of consuming the package timeout. Both populated and empty-input tests use this seam.
- `channels8` now runs both ready-input cases through a result goroutine with a 200 ms watchdog. The fast-ready timeout closes the otherwise silent slow channel, matching the slow-ready cleanup and preventing a leaked blocked goroutine.

No exercise or solution production source changed. The raw-channel-only constraints remain intact; the bounded watchdogs are test-harness-only.

### Verification

Using Go 1.26.5 with `GOCACHE=/private/tmp/gostlings-go-cache`:

```sh
go test -count=1 -timeout 5s \
  ./solutions/13_channels/channels1 \
  ./solutions/13_channels/channels6 \
  ./solutions/13_channels/channels8
go vet ./solutions/13_channels/...
sh scripts/verify_exercise_starters.sh exercises/13_channels
git diff --check
```

The focused solution tests were run with the revised exercise harnesses overlaid in a disposable clean clone; all three passed. The starter verifier, vet, and whitespace checks also passed.

### Re-review fix — channels6 watchdog starts before producer launch

Moved `deadline := time.After(500 * time.Millisecond)` before the `go func()` invocation in `generateWithWatchdog`. The watchdog now measures from before the producer goroutine is launched, while preserving the existing timeout, channel, and test semantics.

Verification was run with Go 1.26.5 and `GOCACHE=/private/tmp/gostlings-go-cache`. The learner starter retains its intentional TODO and therefore fails its focused tests; in a disposable copy, the matching solution was overlaid onto the learner package and passed:

```sh
go test -count=1 -timeout 5s ./solutions/13_channels/channels6
go test -count=1 -timeout 5s ./exercises/13_channels/channels6  # with solutions/13_channels/channels6/main.go overlaid
```
