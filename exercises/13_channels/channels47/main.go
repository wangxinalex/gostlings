// Concept: a raw-channel service owns a result stream and a done signal, but workers own neither close.
// Task: serve requests until jobs closes or stop closes, then wait for active workers before closing results and done.
// request carries value, a per-request reply channel, and an optional err. response carries value and err. Send
// exactly one response on a non-nil reply channel for each request that is accepted before shutdown.
// Hint: workers select on stop before accepting jobs and before both reply/result sends. A coordinator receives one
// exit acknowledgement per worker, closes results, then closes done. Only that coordinator closes shared channels.
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

func serve(stop <-chan struct{}, jobs <-chan request) (<-chan response, <-chan struct{}) {
	return nil, nil // TODO: stop accepting work, join workers, close results, then close done
}

func main() { fmt.Println(serve(make(chan struct{}), nil)) }
