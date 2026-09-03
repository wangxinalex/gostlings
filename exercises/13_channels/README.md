# Channels: learn the protocol before composing the pattern

This chapter teaches channels as small communication protocols. A channel is
not merely a queue: it can transport data, synchronize a handoff, broadcast
cancellation, or announce completion. The important questions are always:

```text
Who sends?
Who receives?
Who knows that no more values will be sent?
Who closes the channel?
What event lets every goroutine stop?
```

Read the exercises from `channels1` through `channels30`. Later exercises reuse
earlier protocols and add one constraint at a time. If an exercise feels
arbitrary, first identify the production failure it prevents; the pattern is
there to solve that failure.

## The channel ledger

Before implementing a channel-based function, write one row per channel:

```text
channel | carries | direction | owner | who closes it | what closing means
```

Typical roles are different even when their types look similar:

| Role | Carries | Usual closer | Meaning of close |
| --- | --- | --- | --- |
| Data channel | Values used by the application | The sender that knows there are no more values | The stream is complete; buffered values can still be drained |
| Stop channel | No data, only cancellation | The caller or coordinator that owns cancellation | Every receiver of `<-stop` may stop its work |
| Done channel | No data, only completion | The goroutine or coordinator that knows work is finished | Waiting receivers may continue after cleanup is complete |
| Exit acknowledgement | One token per worker | Each exiting worker sends one token | This worker has reached its exit point |

Closing a channel and sending a value are not interchangeable. Closing a stop
channel broadcasts one event to every receiver. Closing a done channel reports
completion to every waiter. A data channel should be closed only by the party
that owns the send side and knows that no future send is possible.

## Timeline notation

Hints use short ASCII timelines. A vertical bar means that a goroutine is
waiting; an arrow is a send, receive, or notification; `close(x)` is a state
change, not a data value.

For a cancellable producer with a timeout:

```text
caller                  producer                 result
  | start producer          |                       |
  |----------------------->|                       |
  | wait for value/timeout  | slow work             |
  | timeout                 |                       |
  | close(stop)             |                       |
  |----------------------->| observe stop          |
  |                         | close(done)           |
  |<------------------------|                       |
  | <-done: producer exited |                       |
  | return timed out        |                       |
```

Here `stop` is the request to stop, while `done` is proof that the producer
actually stopped. Returning after `close(stop)` but before `<-done` can leave a
goroutine running. The data channel `result` carries the business result and
must not be used as a substitute for the completion signal.

## Learning stages

### Stage 1: Mechanics and ownership — exercises 1–7

Learn unbuffered rendezvous, buffering, FIFO delivery, closing, `range`,
comma-ok receives, draining, generators, and directional channels. These answer
the basic ownership questions before any cancellation is introduced.

### Stage 2: `select` as an escape path — exercises 8–12

Learn multiplexing, non-blocking attempts, timeouts, and cancellable blocking
operations. The key idea is that `select` adds an alternative at a place that
could otherwise block:

```go
select {
case value := <-in:
    use(value)
case <-stop:
    return
}
```

`default` means “try immediately”; it does not mean “wait eventually”. A
timeout is another channel case that becomes ready later. If both cases are
ready, `select` does not provide priority.

### Stage 3: Results, cancellation, and completion — exercises 13–20

Learn the difference between data, cancellation, and completion. Start with a
done broadcast, then an asynchronous result, then a producer that can be
cancelled while sending. Exercise 18 introduces worker exit acknowledgements;
exercise 19 combines timeout, cancellation, and join; exercise 20 applies the
same protocol to a graceful multi-worker shutdown.

Do not treat `stop`, `exited`, and `done` as three names for the same thing:

```text
stop   = request that work should stop
exited = one worker confirms that it stopped
done   = coordinator confirms that every worker stopped
```

### Stage 4: Fan-out and fan-in — exercises 21–25

Fan-out solves “one jobs stream must be processed concurrently”: workers share
one input and each job is received by one worker. Fan-in solves “several
independent streams must become one stream”: one forwarder is started per input.

For a shared output, forwarders must not close the output themselves:

```text
input A ---> forwarder A --+
input B ---> forwarder B --+--> out ---> consumer
input C ---> forwarder C --+       ^
                                coordinator closes out
```

The coordinator waits for one exit acknowledgement per sender and closes `out`
exactly once. Exercises 23–24 add cancellation to the receive and send sides;
exercise 25 isolates empty, closed, buffered, and nil-input lifecycle edges.

### Stage 5: Worker pools and constraints — exercises 26–30

Build a pool from the pieces already learned: a jobs producer, shared jobs
channel, workers, results channel, and one results closer. Then add one real
constraint at a time:

```text
basic pool -> ordered results -> error cancellation -> directional handles
            -> bounded queue
```

The pool exists to bound concurrency and reuse workers. A bounded jobs channel
adds backpressure: the producer must slow down when workers cannot keep up.
Indexed results restore input order because parallel completion order is not
input order. Error cancellation stops new work but still requires every worker
to exit before the function returns.

## Exercise map

| Exercises | Primary question |
| --- | --- |
| 1–2 | When does an unbuffered or buffered send block? |
| 3–5 | How does a receiver know that a data stream is complete? |
| 6–7 | Who owns a producer's output and its close? |
| 8–11 | How can a wait have another ready alternative? |
| 12 | How do both sides of a relay stop safely? |
| 13–15 | How are completion and result data represented separately? |
| 16–17 | How do cancellation-aware sends and receives prevent leaks? |
| 18–20 | How do workers broadcast stop and prove that cleanup finished? |
| 21 | How are jobs shared by a group of workers? |
| 22–25 | How are input streams merged and safely closed? |
| 26 | How is a basic worker pool built? |
| 27 | How are parallel results returned in input order? |
| 28 | How does error cancellation stop new work and join every worker? |
| 29 | How do directional channel handles express ownership? |
| 30 | How does a bounded queue apply backpressure? |

## A repeatable implementation checklist

For every exercise:

1. Draw the senders, receivers, and closer.
2. Write the channel ledger.
3. Mark every potentially blocking receive and send.
4. Add the exit case for each blocking operation if cancellation is part of the contract.
5. Decide who waits for each goroutine and what signal proves it has exited.
6. Only then write the smallest code change that satisfies the TODO.

The checker is behavior based. Run one exercise while learning, then the full
chapter and starter audit:

```sh
sh check.sh exercises/13_channels/channels1
sh check.sh solutions/13_channels --run-all
sh scripts/verify_exercise_starters.sh exercises/13_channels
```

Use Go 1.26.5, as required by the repository's `go.mod`. `--race` is useful
when a pattern shares state or coordinates multiple goroutines.
