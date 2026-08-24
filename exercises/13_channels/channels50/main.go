// Capstone: combine only protocols already learned: bounded work, request/reply,
// backpressure, cancellation, ordered results, error propagation, and one closer.
// Problem: a service must preserve all of these contracts while shutting down
// under either caller cancellation or the first request error.
// Without this composition: each individual pattern may work, but their close
// and join responsibilities can conflict and leak goroutines.
// Channels: unbuffered indexed jobs provides backpressure; request.reply routes
// one private response; stop/internal cancel request shutdown; results carries
// indexed responses; failure reports the first error; worker exits prove cleanup.
// Timeline: produce request -> worker reply/result -> error or stop -> cancel -> join -> close results/done -> return ordered values
// Hint: build the service in layers. The producer sends indexed requests through
// an unbuffered jobs channel, selecting on stop/internal cancel. Workers receive
// requests, process them, send one reply when reply != nil, and publish indexed
// responses with stop/internal cancel around every potentially blocking operation.
// The failure path sends the first error to a capacity-one channel, closes
// internal cancel once, and stops producing new requests. Join every worker
// before closing results and returning. The collector stores successful indexed
// responses by original index and returns them in request order after results
// closes. Return errStopped for caller cancellation, or the first request error.
// Only the coordinator closes shared channels; request reply channels are
// caller-owned destinations.
package main

import (
	"errors"
	"fmt"
)

type response struct {
	value int
	err   error
}

type request struct {
	value int
	reply chan response
	err   error
}

var errStopped = errors.New("service stopped")
var processServiceRequest = func(current request) response {
	return response{value: current.value * 2, err: current.err}
}

func runService(stop <-chan struct{}, workers int, requests []request) ([]response, error) {
	return nil, nil // TODO: bound, cancel, reply, order, propagate errors, and join one raw-channel service
}

func main() { fmt.Println(runService(make(chan struct{}), 1, nil)) }
