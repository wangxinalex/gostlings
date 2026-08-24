// Problem: a raw-channel service must serve requests, route per-request replies,
// publish a shared result stream, and shut down without abandoning workers.
// Without this pattern: workers can outlive the service, callers can miss their
// replies, or several workers can race to close shared channels.
// Channels: jobs carries requests; each request.reply is a private destination;
// stop cancels; results carries shared responses; exited joins workers; done is
// the final completion signal. Workers close none of the shared channels.
// Timeline: receive job/stop -> reply and publish result -> worker exits -> coordinator closes results -> close(done)
// Hint: follow the lifecycle in this order:
//
//	workers select on stop before receiving a job; a closed jobs channel ends a worker normally.
//	For an accepted request, compute one response and attempt one reply send with a stop case.
//	If cancellation has not won, publish the same response to results with another stop case.
//	Each worker sends one exit acknowledgement on every return path.
//	A coordinator receives one acknowledgement per worker, closes results, then closes done.
//	Workers never close shared channels; done cannot close until every worker has joined.
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
