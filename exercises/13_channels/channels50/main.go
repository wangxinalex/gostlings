// Concept: a capstone service combines bounded workers, request/reply, backpressure, cancellation, ordered results,
// error propagation, and coordinator-owned closure using raw channels only.
// request carries value, a per-request reply channel, and err. response carries value and err. Send exactly one
// response on a non-nil reply channel for every request accepted before stop; request.err becomes response.err and
// the returned error. A closed stop is the caller-provided timeout/cancellation signal.
// Hint: feed indexed requests through an unbuffered jobs channel, start only workers workers, use a capacity-one
// error channel, select on stop/internal cancel around every handoff, and have one coordinator close results after
// all workers exit. Store received indexed responses by their original index before returning.
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
