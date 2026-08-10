# Channel Exercise Progression Design

## Goal

Rebuild `exercises/13_channels` as a sequential 24-exercise curriculum that teaches Go channels in the order developers usually encounter the ideas: basic blocking and ownership first, then lifecycle and cancellation, then reusable composition patterns, and finally bounded concurrency and less-common `select` techniques.

The chapter will be implemented in the main repository worktree lineage and published as a pull request. The user's separate `working-progress` worktree is intentionally out of scope and must not be modified.

## Design principles

1. **Real-world frequency drives the sequence.** Closing and ranging, `select`, timeouts, cancellation, fan-out/fan-in, worker pools, pipelines, and bounded concurrency receive most of the exercises. Nil-channel switching is covered once as an advanced technique.
2. **Every exercise builds on earlier vocabulary.** A later exercise may combine earlier patterns, but it will not require an unexplained concept from a later chapter.
3. **Important patterns repeat with new failure modes.** Closing/lifecycle appears in producers, fan-in, worker pools, and pipelines. Cancellation appears first as a single stoppable sender, then as broadcast cancellation, cancellable fan-in, and cancellable pipelines.
4. **Tests prove lifecycle, not just values.** Tests must check that output channels close, calls return for empty input, and cancellable goroutines can finish. Tests may use bounded timeouts as deadlock guards but must not depend on a lucky scheduling order.
5. **Exercises and solutions mirror exactly.** Every `exercises/13_channels/channelsN` directory has a matching `solutions/13_channels/channelsN` directory and a runnable reference `main`.

## Sequential curriculum

| Exercise | Topic | New idea | Why it appears here |
| --- | --- | --- | --- |
| 1 | Unbuffered handoff | A send waits for a receiver | Establishes the blocking model. |
| 2 | Buffered channel | Capacity delays blocking and creates backpressure | Extends the same send/receive model without introducing a new coordination pattern. |
| 3 | Close plus `range` | The sender closes after the final value | Introduces channel lifetime and normal consumer completion. |
| 4 | Comma-ok receive | Closed channel and received zero are different states | Makes the close protocol explicit before it is hidden by `range`. |
| 5 | Directional generator | `<-chan`/`chan<-` express ownership and use direction | Turns the lifecycle rules into a reusable function API. |
| 6 | Basic `select` | Wait for whichever channel operation becomes ready | Builds on ordinary receives and sends. |
| 7 | Select timeout | Stop waiting after a deadline | Applies `select` to a common service-call boundary. |
| 8 | Non-blocking `select` | `default` performs an immediate attempt | Contrasts bounded waiting with polling and introduces the busy-loop warning. |
| 9 | Done notification | `close(done)` broadcasts completion without carrying data | Separates completion coordination from result channels. |
| 10 | Cancellable blocking send | A sender selects between output and cancellation | First direct goroutine-leak prevention exercise. |
| 11 | Timeout, cancellation, and join | The caller cancels and waits for producer exit | Combines exercises 7, 9, and 10 into a complete cleanup protocol. |
| 12 | Broadcast cancellation | One closed signal stops multiple workers | Generalizes cancellation from one goroutine to a task group. |
| 13 | Fan-out | Multiple workers consume one closed jobs channel | Reuses producer ownership and broadcast completion for parallel work. |
| 14 | Fan-in | Multiple inputs forward to one output; one coordinator closes it | Introduces `WaitGroup` only after its channel lifecycle need is visible. |
| 15 | Cancellable fan-in | Forwarders can stop when downstream abandons output | Applies the leak rule to a composite pattern. |
| 16 | Worker pool | Jobs, workers, results, and close order form one lifecycle | Combines fan-out, fan-in, and waiting into the most common reusable shape. |
| 17 | Ordered worker-pool results | Carry an index and restore order after parallel work | Addresses the common mismatch between completion order and business order. |
| 18 | First-error cancellation | A worker reports an error and all remaining work can stop | Adds production-style failure handling to the worker pool. |
| 19 | Single pipeline stage | A stage ranges input and closes its own output | Reuses generator ownership and makes stage boundaries explicit. |
| 20 | Multi-stage pipeline | Several stages compose through channels | Shows that each stage needs an independent close path. |
| 21 | Cancellable pipeline | Cancellation must cover both receiving and sending | Repeats the most important leak-prevention rule at pipeline scale. |
| 22 | Semaphore | A buffered channel limits active work | Uses channel capacity as a concurrency bound rather than a data queue. |
| 23 | Rate limiter | A ticker channel gates operation starts | Covers a common service-protection pattern in one focused exercise. |
| 24 | Nil channel switching | A nil case can disable one `select` branch | A less-common but useful advanced technique, kept at the end. |

## Exercise contract

Each learner directory contains:

- `main.go` with a short concept, task, expected behavior, and a hint that names the intended pattern without giving the full implementation;
- `main_test.go` with deterministic value and lifecycle assertions;
- no dependency on the reference solution.

Each solution directory contains a complete `main.go` that demonstrates the same pattern with a small runnable example. The solution may use Go 1.26's `sync.WaitGroup.Go`, while the learner hint and README also explain the equivalent `Add`/`Done` form when that distinction is pedagogically useful.

## Testing and verification

The PR must pass:

- `gofmt` over all tracked Go files;
- exercise/solution directory parity in the existing CI layout check;
- `go test ./solutions/...`;
- `go vet ./solutions/...`;
- `sh check.sh solutions --run-all`;
- `sh check.sh solutions --run-all --race`;
- focused tests for all 24 channel exercises with `-race`.

Tests for cancellation and cleanup will use explicit completion channels or a bounded test timeout. They will not use `runtime.NumGoroutine` as the sole assertion because scheduler and test-runner goroutines make that count noisy.

## Documentation

`exercises/13_channels/README.md` will be rewritten as the chapter guide. It will contain:

- the sender/receiver/closer ownership model;
- a 24-exercise map with frequency and difficulty bands;
- a short “what the previous exercise prepared” note for each band;
- timing diagrams in prose for the cancellation and close-order exercises;
- a production note that `context.Context` is preferred for request-scoped cancellation and is covered by `exercises/15_context`;
- a troubleshooting checklist for deadlocks, blocked sends, missing closes, and leaked goroutines.

The root `README.md` channel entry will be updated from the old six-exercise count to 24 and will point readers to the chapter guide.

## Migration and scope

The current `channels1`–`channels8` directories will be moved or rewritten into their new sequential positions as part of this PR. The corresponding solution directories and chapter documentation will move together. No other chapter, the `working-progress` worktree, or unrelated user changes will be modified.
