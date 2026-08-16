// Concept: a raw-channel service owns a result stream and a done signal, but workers own neither close.
// Task: serve requests until jobs closes or stop closes, then wait for active workers before closing results and done.
// request carries value, a per-request reply channel, and an optional err. response carries value and err. Send
// exactly one response on a non-nil reply channel for each request that is accepted before shutdown.
// Hint: follow the lifecycle in this order:
//       workers select on stop before receiving a job; a closed jobs channel ends a worker normally.
//       For an accepted request, compute a response and attempt exactly one send on its non-nil
//       reply channel, with a stop case. If cancellation has not won, publish the same response
//       to results, also with a stop case.
//       Each worker sends one exit acknowledgement on every return path.
//       A coordinator receives one acknowledgement per worker, closes results, then closes done.
//       Workers never close shared channels. stop can cancel a blocked reply or result send, and done
//       must not close until all workers have joined.
package main

import "fmt"

type response struct {
	value int
	err   error
}

type request struct {
	value int
	reply chan response
	err   error
}

var onServeBeforeResult = func() {}

var serveWorkerCount = 2

func serve(stop <-chan struct{}, jobs <-chan request) (<-chan response, <-chan struct{}) {
	return nil, nil // TODO: stop accepting work, join workers, close results, then close done
}

func main() { fmt.Println(serve(make(chan struct{}), nil)) }
