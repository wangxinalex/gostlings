// Problem: even a correctly joined worker pool can leak if cancellation reaches
// only its input or only its output.
// Without this pattern: a worker may wait forever for a job or for a result
// receiver after the caller has abandoned the pool.
// Channels: jobs carries work; stop is a receive-only cancellation broadcast;
// out carries results; exited confirms every worker returned.
// Timeline: select jobs/stop -> compute -> select out/stop -> all exited -> close(out)
// Hint: protect both blocking operations with select. Workers acknowledge exit
// on every return path; the coordinator receives every acknowledgement before
// closing out. This is the worker-pool version of the two-select relay pattern.
package main

import "fmt"

var onFanOutBeforeSend = func() {}

func fanOut(stop <-chan struct{}, jobs <-chan int, workers int) <-chan int {
	return nil // TODO: make receives and sends cancellable
}
func main() { fmt.Println(fanOut(make(chan struct{}), nil, 1)) }
