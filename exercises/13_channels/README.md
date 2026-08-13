# Channels: from raw protocols to concurrent composition

This chapter uses 50 small exercises to build a lifecycle model for channels:
who sends, who receives, who knows no more values will be sent, who closes,
and who cancels when the caller gives up early. It deliberately sticks to raw
channels; `context.Context`, `sync`, timers, and integrated concurrency
patterns belong to later chapters. After finishing this chapter, hand the more
complete pipeline work to
[`24_concurrency_patterns`](../24_concurrency_patterns).

This chapter does not build reusable timer/ticker abstractions; a few timeout
or cancellation tests/examples may still use `time.After`. `21_time` is
dedicated to constructing, resetting, and managing the lifecycle of
`time.Timer`/`time.Ticker` and to periodic designs.

## Ownership model

- The sender usually closes; only the party that knows "no more values will
  be sent" may `close(ch)`.
- Receivers usually `range ch` or use comma-ok receives; they do not close the
  producer's output.
- Closing marks the end of the lifecycle, not a special data value; buffered
  values are still drained first.
- A `done` channel signals completion, a data channel carries values, and a
  `stop`/`cancel` channel broadcasts cancellation.
- A shared output can only be closed by an explicit coordinator after every
  sender has exited.
- Every potentially blocking receive, send, or token wait needs an exit path
  in the cancellation protocol.

## Skill matrix and difficulty

| Skill | Exercises |
| --- | --- |
| Unbuffered/buffered, close, comma-ok, drain | 1–5 |
| Generator, directional constraints, `select`, `default`, timeout | 6–11 |
| Cancelable receive/send, done, result, stop, join | 12–20 |
| Fan-out, fan-in, worker pool, errors and direction | 21–30 |
| First result, request/reply, shared-output coordination, control state | 31–40 |
| nil channel, canceling producers, semaphores, rate limiting, service shutdown | 41–50 |

| Difficulty band | Exercises | Goal |
| --- | --- | --- |
| Basics | 1–10 | Understand blocking, closing, and one-shot `select` |
| Transition | 11–20 | Combine timeout, cancellation, and completion into waitable lifecycles |
| Composition | 21–30 | Organize workers, shared output, and backpressure |
| Advanced | 31–40 | Handle first results, replies, state machines, and coordinated shutdown |
| Capstone | 41–50 | Keep ownership clear in bounded, ordered, cancellable services |

## 1–50 learning map

| # | Pattern |
| ---: | --- |
| 1 | Unbuffered handoff: send and receive must be in different goroutines |
| 2 | Buffered channel decouples send and receive |
| 3 | Closing a channel stops `range` |
| 4 | Comma-ok distinguishes close from a real zero value |
| 5 | Drain buffered values before closing the channel |
| 6 | Producer owns and closes the generator output |
| 7 | Directional channels express generator ownership |
| 8 | `select` multiplexes multiple inputs |
| 9 | `select` + `default` non-blocking receive |
| 10 | `select` + `default` non-blocking send |
| 11 | `select` timeout |
| 12 | Relay cancellable at both the receive and the send end |
| 13 | Closing a done channel broadcasts completion |
| 14 | Capacity-one asynchronous result |
| 15 | Separate result from done signal |
| 16 | Cancelable producer releases a blocked send |
| 17 | Cancelable forwarder that can block at both ends |
| 18 | Start, broadcast stop, and join workers with raw channels |
| 19 | Full cleanup: timeout → cancel → join |
| 20 | Stop, join workers, then report graceful shutdown |
| 21 | Fan-out workers share jobs, then fan-in results |
| 22 | Fan-in multiple inputs and coordinate closing the output |
| 23 | Cancelable fan-in with exit paths at both ends |
| 24 | Worker pool with cancellable jobs receive and result send |
| 25 | Fan-in that handles nil, closed, and buffered inputs |
| 26 | Basic pool: jobs producer, workers, results closer |
| 27 | Carry an index to restore input order in a worker pool |
| 28 | First error stops new work and joins all workers |
| 29 | Expose a minimal interface through channel direction |
| 30 | Bounded jobs channel provides backpressure |
| 31 | First completed result cancels peers and waits for exit |
| 32 | Request carries a private reply channel |
| 33 | Multiple forwarders, one coordinator closes the shared output |
| 34 | Define pool behavior for zero workers |
| 35 | Cancelable worker group protects receive and send |
| 36 | Result envelope carries both value and error |
| 37 | Notify upstream when the downstream gives up |
| 38 | One done channel broadcasts to multiple observers |
| 39 | pause/resume/stop control commands coexist with work |
| 40 | Relay cancels from a blocked receive or send |
| 41 | Set a closed input to nil so a closed channel is not always ready |
| 42 | Stop releases a producer send with no receiver |
| 43 | Buffered token channel as semaphore, restoring input order |
| 44 | Semaphore acquisition and result publishing are both cancellable |
| 45 | Raw-channel rate limiter with caller-provided tokens |
| 46 | Rate limiter with cancellable input, tokens, and output |
| 47 | Raw-channel service with results/done closed by the coordinator |
| 48 | Bounded, ordered, cancellable worker pool |
| 49 | First job error cancels new work; a single owner closes cancel |
| 50 | Bounded request/reply service: backpressure, cancellation, order, errors, and shutdown |

## Cancellation timing

Normal shutdown:

```text
producer -- values --> out -- range --> consumer
   |                       |
   +------ close(out) -----+--> range ends after buffered values drain
```

When the downstream exits early, the send side must be cancellable; you must
also join after cancelling:

```text
caller -- close(stop) --> producer/forwarders -- exit --> coordinator
caller <-- return only after done/result closure ----------------------+
```

The common sequence for fan-in, worker pools, and the capstones is: stop
delivering new work, wait for all senders to exit, then let one coordinator
close the shared output (and the done channel if needed). Never rely on
`Sleep` to guess that a goroutine has finished.

## Hints are strict contracts

Each `main.go`'s Concept / Task / Expected behavior / Hint is the minimal
contract for that exercise, not optional advice. In particular, keep the
closer, channel directions, cancellation coverage, buffer capacity, index,
reply count, and join order as the hint specifies; do not use empty slices,
early returns, `Sleep`, or ignored errors to create implementations that
merely look like they pass.
When stuck, first draw the senders, receivers, closers, and every exit edge,
then change the smallest piece.

## Checker and exercise order

Single exercise, full chapter, and starter verification:

```sh
sh check.sh exercises/13_channels/channels1
sh check.sh exercises/13_channels --run-all
sh scripts/verify_exercise_starters.sh exercises/13_channels
```

The checker is behavior based: every exercise ships a `main_test.go`, so an
untouched starter fails its focused behavioral test and a solution that merely
exits 0 cannot pass. `verify_exercise_starters.sh` additionally checks that
every unmodified starter really fails, so a vacuous implementation cannot
pass. `--race` adds `-race` only to exact paths listed in `race.list`;
`--race-all` applies it to every selected target and is the expensive full
audit:

```sh
sh check.sh solutions/16_sync/sync1 --race
sh check.sh solutions --run-all --race
sh check.sh solutions --run-all --race-all
```

A practical order is to finish the semantics and exit protocols of 1–20, then
the compositions of 21–40, and finally the boundary and service-shutdown
exercises 41–50. `15_context`, `16_sync`, `21_time`, and
`24_concurrency_patterns` migrate these raw-channel timings to higher-level
APIs; do not use them to replace channel protocols in this chapter.
