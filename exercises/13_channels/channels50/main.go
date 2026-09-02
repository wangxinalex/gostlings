// Capstone: combine only protocols already learned: bounded work, request/reply,
// backpressure, cancellation, ordered results, error propagation, and one closer.
// Problem: a service must preserve all of these contracts while shutting down
// under either caller cancellation or the first request error.
// Without this composition: each individual pattern may work, but their close
// and join responsibilities can conflict and leak goroutines.
// Channels: unbuffered indexed jobs provides backpressure; request.reply routes
// one private response; stop/internal cancel request shutdown; results carries
// indexed responses; failure reports the first error; worker exits prove cleanup.
// Timeline:
//   producer sends a request
//   worker sends a private reply and shared result
//   an error or stop requests cancellation
//   coordinator joins workers
//   coordinator closes results and done
//   collector returns values in input order
//
// Hint:
//   Build the service in layers.
//   The producer sends indexed requests through an unbuffered jobs channel.
//   It selects on stop and internal cancel.
//   Workers receive requests and process them.
//   They send one reply when reply != nil.
//   They publish indexed responses with cancellation around every blocking
//   operation.
//   The failure path sends the first error to a capacity-one channel.
//   It closes internal cancel once and stops producing new requests.
//   Join every worker before closing results and returning.
//   The collector stores successful indexed responses by original index.
//   It returns them in request order after results closes.
//   Return errStopped for caller cancellation, or the first request error.
//   Only the coordinator closes shared channels.
//   Request reply channels are caller-owned destinations.

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
	// TODO: bound, cancel, reply, order, propagate errors, and join one
	// raw-channel service.
	return nil, nil
}

func main() { fmt.Println(runService(make(chan struct{}), 1, nil)) }
