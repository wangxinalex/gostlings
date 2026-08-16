// Concept: a capstone service combines bounded workers, request/reply, backpressure, cancellation, ordered results,
// error propagation, and coordinator-owned closure using raw channels only.
// request carries value, a per-request reply channel, and err. response carries value and err. Send exactly one
// response on a non-nil reply channel for every request accepted before stop; request.err becomes response.err and
// the returned error. A closed stop is the caller-provided timeout/cancellation signal.
// Hint: build the service in layers:
//       producer: send indexed requests through an unbuffered jobs channel, selecting on stop/internal cancel;
//       workers: receive requests, process them, send one reply when reply != nil, and publish indexed
//       responses, with stop/internal cancel around every potentially blocking send or receive;
//       failure path: send the first error to a capacity-one channel, close internal cancel once, and stop
//       producing new requests; join every worker before closing results and returning.
//       The collector stores successful indexed responses by original index. After results closes, return
//       them in request order. Return errStopped for caller cancellation, or the first request error.
//       Only the coordinator closes shared channels; request reply channels are caller-owned destinations.
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
